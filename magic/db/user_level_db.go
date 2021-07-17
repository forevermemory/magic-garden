package db

import (
	"magic/global"
	"time"

	"github.com/jinzhu/gorm"
)

/*
date:2021-07-17 14:20:14

*/

// UserLevel UserLevel
type UserLevel struct {
	ID         int    `json:"id" form:"id" gorm:"column:id;primary_key;auto_increment;comment:'主键'"`
	OrderIndex int    `json:"order_index" form:"order_index" gorm:"column:order_index;comment:'排序'"`
	LevelName  string `json:"level_name" form:"level_name" gorm:"column:level_name;comment:'等级名称'"`
	LevelImg   string `json:"level_img" form:"level_img" gorm:"column:level_img;comment:'图片'"`

	OtherDesc  string `json:"other_desc" form:"other_desc" gorm:"column:other_desc;comment:'备注'"`
	CreateTime string `json:"-" form:"-" gorm:"column:create_time;comment:'创建时间'"`
	UpdateTime string `json:"-" form:"-" gorm:"column:update_time;comment:'更新时间'"`
	IsDelete   uint   `json:"-" form:"-" gorm:"column:is_delete;default:0"` // 0 未删除

	PageNo   int `json:"-" form:"page" gorm:"-"`
	PageSize int `json:"-" form:"page_size" gorm:" - "`
}

// TableName 表名
func (o *UserLevel) TableName() string {
	return "user_level"
}

// DeleteUserLevel 根据id删除
func DeleteUserLevel(id int, tx ...*gorm.DB) error {
	db := global.MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	sql := "update user_level set is_delete = 1,update_time = ? where id = ?"
	err := db.Exec(sql, time.Now().Format("2006-01-02 15:04:05"), id).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserLevelByID 根据id查询一个
func GetUserLevelByID(id int) (*UserLevel, error) {
	db := global.MYSQL
	o := &UserLevel{}
	err := db.Table("user_level").Where("is_delete = 0").Where("id = ?", id).First(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// GetUserLevelByLevel 根据级别查询一个
func GetUserLevelByLevel(le int) (*UserLevel, error) {
	db := global.MYSQL
	o := &UserLevel{}
	err := db.Table("user_level").Where("is_delete = 0").Where("order_index = ?", le).First(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// AddUserLevel 新增
func AddUserLevel(o *UserLevel, tx ...*gorm.DB) (*UserLevel, error) {
	db := global.MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	o.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	err := db.Create(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// UpdateUserLevel 修改
func UpdateUserLevel(o *UserLevel, tx ...*gorm.DB) (*UserLevel, error) {
	db := global.MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	o.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	err := db.Table("user_level").Where("id = ?", o.ID).Update(o).First(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// ListUserLevel 分页条件查询
func ListUserLevel(o *UserLevel) ([]*UserLevel, error) {
	db := global.MYSQL
	res := make([]*UserLevel, 0)
	err := db.Table("user_level").Where("is_delete = 0").Where(o).Offset((o.PageNo - 1) * o.PageSize).Limit(o.PageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CountUserLevel 条件数量
func CountUserLevel(o *UserLevel) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("user_level").Where("is_delete = 0").Where(o).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
