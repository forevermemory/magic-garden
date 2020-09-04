package service

/*
date:2020-07-30 20:00:37
*/

import (
	"errors"
	"fmt"
	"magic/db"
	"magic/global"
	"magic/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// InitGarden 初始化花园
func InitGarden(req *global.UserAddGamesParams) error {
	_, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return err
	}
	// 首先需要判断是否需要重新初始化 important
	_, err = db.GetGardenByID(req.UserID)
	if err == gorm.ErrRecordNotFound {
		// 开始事物
		tx := global.MYSQL.Begin()
		// 1. 创建一个花园
		garden := &db.Garden{
			ID:    req.UserID,
			GName: req.Gname,
			GInfo: req.Ginfo,
			// GName:      user.Username + "的花园",
			// GInfo:      "劳动可耻,偷窃光荣!",
			GLevel:     1,
			GAtlas:     0,
			IsSignin:   0,
			SignDays:   0,
			GCurrentEx: 0,
		}
		if err = tx.Create(garden).Error; err != nil {
			tx.Rollback()
			return err
		}
		// 2.初始化10个花盆
		for i := 1; i < 11; i++ {
			// 2.1 只设置两个花盆解锁
			var islock = 1
			if i <= 2 {
				islock = 2
			}
			huapen := &db.GardenFlowerpot{
				UserID:   req.UserID,
				GardenID: req.UserID,
				Number:   i,
				IsLock:   islock,
				IsSow:    1,
			}
			if err = tx.Create(huapen).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		// 初始化花房 花瓶 好像不需要初始化哦 到时候写在收获中

		// 初始化背包 道具 种子

		//
		tx.Commit()
	} else {
		return nil
	}

	return nil
}

// ListGardenKnapsack  查询列表
func ListGardenKnapsack(b *db.GardenFlowerKnapsack) (*db.DataStore, error) {
	list, err := db.ListGardenFlowerKnapsack(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenFlowerKnapsack(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + global.PageSize - 1) / global.PageSize}, nil
}

// UpdateGarden update 修复了会更新其它字段的bug 只会更新名称和公告
func UpdateGarden(b *db.Garden) (*db.Garden, error) {
	return db.UpdateGarden(b)
}

// GetGardenByID get by id
func GetGardenByID(id string) (*db.Garden, error) {
	return db.GetGardenByID(id)
}

// GetGardenHelpByID get by id
func GetGardenHelpByID(id int) (*db.GardenHelp, error) {
	return db.GetGardenHelpByID(id)
}

// GetGardenHelpTitles 花园帮助标题列表
func GetGardenHelpTitles() ([]*db.GardenHelp, error) {
	return db.GetGardenHelpTitles()
}

// GardenEveryDaySignin 花园签到  TODO
func GardenEveryDaySignin(req *global.GardenParams) (interface{}, error) {
	// 1.查询花园
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	results := make([]map[string]interface{}, 0)
	// 2.是否签到了
	if garden.IsSignin == 1 {
		//2.1 未签到
		days := garden.SignDays
		if days < 7 {
			days++
			garden.GSigninTime = utils.GetNowTimeString()
		}
		// 事物对象
		tx := global.MYSQL.Begin()

		// 查询签到的奖励 道具和种子
		seeds, err := db.GetSignInRewardsSeed(days)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		props, err := db.GetSignInRewardsProp(days)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		var signinMoney int
		var signinExp int
		for _, seed := range seeds {
			temp := make(map[string]interface{})
			temp["o_name"] = seed.RSeedName
			temp["o_num"] = seed.RSeedNum
			temp["o_id"] = seed.RSeedID
			results = append(results, temp)
			// 将种子存入背包
			// err = tx.Table("garden_flower_knapsack").Where("garden_id = ? and seed_id = ?").Error
			gardenKnap, err := db.IsExistGardenFlowerKnapsackSeed(garden.ID, seed.RSeedID)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					// 不存在记录 新建一条记录
					sqlSeed := "INSERT INTO garden_flower_knapsack  VALUES (null, ?, ?, ?, 1, NULL, NULL);"
					if err = tx.Exec(sqlSeed, garden.ID, seed.RSeedID, seed.RSeedNum).Error; err != nil {
						tx.Rollback()
						return nil, err
					}
					// 插入成功
					continue
				} else {
					// 其它查询错我
					tx.Rollback()
					return nil, err
				}
			}
			// 更新该种子的数量
			gardenKnap.SeedNum += seed.RSeedNum
			if err = tx.Exec("update garden_flower_knapsack set seed_num = ? where garden_id = ? and seed_id = ?", gardenKnap.SeedNum, garden.ID, seed.RSeedID).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

		}

		for _, prop := range props {
			temp2 := make(map[string]interface{})
			temp2["o_name"] = prop.RPropName
			temp2["o_num"] = prop.RPropNum
			temp2["o_id"] = prop.RPropID
			signinMoney = prop.RGb
			signinExp = prop.RExp

			results = append(results, temp2)
			// 将道具存入背包
			gardenKnapProp, err := db.IsExistGardenFlowerKnapsackProp(garden.ID, prop.RPropID)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					// 不存在记录 新建一条记录
					sqlProp := "INSERT INTO garden_flower_knapsack  VALUES (null, ?, null, null, 2, ?, ?);"
					if err = tx.Exec(sqlProp, garden.ID, prop.RPropID, prop.RPropNum).Error; err != nil {
						tx.Rollback()
						return nil, err
					}
					// 插入成功
					continue
				} else {
					// 其它查询错误
					tx.Rollback()
					return nil, err
				}
			}
			// 更新该种子的数量 有问题哦
			gardenKnapProp.PropNum += prop.RPropNum
			if err = tx.Exec("update garden_flower_knapsack set prop_num = ? where garden_id = ? and prop_id = ?", gardenKnapProp.PropNum, garden.ID, prop.RPropID).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		// 需要计算是否当天可以继续获得经验值 TODO

		// exp和gb也返回
		expAndGb := make(map[string]interface{})
		expAndGb2 := make(map[string]interface{})
		expAndGb["o_name"] = "经验值"
		expAndGb["o_num"] = signinExp
		results = append(results, expAndGb)
		expAndGb2["o_name"] = "GB"
		expAndGb2["o_num"] = signinMoney
		results = append(results, expAndGb2)
		// 更新 gb
		user, err := db.GetUsersByID(req.GardenID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		gmoney := user.GBMoney
		gmoney += signinMoney

		updateGmoneySQL := "update users set gb_money = ? where _id = ?"
		if err = tx.Exec(updateGmoneySQL, strconv.Itoa(gmoney), user.ID).First(user).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		// exp // 最后更新 登陆天数
		totalExp := garden.GTotaltEx + signinExp
		curExp := garden.GCurrentEx + signinExp
		updateGardenExpSQL := "update garden set g_current_ex = ?,g_total_ex = ?,is_signin = 2,sign_days = ?,g_signin_time = ? where _id = ?"
		if err = tx.Exec(updateGardenExpSQL, curExp, totalExp, strconv.Itoa(days), utils.GetNowTimeString(), garden.ID).First(garden).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return nil, err
		}
		// 需要搞一个方法 根据总的经验数量计算花园的等级
		if garden, err = garden.ComputeCurrentLevel(tx); err != nil {
			tx.Rollback()
			return nil, err
		}
		// 存入历史表
		if _, err = db.AddGardenSigninHistory(&db.GardenSigninHistory{
			HGardenID: req.GardenID,
			HUserID:   req.GardenID,
			HTime:     utils.GetNowTimeString(),
		}, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
		// 事物提交
		if err = tx.Commit().Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		// 存入历史 这里用消息队列比较好 算了吧
		go db.SaveGbHistory(signinMoney, "每日签到", "", garden)
		return results, nil
	} else if garden.IsSignin == 2 {
		// 2.2当天已经签到
		results = append(results, map[string]interface{}{
			"o_name": "你今天签到过了,明天记得来领取签到奖励哦!",
		})
		return results, nil
	} else {
		return "请正确操作", errors.New("err please operate correctly")
	}

}

// ListGardenSigninHistory  查询花园签到历史
func ListGardenSigninHistory(b *db.GardenSigninHistory) (*db.DataStore, error) {
	list, err := db.ListGardenSigninHistory(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenSigninHistory(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + global.PageSize - 1) / global.PageSize}, nil
}

// GardeFlowerpotList  查看可用花盆列表
func GardeFlowerpotList(b *global.GardenPotParams) ([]*db.GardenFlowerpot, error) {
	return db.ListGardenFlowerpot(b.GardenID)
}

// GardeFlowerpotDetail  查看花盆详情
func GardeFlowerpotDetail(b *global.GardenPotParams) (*db.GardenFlowerpot, error) {
	return db.GetGardenFlowerpotByID(b.GardenID, b.FlowerpotID)
}

// GardeFlowerpotSow  花盆播种
func GardeFlowerpotSow(req global.GardenPotParams) (interface{}, error) {

	result := make(map[string]interface{})
	tx := global.MYSQL.Begin()

	// 查询出 garden
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 查询出 背包中种子的数量
	seedKnap, err := db.CountGardenFlowerKnapsackV2(req.GardenID, req.SeedID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result["result"] = "您没有该种子哦,,请正确操作"
			tx.Rollback()
			return result, nil
		}
		tx.Rollback()
		return nil, err
	}
	if seedKnap.SeedNum == 0 {
		result["result"] = "该种子已经用完了哦,请正确操作"
		tx.Rollback()
		return result, nil
	}
	fmt.Println("req.IsVip::", req.IsVip)

	// 1.判断是不是vip
	if req.IsVip == 2 {
		user, err := db.GetUsersByID(req.GardenID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// 是否是vip
		if user.IsVip == 1 {
			// 不是vip
			result["result"] = "您还不是vip哦,无法享受一键播种的特权"
			tx.Rollback()
			return result, nil
		} else if user.IsVip == 2 {
			// 进行一键播种
			// TODO
			pots, err := db.GetCanSowGardenFlowerpots(req.GardenID)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			// 遍历花🌹盆们
			if len(pots) == 0 {
				result["result"] = "您当前无可播种花盆"
				tx.Rollback()
				return result, nil
			}
			totalJy := 0
			// ps  因为这里用到了事物 即使有下面的 update 但是事物没有提交 查询到的还是之前的数据
			// times := 0
			// seedCount 不需要再查询一遍 上面已经查询过了
			// seedCount, err := db.CountGardenFlowerKnapsackV2(req.GardenID, req.SeedID)
			// // 没有该种子退出循环
			// if err != nil {
			// 	tx.Rollback()
			// 	return nil, err
			// }
			for _, pot := range pots {
				// 当存在的种子数量比空花盆少的时候 只播种种子的数量花盆
				fmt.Println("seedCount.SeedNum start---", seedKnap.SeedNum)
				getjy, err := sow(req.SeedID, pot, tx, garden)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
				totalJy += getjy
				// 种子-- TODO
				seedKnap.SeedNum = seedKnap.SeedNum - 1
				// 更新到数据库 这里还是感觉写的不好 可以改进 seedKnap.SeedNum == 0  下面在循环结束后还是需要判断update
				if seedKnap.SeedNum == 0 {
					// _, err = db.UpdateGardenFlowerKnapsackHandleSeedNumZelo(seedKnap, tx)
					// if err != nil {
					// 	tx.Rollback()
					// 	return nil, err
					// }
					break
				}
				fmt.Println("seedKnap.SeedNum end----", seedKnap.SeedNum)

			}
			//  下面在循环结束后还是需要判断update
			_, err = db.UpdateGardenFlowerKnapsackHandleSeedNumZelo(seedKnap, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			result["result"] = fmt.Sprintf("播种成功,您获得%v点经验值", totalJy)
			tx.Commit()
			return result, nil
		}
	}
	// 2.普通播种 需要传花盆编号
	// 查询花盆 is_lock = 2 /
	pot, err := db.GetGardenFlowerpotByIDIsCanSow(req.GardenID, req.FlowerpotID)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			result["result"] = "您当前无可播种花盆"
			return result, nil
		}
		return nil, err
	}
	getjy, err := sow(req.SeedID, pot, tx, garden)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 种子-- TODO
	seedKnap.SeedNum = seedKnap.SeedNum - 1
	// 更新到数据库
	fmt.Println("更新到数据库 又是0.。。。。。")
	if _, err = db.UpdateGardenFlowerKnapsackHandleSeedNumZelo(seedKnap, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	result["result"] = fmt.Sprintf("播种成功,您获得%v点经验值", getjy)
	tx.Commit()
	return result, nil

}

// sow 抽离播种方法 返回一个盆播种之后获得的经验值
func sow(seedID int, pot *db.GardenFlowerpot, tx *gorm.DB, garden *db.Garden) (int, error) {
	// 更新该花盆的状态
	pot.IsSow = 2
	pot.SeedID = seedID
	// pot.Disaster = 2 // 干旱
	pot.SowTime = utils.GetNowTimeString()

	// 种子
	seed, err := db.GetGardenSeedsByID(seedID)
	if err != nil {
		return 0, err
	}
	pot.IsChangeColor = seed.IsChangeColor
	_, err = db.UpdateGardenFlowerpot(pot, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// 计算获得的经验数量
	getjy, err := garden.ComputeEmpirical(1, 1, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return getjy, nil
}

// GardeFlowerpotLookAfter  花盆浇水施肥除草 disaster  2干旱(浇水) 3有虫(除虫) 4有草(除草)
func GardeFlowerpotLookAfter(req global.GardenPotParams) (interface{}, error) {

	result := make(map[string]interface{})
	result["total"] = 0
	tx := global.MYSQL.Begin()

	// 查询出 garden0
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 1.判断是不是vip
	if req.IsVip == 2 {
		user, err := db.GetUsersByID(req.GardenID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if user.IsVip == 1 {
			result["result"] = fmt.Sprintf("您还不是vip,无法享受vip特权哦,请正确操作")
			return result, nil
		}
		// 一键操作 TODO
		for _, val := range req.Handle {
			result, err = handleDisaster(tx, req, result, garden, val)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		// 全部操作完成后
		tx.Commit()
		result["result"] = fmt.Sprintf("一键操作成功,您获得%v点经验值", result["total"].(int))
		return result, nil
	}
	// 操作一个花盆
	// [{"number":1,"kind":2}]
	hand := req.Handle[0]
	result, err = handleDisaster(tx, req, result, garden, hand)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, nil

}

func handleDisaster(tx *gorm.DB, req global.GardenPotParams, result map[string]interface{}, garden *db.Garden, hand map[string]int) (map[string]interface{}, error) {
	number := hand["number"]
	kind := hand["kind"]
	pot, err := db.GetGardenFlowerpotByID(req.GardenID, number)
	if err != nil {
		return nil, err
	}
	seed, err := db.GetGardenSeedsByID(pot.SeedID)
	if err != nil {
		return nil, err
	}
	// 先判断是否需要该操作  状态不匹配
	if pot.Disaster != kind || !utils.IsValInSlice(kind, []int{2, 3, 4}) {
		result["result"] = fmt.Sprintf("您无法进行该操作,请正确操作")
		return result, nil
	}
	// 设置花盆状态
	pot.Disaster = 1
	// 花朵🌺数量++ 每处理一次 FlowerNumHandle 增加一
	if seed.ForecastNum > 1 && pot.FlowerNumHandle < seed.ForecastNum*2 {
		pot.FlowerNumHandle++
	}
	if _, err = db.UpdateGardenFlowerpot(pot, tx); err != nil {
		return nil, err
	}
	// 更新经验值
	handleJy, err := garden.ComputeEmpirical(2, 1, tx)
	if err != nil {
		return nil, err
	}

	result["result"] = fmt.Sprintf("%v成功,您获得%v点经验值", global.PotStatus[kind], handleJy)
	result["total"] = handleJy + result["total"].(int)
	return result, nil
}

// GardeFlowerpotRemove  移除花盆中成长的花朵🌹
func GardeFlowerpotRemove(req global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})
	// 更新该花盆的状态
	pot, err := db.GetGardenFlowerpotByID(req.GardenID, req.FlowerpotID)
	if err != nil {
		return nil, err
	}
	tx := global.MYSQL.Begin()

	if err = db.UpdateGardenFlowerpotWithRemove(pot, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result["result"] = "移除成功"
	return result, nil
}

// GardeFlowerpotDyeing  染色
func GardeFlowerpotDyeing(req global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})

	// 背包是否存在染色剂 --
	die, err := db.IsExistGardenFlowerKnapsackProp(req.GardenID, 7)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result["result"] = "您没有染色剂,请正确操作"
			return result, nil
		}
		return nil, err
	}
	if die.PropNum == 0 {
		result["result"] = "您的染色剂已经用完,请正确操作"
		return result, nil
	}

	// 花盆
	pot, err := db.GetGardenFlowerpotByID(req.GardenID, req.FlowerpotID)
	if err != nil {
		return nil, err
	}
	// 种子
	seed, err := db.GetGardenSeedsByID(pot.SeedID)
	if err != nil {
		return nil, err
	}

	// 不可以染色
	if seed.IsChangeColor == 1 {
		result["result"] = "当前花朵暂不支持染色,请正确操作"
		return result, nil
	}
	// 染色
	// 普通的多色花朵 可以染色

	// 图谱s
	atlas, err := db.GetGardenAtlasBySeedID(pot.SeedID)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(atlas)) // [0,n)
	pot.SeedResult = atlas[index].ID
	pot.ChangeResult = atlas[index].FlowerCateName
	pot.SeedResultStr = atlas[index].FlowerCateName
	pot.FlowerNum = seed.ForecastNum + pot.FlowerNumHandle
	pot.IsUseDye = 2
	tx := global.MYSQL.Begin()
	if _, err = db.UpdateGardenFlowerpot(pot, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	// 染色剂 -- TODO
	if err = db.DieReduce(die.PropNum-1, pot.GardenID, die.PropID, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result["result"] = fmt.Sprintf("染色成功,您获得了%s", pot.ChangeResult)
	return result, nil
}

// GardeFlowerpotFertilizer  施肥
func GardeFlowerpotFertilizer(req global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})

	// 查询出该道具
	prop, err := db.GetGardenPropsByID(req.PropID)
	if err != nil {
		return nil, err
	}

	// 背包是否存在肥料 --
	die, err := db.IsExistGardenFlowerKnapsackProp(req.GardenID, req.PropID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result["result"] = fmt.Sprintf("您没有%v,请正确操作", prop.PName)
			return result, nil
		}
		return nil, err
	}
	if die.PropNum == 0 {
		result["result"] = fmt.Sprintf("%v已经用完,请正确操作", prop.PName)
		return result, nil
	}

	// 花盆
	pot, err := db.GetGardenFlowerpotByID(req.GardenID, req.FlowerpotID)
	if err != nil {
		return nil, err
	}
	if pot.IsHarvest == 2 {
		result["result"] = fmt.Sprintf("已经可以收获了哦,请不要重复操作")
		return result, nil
	}
	// 种子
	// seed, err := db.GetGardenSeedsByID(pot.SeedID)
	// if err != nil {
	// 	return nil, err
	// }
	// 计算剩余参数
	// 计算缩短的时间
	var interval int
	switch req.PropID {
	case 1:
		interval = 0.5 * 3600
	case 2:
		interval = 3600
	case 3:
		interval = 2 * 3600
	case 4:
		interval = 4 * 3600
	case 5:
		interval = 12 * 3600
	case 6:
		interval = 24 * 3600 * 10
	default:
		interval = 0
	}
	// 更新播种时间即可 TODO
	var newSowTime = ""
	timeL, err := time.ParseInLocation("2006-01-02 15:04:05", pot.SowTime, time.Local)
	if err != nil {
		return nil, err
	}
	td, err := time.ParseDuration(strconv.Itoa(int(interval)) + "s")
	if err != nil {
		return nil, err
	}

	newSowTime = timeL.Add(-td).Format("2006-01-02 15:04:05")
	// 经过的时间 总的秒数
	pot.SowTime = newSowTime
	tx := global.MYSQL.Begin()
	if _, err = pot.ComputeImportantParams(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 该道具数量 -- TODO
	if err = db.DieReduce(die.PropNum-1, pot.GardenID, die.PropID, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result["result"] = fmt.Sprintf("操作成功")
	return result, nil
}

// HarvestFlower  收获 TODO
func HarvestFlower(req global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})
	result["total"] = 0
	tx := global.MYSQL.Begin()

	// 查询出 garden0
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	user, err := db.GetUsersByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 1.判断是不是vip
	if req.IsVip == 2 {

		if user.IsVip == 1 {
			result["result"] = fmt.Sprintf("您还不是vip,无法享受vip特权哦,请正确操作")
			return result, nil
		}
		// 一键收获
		// 查询所有可以收获的花盆
		pots, err := db.GetGardenFlowerpotCanHavrest(req.GardenID)
		if err != nil {
			return nil, err
		}
		resMsgV2 := make([]string, 0)
		for _, pot := range pots {
			result, err = handleHarvestFlower(&pot, req, result, garden, user, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			// 拼接 array  这里有问题 TODO
			fmt.Println(result["result"].([]string))
			resMsgV2 = append(resMsgV2, result["result"].([]string)...)
		}
		// 全部操作完成后 TODO
		tx.Commit()
		result["result"] = resMsgV2
		return result, nil
	}
	// 操作一个花盆
	pot, err := db.GetGardenFlowerpotByID(req.GardenID, req.FlowerpotID)
	if err != nil {
		return nil, err
	}
	result, err = handleHarvestFlower(pot, req, result, garden, user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, nil

}

func handleHarvestFlower(pot *db.GardenFlowerpot, req global.GardenPotParams, result map[string]interface{}, garden *db.Garden, user *db.Users, tx *gorm.DB) (map[string]interface{}, error) {
	resMsg := make([]string, 0)
	// 花盆
	if pot.IsHarvest == 1 {
		result["result"] = fmt.Sprintf("当前花盆还未开花,请正确操作")
		result["result2"] = false
		return result, nil
	}
	// 种子
	seed, _ := db.GetGardenSeedsByID(pot.SeedID)

	// 收割
	resMsg = append(resMsg, fmt.Sprintf("收获成功,经验值+2，GB+%v", seed.RawPrice*3))
	garden.GTotaltEx += 2
	garden.GCurrentEx += 2
	// 1.1 是否点亮图谱
	gardenHouse, err := db.IsLightupAtlas(pot.SeedResult, req.GardenID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 没有点亮图谱 TODO 计算经验值 GB
			newEx := seed.LevelNum * 30
			garden.GTotaltEx += newEx
			garden.GCurrentEx += newEx
			// 新增一条item
			if gardenHouse, err = db.AddGardenFlowerHouse(&db.GardenFlowerHouse{
				GardenID: req.GardenID,
				AtlasID:  pot.SeedResult,
				Cate:     1,
				Num:      pot.FlowerNum + pot.FlowerNumHandle,
			}, tx); err != nil {
				return nil, err
			}
			// 更新花之图谱总数
			garden.GAtlas++
			// 恭喜您种出了一种新的花朵！
			// 收获成功！经验值+10，GB+2。
			// 您一共收获了4朵黄色清辉月韵！
			resMsg = append(resMsg, fmt.Sprintf("恭喜您种出了一种新的花朵[%v]!!!经验值+%v", pot.SeedResultStr, newEx))
		} else {
			fmt.Println("dsadl,.......")
			return nil, err
		}
	} else {
		// 1.2 更新花房的数量 获得的经验 gb
		fmt.Println("pot::", pot.FlowerNum, pot.FlowerNumHandle)
		gardenHouse.Num = pot.FlowerNum + pot.FlowerNumHandle + gardenHouse.Num
		if _, err = db.UpdateGardenFlowerHouse(gardenHouse, tx); err != nil {
			return nil, err
		}
		// fmt.Println(gardenHouse.Num)
	}
	resMsg = append(resMsg, fmt.Sprintf("您一共收获了%v朵%v！", pot.FlowerNum+pot.FlowerNumHandle, pot.SeedResultStr))
	user.GBMoney += seed.RawPrice * 3
	// 1.3 移除花盆
	if err = db.UpdateGardenFlowerpotWithRemove(pot, tx); err != nil {
		return nil, err
	}
	// 最后更新花园 更新用户 花房

	if _, err = db.UpdateGarden(garden, tx); err != nil {
		return nil, err
	}
	if _, err = db.UpdateUsers(user, tx); err != nil {
		return nil, err
	}
	// // 又是这个问题 事物还没提交 这个只是存在某个队列中 无法查到结果
	// if _, err = db.UpdateGardenFlowerHouse(gardenHouse, tx); err != nil {
	// 	return nil, err
	// }
	result["result"] = resMsg
	result["result2"] = true
	if seed.RawPrice*3 > 0 {
		go db.SaveGbHistory(seed.RawPrice*3, "收获花朵", "", garden)
	}
	return result, nil
}

//

// ListGardenGbDetail  查看gb获得历史记录
func ListGardenGbDetail(b *db.GardenGbDetail) (*db.DataStore, error) {
	list, err := db.ListGardenGbDetail(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenGbDetail(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + b.PageSize - 1) / b.PageSize}, nil
}

// BuyShopSeed todo
func BuyShopSeed(req *global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})
	// 查询出 garden
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 user
	user, err := db.GetUsersByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 seed
	seed, err := db.GetGardenSeedsByID(req.SeedID)
	if err != nil {
		return nil, err
	}
	// 判断种子是否可购买
	if seed.Rarity > 0 {
		result["result"] = "您无法购买该种子"
		return result, nil
	}
	// 计算数量 价格
	totalMoney := seed.RawPrice * req.SeedNum
	if user.IsVip == 2 {
		totalMoney = totalMoney / 5 * 4
	}
	if totalMoney > user.GBMoney {
		result["result"] = "您没有足够的GB哦😯"
		return result, nil
	}
	tx := global.MYSQL.Begin()
	// 更新背包
	kn, err := db.IsExistGardenFlowerKnapsackSeed(req.GardenID, req.SeedID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// new item
			if err = db.AddGardenFlowerKnapsack(&db.GardenFlowerKnapsack{
				Cate:     1,
				GardenID: req.GardenID,
				SeedID:   req.SeedID,
				SeedNum:  req.SeedNum,
			}, tx); err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			return nil, err // 这里竟然被返回了
		}
	} else if err == nil {
		fmt.Println("33")
		// update num
		kn.SeedNum += req.SeedNum
		if err = db.UpdateGardenFlowerKnapsack(kn, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 更新gb数量
	user.GBMoney -= totalMoney
	if _, err = db.UpdateUsersGB(user, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	// 存入gb历史
	go db.SaveGbHistory(totalMoney, "购买花种", fmt.Sprintf("购买了%v颗%v,花费%vGB", req.SeedNum, seed.SeedName, totalMoney), garden)
	result["result"] = fmt.Sprintf("购买成功,您获得%v颗%v,花费%vGB", req.SeedNum, seed.SeedName, totalMoney)
	// 提交事物
	tx.Commit()
	fmt.Println(result)
	return result, nil
}

// BuyShopProp 购买道具
func BuyShopProp(req *global.GardenPotParams) (interface{}, error) {
	result := make(map[string]interface{})
	// 查询出 garden
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 user
	user, err := db.GetUsersByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 prop
	prop, err := db.GetGardenPropsByID(req.PropID)
	if err != nil {
		return nil, err
	}
	// 判断种子是否可购买
	if prop.ISbuy != 2 {
		result["result"] = "您无法购买该道具"
		return result, nil
	}
	// 计算数量 价格
	totalMoney := prop.PPrice * req.PropNum
	if totalMoney > user.GBMoney {
		result["result"] = "您没有足够的GB哦😯"
		return result, nil
	}
	tx := global.MYSQL.Begin()
	// 更新背包
	kn, err := db.IsExistGardenFlowerKnapsackProp(req.GardenID, req.PropID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// new item
			if err = db.AddGardenFlowerKnapsack(&db.GardenFlowerKnapsack{
				Cate:     1,
				GardenID: req.GardenID,
				PropID:   req.PropID,
				PropNum:  req.PropNum,
			}, tx); err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			return nil, err // 这里竟然被返回了
		}
	} else if err == nil {
		// update num
		kn.PropNum += req.PropNum
		if err = db.UpdateGardenFlowerKnapsack(kn, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 更新gb数量
	user.GBMoney -= totalMoney
	if _, err = db.UpdateUsersGB(user, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	// 存入gb历史
	go db.SaveGbHistory(totalMoney, "购买道具", fmt.Sprintf("购买了%v个%v,花费%vGB", req.SeedNum, prop.PName, totalMoney), garden)
	result["result"] = fmt.Sprintf("购买成功,您获得%v个%v,花费%vGB", req.SeedNum, prop.PName, totalMoney)
	// 提交事物
	tx.Commit()
	return result, nil
}

// ListGardenMagician  查询魔法屋列表
func ListGardenMagician(b *db.GardenSeeds) (*db.DataStore, error) {
	list, err := db.ListGardenMagician(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenMagician(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + global.PageSize - 1) / global.PageSize}, nil
}

// GardenMagicianDetail  合成条件
func GardenMagicianDetail(b *global.MagicianParams) ([]db.MagicianSeedSynthesisMethods, error) {
	list, err := db.GardenMagicianDetail(b.SeedID, b.GardenID)
	return list, err
}

// GardenMagicianSynthesis  合成
func GardenMagicianSynthesis(req *global.MagicianParams) (interface{}, error) {
	result := make(map[string]interface{})
	tmpArr := make([]string, 0)
	//  查询合成条件
	items, err := GardenMagicianDetail(req)
	if err != nil {
		return nil, err
	}
	length := len(items)
	tmpLength := 0
	for _, item := range items {
		if item.TotalNum >= item.Num {
			tmpLength++
		}
	}
	if length != tmpLength {
		// 不满足达到合成条件
		tmpArr = append(tmpArr, "不满足合成条件")
		result["result"] = tmpArr
		return result, nil
	}
	// 满足合成条件 ------------
	// 查询出 garden
	garden, err := db.GetGardenByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 user
	user, err := db.GetUsersByID(req.GardenID)
	if err != nil {
		return nil, err
	}
	// 查询出 seed
	seed, err := db.GetGardenSeedsByID(req.SeedID)
	if err != nil {
		return nil, err
	}
	tx := global.MYSQL.Begin()
	// 新增背包一个种子 或者更新数量
	// 更新背包
	kn, err := db.IsExistGardenFlowerKnapsackSeed(req.GardenID, req.SeedID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// new item
			if err = db.AddGardenFlowerKnapsack(&db.GardenFlowerKnapsack{
				Cate:     1,
				GardenID: req.GardenID,
				SeedID:   req.SeedID,
				SeedNum:  1,
			}, tx); err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			return nil, err
		}
	} else if err == nil {
		// update num
		kn.SeedNum++
		if err = db.UpdateGardenFlowerKnapsack(kn, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// 减去花房中的花朵数量
	for _, val := range items {
		if err = db.UpdateGardenFlowerHouseNumber(val.TotalNum-val.Num, val.GardenID, val.AtlasID, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// 计算合成经验
	exp := seed.LevelNum * 10
	garden.GTotaltEx += exp
	garden.GCurrentEx += exp
	if garden, err = db.UpdateGarden(garden, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	// 给点GB吧
	gb := seed.LevelNum * 100
	user.GBMoney += gb
	if _, err = db.UpdateUsers(user, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	go db.SaveGbHistory(gb, "合成花朵", fmt.Sprintf("合成了%s", seed.SeedName), garden)
	tx.Commit()

	tmpArr = append(tmpArr, fmt.Sprintf("合成成功,您获得了一颗%s", seed.SeedName))
	tmpArr = append(tmpArr, fmt.Sprintf("您获得%v经验", exp))
	tmpArr = append(tmpArr, fmt.Sprintf("您获得%vGB", gb))
	result["result"] = tmpArr
	return result, nil
}

// GardenHouseList  花房花朵分页查询
func GardenHouseList(b *global.MagicianParams) (*db.DataStore, error) {
	list, err := db.ListGardenFlowerHouse(b.GardenID, b.Cate, b.Page)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenFlowerHouse(b.GardenID, b.Cate)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + global.PageSize - 1) / global.PageSize}, nil
}

// GardenHouseStatistics  花房花朵分页查询
func GardenHouseStatistics(req *global.MagicianParams) (interface{}, error) {
	return db.GardenHouseStatistics(req.GardenID, req.Cate)
}
