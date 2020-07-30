package service

/*
date:2020-07-30 20:00:37
*/

import (
	"magic/db"
	"magic/global"
)

// InitGarden 初始化花园
func InitGarden(req *global.UserAddGamesParams) error {
	user, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return err
	}
	garden := &db.Garden{
		GName:      user.Nickname + "的花园",
		GInfo:      "劳动光荣,偷窃可耻!",
		GLevel:     1,
		GAtlas:     "0",
		IsSignin:   0,
		SignDays:   0,
		GCurrentEx: 0,
	}
	// 初始化花盆

	// 初始化花房 花瓶

	// 初始化背包 道具 种子

	//
	return db.AddGarden(garden)
}

// UpdateGarden update
func UpdateGarden(b *db.Garden) error {
	return db.UpdateGarden(b)
}

// GetGardenByID get by id
func GetGardenByID(id int) (*db.Garden, error) {
	return db.GetGardenByID(id)
}
