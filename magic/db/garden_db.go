package db

import "magic/global"

/*
date:2020-07-30 20:00:37
*/

// Garden Garden
type Garden struct {
	ID int `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`

	GName      string `json:"g_name" form:"g_name" gorm:"column:g_name;comment:'花园名称'"`
	GInfo      string `json:"g_info" form:"g_info" gorm:"column:g_info;comment:'花园公告'"`
	GLevel     int    `json:"g_level" form:"g_level" gorm:"column:g_level;comment:'花园等级'"`
	IsSignin   int    `json:"is_signin" form:"is_signin" gorm:"column:is_signin;comment:'当天是否签到'"`
	SignDays   int    `json:"sign_days" form:"sign_days" gorm:"column:sign_days;comment:'累计签到天数'"`
	GAtlas     string `json:"g_atlas" form:"g_atlas" gorm:"column:g_atlas;comment:'花园图谱数量'"`
	GCurrentEx int    `json:"g_current_ex" form:"g_current_ex" gorm:"column:g_current_ex;comment:'当天获得的经验'"`
}

// TableName 表名
func (o *Garden) TableName() string {
	return "garden"
}

// GetGardenByID 根据id查询一个
func GetGardenByID(id int) (*Garden, error) {
	db := global.MYSQL
	o := &Garden{}
	err := db.Table("garden").Where("_id = ?", id).First(o).Error
	return o, err
}

// AddGarden 新增
func AddGarden(o *Garden) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// UpdateGarden 修改
func UpdateGarden(o *Garden) error {
	db := global.MYSQL
	return db.Table("garden").Where("_id=?", o.ID).Update(o).Error
}
