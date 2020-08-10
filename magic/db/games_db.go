package db

import "magic/global"

/*
date:2020-07-30 20:20:24
*/

// Games Games
type Games struct {
	ID int `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`

	GName  string `json:"g_name" form:"g_name" gorm:"column:g_name;comment:'游戏名称'"`
	GState int    `json:"g_state" form:"g_state" gorm:"column:g_state;comment:'是否上线'"`

	PageNo   int `json:"page" form:"page" gorm:"-"`
	PageSize int `json:"page_size" form:"page_size" gorm:" - "`
}

// TableName 表名
func (o *Games) TableName() string {
	return "games"
}

// DeleteGames 根据id删除
func DeleteGames(id int) error {
	db := global.MYSQL
	return db.Table("games").Where("_id = ?", id).Update("g_state", 1).Error
}

// GetGamesByID 根据id查询一个
func GetGamesByID(id int) (*Games, error) {
	db := global.MYSQL
	o := &Games{}
	err := db.Table("games").Where("_id = ?", id).First(o).Error
	return o, err
}

// AddGames 新增
func AddGames(o *Games) (*Games, error) {
	db := global.MYSQL
	err := db.Create(o).Error
	return o, err
}

// UpdateGames 修改
func UpdateGames(o *Games) (*Games, error) {
	db := global.MYSQL
	err := db.Table("games").Where("_id=?", o.ID).Update(o).Error
	return o, err
}

// ListGames 分页条件查询
func ListGames(o *Games) ([]*Games, error) {
	db := global.MYSQL
	res := make([]*Games, 0)
	err := db.Table("games").Where(o).Offset((o.PageNo - 1) * o.PageSize).Limit(o.PageSize).Find(&res).Error
	return res, err
}

// CountGames 条件数量
func CountGames(o *Games) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("games").Where(o).Count(&count).Error
	return count, err
}
