package db

import (
	"magic/global"
	"magic/utils"
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
)

/*
date:2020-07-30 20:00:37

关于经验和等级?<br/>
您的经验值会从以下几方面进行计算:<br/>
1.播种:+2,★每日上限20点★<br/>
2.浇水,除草,除虫:+1,★每日上限50点★<br/>
3.收获花朵会获得经验,等级越高的花朵经验值越高<br/>
4.合成花种会获得20至200点经验,等级越高的花种经验值越高<br/>
5.点亮花谱(种出以前未种出的花)可获得50至300点经验,等级越高的花朵经验值越高<br/>
6.每级升级所需经验为:当前级别*(200点)<br/>
----------<br/>

CREATE TABLE `garden` (
  `_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '和用户id一样',
  `g_name` varchar(255) NOT NULL COMMENT '花园名称',
  `g_info` varchar(255) NOT NULL COMMENT '花园公告',
  `g_level` int(11) NOT NULL COMMENT '花园等级',
  `g_atlas` varchar(255) NOT NULL COMMENT '花园图谱数量',
  `is_signin` int(11) NOT NULL DEFAULT '1' COMMENT '1没有签到 2签到',
  `sign_days` varchar(255) NOT NULL COMMENT '累计签到天数',
  `g_current_ex` int(11) NOT NULL DEFAULT '0' COMMENT '当天获得的经验',
  `g_total_ex` int(11) NOT NULL DEFAULT '0' COMMENT '累计获得的经验',
  `g_level_str` varchar(255) DEFAULT NULL COMMENT '当前等级str',
  `g_level_cha` varchar(255) DEFAULT NULL COMMENT '距离下一个等级的经验值',
  `g_sow_exp` int(11) DEFAULT NULL COMMENT '每日播种经验',
  `g_handle_exp` int(11) DEFAULT NULL COMMENT '每日除草经验',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='花园表'
*/

// Garden Garden
type Garden struct {
	ID          string `json:"_id" form:"_id" gorm:"column:_id;primary_key;comment:'主键'"`
	GName       string `json:"g_name" form:"g_name" gorm:"column:g_name;comment:'花园名称'"`
	GInfo       string `json:"g_info" form:"g_info" gorm:"column:g_info;comment:'花园公告'"`
	GLevel      int    `json:"g_level" form:"g_level" gorm:"column:g_level;comment:'花园等级'"`
	IsSignin    int    `json:"is_signin" form:"is_signin" gorm:"column:is_signin;comment:'当天是否签到 1没有签到 2签到'"`
	GSigninTime string `json:"g_signin_time" form:"g_signin_time" gorm:"column:g_signin_time"`
	SignDays    int    `json:"sign_days" form:"sign_days" gorm:"column:sign_days;comment:'累计签到天数'"`
	GAtlas      int    `json:"g_atlas" form:"g_atlas" gorm:"column:g_atlas;comment:'花园图谱数量'"`
	GCurrentEx  int    `json:"g_current_ex" form:"g_current_ex" gorm:"column:g_current_ex;comment:'当天获得的经验'"`
	GTotaltEx   int    `json:"g_total_ex" form:"g_total_ex" gorm:"column:g_total_ex;comment:'累计获得的经验'"`
	GLevelStr   string `json:"g_level_str" form:"g_level_str" gorm:"column:g_level_str;comment:'xx'"`  // 见习魔法学徒
	GLevelCha   string `json:"g_level_cha" form:"g_level_cha"  gorm:"column:g_level_cha;comment:'xx'"` // (经验:3/200)
	GSowExp     int    `json:"g_sow_exp" form:"g_sow_exp"  gorm:"column:g_sow_exp"`
	GHandleExp  int    `json:"g_handle_exp" form:"g_handle_exp"  gorm:"column:g_handle_exp"`
}

// HandleGbUpdateInsertHistory gb变化时存入历史表
func (o *Garden) HandleGbUpdateInsertHistory(tx ...*gorm.DB) string {
	return "garden"
}

// TableName 表名
func (o *Garden) TableName() string {
	return "garden"
}

// ComputeEmpirical 有操作时候更新当日经验值
func (o *Garden) ComputeEmpirical(source int, num int, tx ...*gorm.DB) (int, error) {
	// 播种
	if source == 1 {
		tmp := o.GSowExp
		if tmp == 20 {
			return 0, nil
		}
		dtmp := num * 2 // 18 24
		if tmp+dtmp >= 20 {
			dtmp = 20 - o.GSowExp
		}
		o.GCurrentEx += dtmp
		o.GTotaltEx += dtmp
		o.GSowExp += dtmp
		if _, err := UpdateGarden(o, tx...); err != nil {
			return 0, err
		}
		return dtmp, nil
	} else if source == 2 { // 照看
		tmp := o.GHandleExp
		// 浇水除虫除草
		if tmp == 50 {
			return 0, nil
		}
		dtmp := num * 1
		if tmp+dtmp >= 50 {
			dtmp = 50 - o.GHandleExp
		}
		o.GCurrentEx += dtmp
		o.GTotaltEx += dtmp
		o.GHandleExp += dtmp
		if _, err := UpdateGarden(o, tx...); err != nil {
			return 0, err
		}
		return dtmp, nil
	}
	return 0, nil
}

// ComputeCurrentLevel 计算等级
func (o *Garden) ComputeCurrentLevel(tx ...*gorm.DB) (*Garden, error) {
	// o.GTotaltEx
	// 200 400 600 等差数列 a1 = 200 a2 = 400 an = 200n sn = n*(200 + 200n) /2 = 100n + 100n^2
	// 100n^2  + 100n - GTotaltEx = 0  n =
	// n:=
	level := int(math.Floor((math.Abs(100*(-1)-math.Sqrt(float64(10000+4*100*o.GTotaltEx))) / 200))) - 1
	sn := 100*level*level + 100*level
	cha := o.GTotaltEx - sn
	// 设置差值
	o.GLevelCha = strconv.Itoa(cha)
	o.GLevel = level
	// 更新等级 g_level_str
	gardenLevel, err := GetGardenFlowerLevelByLevelNum(level)
	if err != nil {
		return nil, err
	}
	o.GLevelStr = gardenLevel.LevelName
	if _, err = UpdateGarden(o, tx...); err != nil {
		return nil, err
	}

	return o, nil
}

// GetGardenByID 根据id查询一个
func GetGardenByID(id string) (*Garden, error) {
	db := global.MYSQL
	o := &Garden{}
	err := db.Table("garden").Where("_id = ?", id).First(o).Error
	o.ComputeCurrentLevel()
	return o, err
}

// ListGarden 查询所有花园
func ListGarden() ([]*Garden, error) {
	res := make([]*Garden, 0)
	db := global.MYSQL
	sql := "select * from garden where sign_days >1"
	err := db.Raw(sql).Scan(&res).Error
	return res, err
}

// AddGarden 新增
func AddGarden(o *Garden, tx ...*gorm.DB) error {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	return db.Create(o).Error
}

// UpdateGarden 修改
func UpdateGarden(o *Garden, tx ...*gorm.DB) (*Garden, error) {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	// sql := "update garden set g_name = ?,g_info = ? where _id = ?"
	// err := db.Exec(sql, o.GName, o.GInfo, o.ID).First(o).Error
	err := db.Table("garden").Where("_id=?", o.ID).Update(o).First(o).Error
	return o, err
}

/*
CREATE TABLE `garden_signin_reward` (
  `_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `r_day` int(11) NOT NULL COMMENT '第n天',
  `r_prop_id` int(11) DEFAULT NULL COMMENT '道具id',
  `r_prop_num` int(11) DEFAULT NULL COMMENT '道具数量',
  `r_prop_name` varchar(255) DEFAULT NULL COMMENT '道具名称',
  `r_seed_id` int(11) DEFAULT NULL COMMENT '种子id',
  `r_seed_num` int(11) DEFAULT NULL COMMENT '种子数量',
  `r_seed_name` varchar(255) DEFAULT NULL COMMENT '种子名称',
  `r_exp` int(11) NOT NULL COMMENT '获得的经验',
  `r_gb` int(11) DEFAULT NULL COMMENT '获得GB',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COMMENT='签到奖励'
*/

// GardenSigninReward 签到奖励
type GardenSigninReward struct {
	ID        int    `gorm:"column:_id" json:"_id" form:"_id"`
	RDay      int    `gorm:"column:r_day" json:"r_day" form:"r_day"`
	RPropID   int    `gorm:"column:r_prop_id" json:"r_prop_id" form:"r_prop_id"`
	RPropNum  int    `gorm:"column:r_prop_num" json:"r_prop_num" form:"r_prop_num"`
	RPropName string `gorm:"column:r_prop_name" json:"r_prop_name" form:"r_prop_name"`
	RSeedID   int    `gorm:"column:r_seed_id" json:"r_seed_id" form:"r_seed_id"`
	RSeedNum  int    `gorm:"column:r_seed_num" json:"r_seed_num" form:"r_seed_num"`
	RSeedName string `gorm:"column:r_seed_name" json:"r_seed_name" form:"r_seed_name"`
	RExp      int    `gorm:"column:r_exp" json:"r_exp" form:"r_exp"`
	RGb       int    `gorm:"column:r_gb" json:"r_gb" form:"r_gb"`
}

// TableName 表名
func (o *GardenSigninReward) TableName() string {
	return "garden_signin_reward"
}

// GetSignInRewardsSeed 查询签到奖励 种子
func GetSignInRewardsSeed(days int) ([]*GardenSigninReward, error) {
	db := global.MYSQL
	sql := "select _id,r_day,r_seed_id,r_seed_num,r_seed_name,r_exp,r_gb from garden_signin_reward where r_prop_id is null and r_day = ?"
	var res []*GardenSigninReward
	err := db.Raw(sql, days).Scan(&res).Error
	return res, err
}

// GetSignInRewardsProp 查询签到奖励 道具
func GetSignInRewardsProp(days int) ([]*GardenSigninReward, error) {
	db := global.MYSQL
	sql := "select _id,r_day,r_prop_id,r_prop_num,r_prop_name,r_exp,r_gb from garden_signin_reward where r_seed_id is null and r_day = ?"
	var res []*GardenSigninReward
	err := db.Raw(sql, days).Scan(&res).Error
	return res, err
}

// -----------gb  历史

// GardenGbDetail GardenGbDetail
type GardenGbDetail struct {
	ID       int    `json:"id" form:"id" gorm:"column:id;primary_key;auto_increment;comment:'主键'"`
	GbNum    int    `json:"gb_num" form:"gb_num" gorm:"column:gb_num;comment:'本次获得的数量'"`
	GardenID string `json:"g_garden_id" form:"g_garden_id" gorm:"column:g_garden_id"`
	GbTime   string `json:"gb_time" form:"gb_time" gorm:"column:gb_time;comment:'时间'"`
	GbSource string `json:"gb_source" form:"gb_source" gorm:"column:gb_source;comment:'获得来源'"`
	GbDetail string `json:"gb_detail" form:"gb_detail" gorm:"column:gb_detail;comment:'描述'"`
	GbTotal  int    `json:"gb_total" form:"gb_total" gorm:"column:gb_total;comment:'剩余总数'"`
	PageNo   int    `json:"page" form:"page" gorm:"-"`
	PageSize int    `json:"page_size" form:"page_size" gorm:"-"`
}

// TableName 表名
func (o *GardenGbDetail) TableName() string {
	return "garden_gb_detail"
}

// SaveGbHistory SaveGbHistory
func SaveGbHistory(num int, source, detail string, garden *Garden) {
	// global.MYSQL
	user, _ := GetUsersByID(garden.ID)
	o := &GardenGbDetail{
		GbNum:    num,
		GardenID: garden.ID,
		GbTime:   utils.GetNowTimeString(),
		GbSource: source,
		GbTotal:  user.GBMoney,
		GbDetail: detail,
	}
	AddGardenGbDetail(o)

}

// AddGardenGbDetail 新增
func AddGardenGbDetail(o *GardenGbDetail, tx ...*gorm.DB) (*GardenGbDetail, error) {
	db := global.MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Create(o).Error
	return o, err
}

// ListGardenGbDetail 分页条件查询
func ListGardenGbDetail(o *GardenGbDetail) ([]*GardenGbDetail, error) {
	db := global.MYSQL
	res := make([]*GardenGbDetail, 0)
	err := db.Table("garden_gb_detail").Where(o).Offset((o.PageNo - 1) * o.PageSize).Limit(o.PageSize).Find(&res).Error
	return res, err
}

// CountGardenGbDetail 条件数量
func CountGardenGbDetail(o *GardenGbDetail) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("garden_gb_detail").Where(o).Count(&count).Error
	return count, err
}
