package db

import (
	"magic/global"
)

/*
date:2020-07-29 14:00:32
*/

/*
CREATE TABLE `users` (
  `_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL COMMENT '用户名(登录账号)',
  `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `is_vip` int(11) DEFAULT '0' COMMENT '是否是vip 1否 2是',
  `change_pass_time` varchar(255) DEFAULT NULL COMMENT '上次修改密码时间',
  `desc` varchar(255) DEFAULT NULL COMMENT '备注',
  `gb_money` varchar(1024) DEFAULT '10000' COMMENT 'GB',
  `yuanbao` varchar(255) DEFAULT '0' COMMENT '元宝',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'
*/

// Users Users
type Users struct {
	ID             int    `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`
	Username       string `json:"username" form:"username" gorm:"column:username;comment:'用户名'"`
	Nickname       string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:'昵称'"`
	Password       string `json:"password" form:"password" gorm:"column:password;comment:'密码'"`
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:'手机'"`
	Status         int    `json:"status" form:"status" gorm:"column:status;comment:'状态'"`
	IsVip          int    `json:"is_vip" form:"is_vip" gorm:"column:is_vip;comment:'是否是会员'"`
	Desc           string `json:"desc" form:"desc" gorm:"column:desc;comment:'desc'"`
	GBMoney        string `json:"gb_money" form:"gb_money" gorm:"column:gb_money;comment:'FGB'"`
	Yuanbao        string `json:"yuanbao" form:"yuanbao" gorm:"column:yuanbao;comment:'重置的元宝数量'"`
	ChangePassTime string `json:"change_pass_time" form:"change_pass_time" gorm:"column:change_pass_time;comment:'desc'"`
	PageNo         int    `json:"-" form:"page" gorm:"-"`
	PageSize       int    `json:"-" form:"page_size" gorm:" - "`
}

// TableName 表名
func (o *Users) TableName() string {
	return "users"
}

// DeleteUsers 根据id删除
func DeleteUsers(id int) error {
	return nil
	// return global.MYSQL.Table("users").Where("id = ?", id).Update("IS_DELETE", 1).Error
}

// GetUsersByUsernameAndPassword login
func GetUsersByUsernameAndPassword(username string, password string) (*Users, error) {
	o := &Users{}
	err := global.MYSQL.Table("users").Where("username = ? and password = ?", username, password).First(o).Error
	return o, err
}

// GetUsersByID 根据id查询一个
func GetUsersByID(id int) (*Users, error) {
	o := &Users{}
	err := global.MYSQL.Table("users").Where("_id = ?", id).First(o).Error
	return o, err
}

// GetUserByUsername GetUserByUsername
func GetUserByUsername(username string) (*Users, error) {
	o := &Users{}
	err := global.MYSQL.Table("users").Where("username = ?", username).Limit(1).First(o).Error
	return o, err
}

// UsersLoginByUsernameAndPassword 登陆
func UsersLoginByUsernameAndPassword(o *Users) (*Users, error) {
	newuser := &Users{}
	err := global.MYSQL.Table("users").Where(o).First(newuser).Error
	return newuser, err
}

// AddUsers 新增
func AddUsers(o *Users) error {
	return global.MYSQL.Create(o).Error
}

// UpdateUsers 修改
func UpdateUsers(o *Users) (*Users, error) {
	err := global.MYSQL.Table("users").Where("_id=?", o.ID).Update(o).Error
	return o, err
}

// ListUsers 分页条件查询
func ListUsers(o *Users) ([]*Users, error) {
	res := make([]*Users, 0)
	err := global.MYSQL.Table("users").Where(o).Offset((o.PageNo - 1) * o.PageSize).Limit(o.PageSize).Find(&res).Error
	return res, err
}

// CountUsers 条件数量
func CountUsers(o *Users) (int64, error) {
	var count int64
	err := global.MYSQL.Table("users").Where(o).Count(&count).Error
	return count, err
}
