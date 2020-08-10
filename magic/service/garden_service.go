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
	"strconv"

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
			GName: req.Gname,
			GInfo: req.Ginfo,
			// GName:      user.Username + "的花园",
			// GInfo:      "劳动可耻,偷窃光荣!",
			GLevel:     1,
			GAtlas:     "0",
			IsSignin:   0,
			SignDays:   "0",
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

// UpdateGarden update
func UpdateGarden(b *db.Garden) (*db.Garden, error) {
	return db.UpdateGarden(b)
}

// GetGardenByID get by id
func GetGardenByID(id int) (*db.Garden, error) {
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
		days, err := strconv.Atoi(garden.SignDays)
		if err != nil {
			return nil, err
		}
		if days < 7 {
			days++
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
		gmoney, err := strconv.Atoi(user.GBMoney)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		gmoney += signinMoney
		updateGmoneySQL := "update users set gb_money = ? where _id = ?"
		if err = tx.Exec(updateGmoneySQL, strconv.Itoa(gmoney), user.ID).First(user).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		// exp // 最后更新 登陆天数
		totalExp := garden.GTotaltEx + signinExp
		curExp := garden.GCurrentEx + signinExp
		updateGardenExpSQL := "update garden set g_current_ex = ?,g_total_ex = ?,is_signin = 2,sign_days = ? where _id = ?"
		if err = tx.Exec(updateGardenExpSQL, curExp, totalExp, strconv.Itoa(days), garden.ID).First(garden).Error; err != nil {
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
	pot.IsChangeColor = 1
	_, err := db.UpdateGardenFlowerpot(pot, tx)
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
	if pot.FlowerNumHandle < seed.ForecastNum*2 {
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
	pot.IsSow = 1
	pot.IsHarvest = 1
	if _, err = db.UpdateGardenFlowerpot(pot, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result["result"] = "移除成功"
	return result, nil
}
