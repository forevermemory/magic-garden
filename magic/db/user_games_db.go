package db

import (
	"magic/global"
)

/*
date:2020-07-30 20:08:02
*/

// UserGames UserGames
type UserGames struct {
	ID int `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`

	UserID     int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'用户id'"`
	GameID     int    `json:"game_id" form:"game_id" gorm:"column:game_id;comment:'游戏id'"`
	IsDelete   int    `json:"is_delete" form:"is_delete" gorm:"column:is_delete;comment:'是否删了这个游戏'"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:'添加时间'"`
	DeleteTime string `json:"delete_time" form:"delete_time" gorm:"column:delete_time;comment:'删除时间'"`
}

// TableName 表名
func (o *UserGames) TableName() string {
	return "user_games"
}

// GetUserGamesByUserAndGame UserAndGame
func GetUserGamesByUserAndGame(userid int, gameid int) ([]*UserGames, error) {
	db := global.MYSQL
	var o []*UserGames
	err := db.Table("user_games").Where("user_id = ? and game_id = ?", userid, gameid).Find(&o).Error
	return o, err
}

// AddUserGames 新增
func AddUserGames(o *UserGames) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// UpdateUserGames 修改
func UpdateUserGames(o *UserGames) error {
	db := global.MYSQL
	return db.Table("user_games").Where("_id=?", o.ID).Update(o).Error
}

// UpdateUserGamesSetIsdeleteEqulaZelo 修改
func UpdateUserGamesSetIsdeleteEqulaZelo(o *UserGames) error {
	db := global.MYSQL
	return db.Exec("update user_games set is_delete = 0 where _id = ?", o.ID).Error
	// return db.Table("user_games").Where("_id=?", o.ID).Update("is_delete = 0").Error
}
