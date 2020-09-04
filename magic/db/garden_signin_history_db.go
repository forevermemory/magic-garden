package db

import (
	"magic/global"

	"github.com/jinzhu/gorm"
)

/*
date:2020-08-01 21:28:33
*/

// GardenSigninHistory GardenSigninHistory
type GardenSigninHistory struct {
	ID        int    `json:"id" form:"id" gorm:"column:id;primary_key;auto_increment;comment:'主键'"`
	HGardenID string `json:"h_garden_id" form:"h_garden_id" gorm:"column:h_garden_id;comment:'花园id'"`
	HUserID   string `json:"h_user_id" form:"h_user_id" gorm:"column:h_user_id;comment:'用户id'"`
	HTime     string `json:"h_time" form:"h_time" gorm:"column:h_time;comment:'签到时间'"`
	IsDelete  uint   `json:"-" form:"-" gorm:"column:is_delete;default:0"` // 0 未删除
	PageNo    int    `json:"page" form:"page" gorm:"-"`
}

// TableName 表名
func (o *GardenSigninHistory) TableName() string {
	return "garden_signin_history"
}

// AddGardenSigninHistory 新增
func AddGardenSigninHistory(o *GardenSigninHistory, tx ...*gorm.DB) (*GardenSigninHistory, error) {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	err := db.Create(o).Error
	return o, err
}

// ListGardenSigninHistory 分页条件查询
func ListGardenSigninHistory(o *GardenSigninHistory) ([]*GardenSigninHistory, error) {
	db := global.MYSQL
	res := make([]*GardenSigninHistory, 0)
	err := db.Table("garden_signin_history").Where("is_delete = 0").Where(o).Offset((o.PageNo - 1) * global.PageSize).Limit(global.PageSize).Find(&res).Error
	return res, err
}

// CountGardenSigninHistory 条件数量
func CountGardenSigninHistory(o *GardenSigninHistory) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("garden_signin_history").Where("is_delete = 0").Where(o).Count(&count).Error
	return count, err
}
