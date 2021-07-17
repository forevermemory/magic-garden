package db

import (
	"magic/global"
	"magic/utils"
)

/*
date:2020-07-30 20:08:02
*/

// UserGames UserGames
type UserGames struct {
	ID int `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`

	UserID     string `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'用户id'"`
	GameID     int    `json:"game_id" form:"game_id" gorm:"column:game_id;comment:'游戏id'"`
	IsDelete   int    `json:"is_delete" form:"is_delete" gorm:"column:is_delete;comment:'是否删了这个游戏'"`
	OrderIndex int    `json:"order_index" form:"order_index" gorm:"column:order_index;comment:'排序字段'"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:'添加时间'"`
	UpdateTime string `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间'"`
	DeleteTime string `json:"delete_time" form:"delete_time" gorm:"column:delete_time;comment:'删除时间'"`
}

// UserGamesRes UserGamesRes
type UserGamesRes struct {
	ID int `json:"_id" form:"_id" gorm:"column:_id"`

	GName      string `json:"g_name" form:"g_name" gorm:"column:g_name;comment:'游戏名称'"`
	UserGameID int    `json:"user_game_id" form:"user_game_id" gorm:"column:user_game_id;comment:'游戏名称'"`
	GameID     int    `json:"game_id" form:"game_id" gorm:"column:game_id;comment:'游戏名称'"`
	GState     int    `json:"g_state" form:"g_state" gorm:"column:g_state;comment:'是否上线'"`
	GURL       string `json:"g_url" form:"g_url" gorm:"column:g_url;comment:'路由'"`
	GDesc      string `json:"g_desc" form:"g_desc" gorm:"column:g_desc;comment:'描述'"`
	OrderIndex int    `json:"order_index" form:"order_index" gorm:"column:order_index;comment:'排序字段'"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:'添加时间'"`
}

// TableName 表名
func (o *UserGames) TableName() string {
	return "user_games"
}

// GetGamesByUserID UserAndGame
func GetGamesByUserID(userid string) ([]*UserGamesRes, error) {
	db := global.MYSQL
	var o []*UserGamesRes = make([]*UserGamesRes, 0)
	sql := `
	SELECT c.*,a.add_time,c._id game_id,a._id user_game_id from 
	user_games a 
	INNER JOIN users b on b._id = a.user_id
	INNER JOIN games c on c._id = a.game_id
	
	WHERE b._id = ?
	ORDER BY a.order_index desc, a.update_time desc
	`
	err := db.Raw(sql, userid).Find(&o).Error
	return o, err
}

// GetUserGamesByUserAndGame UserAndGame
func GetUserGamesByUserAndGame(userid string, gameid int) ([]*UserGames, error) {
	db := global.MYSQL
	var o = make([]*UserGames, 0)
	err := db.Table("user_games").Where("user_id = ? and game_id = ? and is_delete = 0", userid, gameid).Find(&o).Error
	return o, err
}

// AddUserGames 新增
func AddUserGames(o *UserGames) error {
	db := global.MYSQL
	o.UpdateTime = o.AddTime
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

// DeleteUserGames 修改
func DeleteUserGames(userid string, gameid int) error {
	db := global.MYSQL
	return db.Exec("delete from  user_games  where user_id = ? and game_id = ?", userid, gameid).Error
}

// UserOrderGames x
func UserOrderGames(index int, userGameId int) error {
	db := global.MYSQL
	return db.Debug().Exec("update   user_games  set order_index = ? ,update_time = ? where _id = ?", index, utils.GetNowTimeString(), userGameId).Error
}
