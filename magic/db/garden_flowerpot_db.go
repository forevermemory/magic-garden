package db

import (
	"errors"
	"fmt"
	"magic/global"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

/*
CREATE TABLE `garden_flowerpot` (
  `_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `garden_id` int(11) DEFAULT NULL COMMENT '花园id和用户id相同',
  `number` int(11) DEFAULT NULL COMMENT '花盆编号 1-10',
  `is_lock` int(11) DEFAULT '0' COMMENT '是否解锁该花盆 1 未解锁 2 已经解锁',
  `is_sow` int(11) DEFAULT NULL COMMENT '空还是播种 1 空 2 播种',
  `seed_id` int(11) DEFAULT NULL COMMENT '种的种子',
  `status` int(11) DEFAULT 1 COMMENT '1 正常 2 干旱 3 有虫 4 有草',
  `seed_result` int(11)   COMMENT '种子开花结果 开出的花是图谱id',
  `seed_result_str` varchar(255)   COMMENT '种子开花结果 开出的花',
  `flower_num` int(11)  COMMENT '最后成花数量 珍惜的只会一朵 其它的0种子 不处理开原数量 浇水或者除虫、除草会翻倍产量',
  `flower_num_haldle` int(11)  COMMENT '每次浇水除草就会多一朵',
  `current_stage` int(11) DEFAULT NULL COMMENT '花的成长阶段  花种期,花苗期,花蕾期',
  `next_stage` varchar(255) DEFAULT NULL COMMENT '花的成长阶段 1 花苗 2 花蕾 3 开花',
  `next_stage_str` varchar(255) DEFAULT NULL COMMENT '成长剩余时间',
  `sow_time` varchar(255) DEFAULT NULL COMMENT '播种时间',
  `is_change_color` int(11) DEFAULT 1 COMMENT '是否可以染色 1 不可以  2 可以',
  `is_use_dye` int(11) DEFAULT 1 COMMENT '是否使用染色剂 1 未使用 2使用了',
  `change_result` varchar(255) DEFAULT NULL COMMENT '染色结果 string',
  `is_harvest` int(11) DEFAULT 1 COMMENT '是否可以收获 1不可以 2 可以',
  `disaster` int(11) DEFAULT 1 COMMENT '自然灾害类型 1健康 2干旱(浇水) 3有虫(除虫) 4有草(除草) ',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='花盆'
*/

// GardenFlowerpot 花盆
type GardenFlowerpot struct {
	ID              int    `gorm:"column:_id" json:"_id" form:"_id"`
	UserID          string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	GardenID        string `gorm:"column:garden_id" json:"garden_id" form:"garden_id"`
	Number          int    `gorm:"column:number" json:"number" form:"number"`
	IsLock          int    `gorm:"column:is_lock" json:"is_lock" form:"is_lock"`
	IsSow           int    `gorm:"column:is_sow" json:"is_sow" form:"is_sow"`
	SeedID          int    `gorm:"column:seed_id" json:"seed_id" form:"seed_id"`
	Status          int    `gorm:"column:status" json:"status" form:"status"`
	SeedResult      int    `gorm:"column:seed_result" json:"seed_result" form:"seed_result"`
	SeedResultStr   string `gorm:"column:seed_result_str" json:"seed_result_str" form:"seed_result_str"`
	FlowerNum       int    `gorm:"column:flower_num" json:"flower_num" form:"flower_num"`
	FlowerNumHandle int    `gorm:"column:flower_num_haldle" json:"flower_num_haldle" form:"flower_num_haldle"`
	CurrentStage    string `gorm:"column:current_stage" json:"current_stage" form:"current_stage"`
	NextStage       string `gorm:"column:next_stage" json:"next_stage" form:"next_stage"`
	NextStageStr    string `gorm:"column:next_stage_str" json:"next_stage_str" form:"next_stage_str"`
	SowTime         string `gorm:"column:sow_time" json:"sow_time" form:"sow_time"`
	IsChangeColor   int    `gorm:"column:is_change_color" json:"is_change_color" form:"is_change_color"`
	IsUseDye        int    `gorm:"column:is_use_dye" json:"is_use_dye" form:"is_use_dye"`
	ChangeResult    string `gorm:"column:change_result" json:"change_result" form:"change_result"`
	IsHarvest       int    `gorm:"column:is_harvest" json:"is_harvest" form:"is_harvest"`
	Disaster        int    `gorm:"column:disaster" json:"disaster" form:"disaster"`
}

// TableName 表名
func (o *GardenFlowerpot) TableName() string {
	return "garden_flowerpot"
}

// SetGFDisaster 设置自然灾害
func (o *GardenFlowerpot) SetGFDisaster(tx ...*gorm.DB) (*GardenFlowerpot, error) {
	rand.Seed(time.Now().Unix())
	rx := rand.Intn(3) + 2 // [0,3) -- [2,5) 234
	o.Disaster = rx
	return UpdateGardenFlowerpot(o, tx...)
	// return nil, nil
}

// ComputeImportantParams 计算相关参数
func (o *GardenFlowerpot) ComputeImportantParams(tx ...*gorm.DB) (*GardenFlowerpot, error) {
	// 计算花期
	// 已经成熟
	if o.IsHarvest == 2 {
		// 可以收获
		return o, nil
	}
	var index int
	// 查询出种子
	seed, err := GetGardenSeedsByID(o.SeedID)
	if err != nil {
		return nil, err
	}
	// 根据播种时间来计算各个阶段
	// fmt.Println(10 / 3)  3
	// fmt.Println(10 % 3)  1
	stage1 := seed.ForecastTime / 3 // 一阶段需要的时间 小时
	stage2 := seed.ForecastTime / 3 * 2
	// stage3 := seed.ForecastTime - seed.ForecastTime/3*2

	timeL, err := time.ParseInLocation("2006-01-02 15:04:05", o.SowTime, time.Local)
	if err != nil {
		return nil, err
	}
	cha := time.Now().Unix() - timeL.Unix()
	td, err := time.ParseDuration(strconv.Itoa(int(cha)) + "s")
	if err != nil {
		return nil, err
	}
	// 经过的时间 总的秒数
	totals := int(td.Seconds())

	// 花种期,花苗期,花蕾期 字符串展示
	nextStageStr := ""
	currentStage := ""
	nextStage := ""
	//  TODO 剩余时间不对 计算剩余时间
	// h := int(td.Hours())
	// m := int(td.Minutes()) - h*60
	// s := int(td.Seconds()) - h*60*60 - m*60
	var ph int
	var pm int
	var ps int
	var pass int

	if totals < stage1*3600 {
		nextStageStr = "进入花苗期"
		nextStage = "花苗期"
		currentStage = "花种期"
		pass = stage1*3600 - totals // 距离下一阶段还剩多少秒

	} else if totals < (stage2+stage1)*3600 && totals > stage1*3600 {
		nextStageStr = "进入花蕾期"
		nextStage = "花蕾期"
		currentStage = "花苗期"

		pass = (stage2+stage1)*3600 - totals // 距离下一阶段还剩多少秒

	} else if totals < seed.ForecastTime*3600 && totals > (stage1+stage2)*3600 {
		nextStageStr = "开花"
		nextStage = "开花"
		currentStage = "花蕾期"
		pass = seed.ForecastTime*3600 - totals // 距离下一阶段还剩多少秒

		// fmt.Println("ph,pm,ps", ph, pm, ps)
	} else if totals > seed.ForecastTime*3600 {
		// 已经成花
		if o.IsHarvest == 1 {
			// 对染色后的花盆不作处理 只对没有染色的花盆作处理
			o.IsHarvest = 2      // 都设置为可以收获
			if o.IsUseDye == 1 { // 未使用染色剂
				atlas, err := GetGardenAtlasBySeedID(o.SeedID)
				if err != nil {
					return nil, err
				}
				// result := 0
				// resultStr := ""
				if len(atlas) == 0 {
					return nil, errors.New("未知错误")
				}
				// 随机取一种颜色
				rand.Seed(time.Now().UnixNano())
				index = rand.Intn(len(atlas)) // [0,n)
				// result = atlas[rx].ID
				fmt.Println("rx---", index)
				o.FlowerNum = seed.ForecastNum
				if seed.ForecastNum > 1 {
					o.FlowerNum = o.FlowerNumHandle + o.FlowerNumHandle
				}
				o.SeedResult = atlas[index].ID // 图谱id
				o.SeedResultStr = atlas[index].FlowerCateName
			}
		}
	}
	// 当前阶段
	o.CurrentStage = currentStage
	o.NextStage = nextStage
	ph = pass / 3600
	pm = (pass - 3600*ph) / 60
	ps = pass - ph*3600 - pm*60

	if totals < seed.ForecastTime*3600 {
		// 说明还没有开花
		if ph > 0 {
			o.NextStageStr = fmt.Sprintf("%v小时%v分钟后%s", ph, pm, nextStageStr)
		} else if ph == 0 {
			if pm == 0 {
				o.NextStageStr = fmt.Sprintf("%v秒后%s", ps, nextStageStr)
			} else {
				o.NextStageStr = fmt.Sprintf("%v分钟后%s", pm, nextStageStr)
			}
		}
	}
	// 更新到数据库
	o, _ = UpdateGardenFlowerpot(o, tx...)
	return o, nil

}

// AddGardenFlowerpot 新增
func AddGardenFlowerpot(o *GardenFlowerpot) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// GetGardenFlowerpotCanHavrest 查询可收获的花盆
func GetGardenFlowerpotCanHavrest(gardenID string) ([]GardenFlowerpot, error) {
	db := global.MYSQL
	var res []GardenFlowerpot
	err := db.Table("garden_flowerpot").Where("garden_id = ? and is_lock = 2 and is_harvest = 2", gardenID).Scan(&res).Error
	// res.ComputeImportantParams()
	return res, err
}

// GetGardenFlowerpotByID get one
func GetGardenFlowerpotByID(gardenID string, flowerpotID int) (*GardenFlowerpot, error) {
	db := global.MYSQL
	var res GardenFlowerpot
	err := db.Table("garden_flowerpot").Where("garden_id = ? and number = ? and is_lock = 2 and is_sow = 2", gardenID, flowerpotID).First(&res).Error
	res.ComputeImportantParams()
	return &res, err
}

// GetGardenFlowerpotByIDIsCanSow 是否花盆可播种 返回一个盆
func GetGardenFlowerpotByIDIsCanSow(gardenID string, flowerpotID int) (*GardenFlowerpot, error) {
	db := global.MYSQL
	var res GardenFlowerpot
	err := db.Table("garden_flowerpot").Where("garden_id = ? and number = ? and is_lock = 2 and is_sow = 1", gardenID, flowerpotID).First(&res).Error
	return &res, err
}

// GetCanSowGardenFlowerpots 是否花盆可播种 返回多个盆
func GetCanSowGardenFlowerpots(gardenID string) ([]*GardenFlowerpot, error) {
	db := global.MYSQL
	var res []*GardenFlowerpot
	err := db.Table("garden_flowerpot").Where("garden_id = ? and is_lock = 2 and is_sow = 1", gardenID).Scan(&res).Error
	return res, err
}

// UpdateGardenFlowerpot 修改
func UpdateGardenFlowerpot(o *GardenFlowerpot, tx ...*gorm.DB) (*GardenFlowerpot, error) {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	err := db.Table("garden_flowerpot").Where("_id = ?", o.ID).Update(o).First(o).Error
	return o, err
}

// UpdateGardenFlowerpotWithRemove 移除花朵更新参数
func UpdateGardenFlowerpotWithRemove(pot *GardenFlowerpot, tx ...*gorm.DB) error {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	// pot.IsSow = 1
	// pot.IsHarvest = 1
	sql := "update garden_flowerpot set seed_id=0,is_sow = 1,seed_result_str = '',flower_num = 0,flower_num_haldle = 0,current_stage='',next_stage = '',next_stage_str = '',sow_time = '',change_result='',is_change_color=1,is_use_dye = 1,is_harvest=1 where _id = ?"
	err := db.Exec(sql, pot.ID).Error
	return err
}

// ListGardenFlowerpot 查询某个花园的花盆列表
func ListGardenFlowerpot(gardenid string) ([]*GardenFlowerpot, error) {
	res := make([]*GardenFlowerpot, 0)
	resnew := make([]*GardenFlowerpot, 0)
	db := global.MYSQL
	err := db.Table("garden_flowerpot").Where(" garden_id = ? and is_lock = 2", gardenid).Scan(&res).Error
	for _, val := range res {
		o, err := val.ComputeImportantParams()
		if err != nil {
			fmt.Println(err, "---")
		}
		resnew = append(resnew, o)
	}

	return resnew, err
}

// AllGardenFlowerpot 查询花园的花盆列表
func AllGardenFlowerpot() (int, error) {
	res := make([]*GardenFlowerpot, 0)
	tx := global.MYSQL.Begin()
	err := tx.Table("garden_flowerpot").Where(" is_lock = 2 and is_harvest = 1 and is_sow = 2").Scan(&res).Error
	for _, val := range res {
		_, err := val.ComputeImportantParams(tx)
		if err != nil {
			fmt.Println(err, "---")
		}
	}
	tx.Commit()
	return len(res), err
}

// SetGardenFlowerpotDisaster 查询花园的花盆列表 并设置自然灾害
func SetGardenFlowerpotDisaster() (int, error) {
	res := make([]*GardenFlowerpot, 0)
	tx := global.MYSQL.Begin()
	err := tx.Table("garden_flowerpot").Where(" is_lock = 2 and is_harvest = 1 and is_sow = 2").Scan(&res).Error
	for _, val := range res {
		_, err := val.SetGFDisaster(tx)
		if err != nil {
			return 0, err
		}
	}
	tx.Commit()
	return len(res), err
}
