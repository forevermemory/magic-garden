package service

/*
date:2020-07-30 20:00:37
*/

import (
	"magic/db"
	"magic/global"

	"github.com/jinzhu/gorm"
)

// InitGarden 初始化花园
func InitGarden(req *global.UserAddGamesParams) error {
	user, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return err
	}
	// 首先需要判断是否需要重新初始化 important
	_, err = db.GetGardenByID(req.UserID)
	if err == gorm.ErrRecordNotFound {
		// 开始事物
		tx := global.MYSQL.Begin()
		// 1. 创建一个花园
		garden := &db.Garden{
			GName:      user.Username + "的花园",
			GInfo:      "劳动可耻,偷窃光荣!",
			GLevel:     1,
			GAtlas:     "0",
			IsSignin:   0,
			SignDays:   "0",
			GCurrentEx: "0",
		}
		if err = tx.Create(garden).Error; err != nil {
			tx.Rollback()
			return err
		}
		// 2.初始化10个花盆
		for i := 1; i < 11; i++ {
			// 2.1 只设置两个花盆解锁
			var islock = 1
			if i <= 2 {
				islock = 2
			}
			huapen := &db.GardenFlowerpot{
				UserID:   req.UserID,
				GardenID: req.UserID,
				Number:   i,
				IsLock:   islock,
				IsSow:    1,
			}
			if err = tx.Create(huapen).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		// 初始化花房 花瓶 好像不需要初始化哦 到时候写在收获中

		// 初始化背包 道具 种子

		//
		tx.Commit()
	} else {
		return nil
	}

	return nil
}

// ListGardenKnapsack  查询列表
func ListGardenKnapsack(b *db.GardenFlowerKnapsack) (*db.DataStore, error) {
	list, err := db.ListGardenFlowerKnapsack(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGardenFlowerKnapsack(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + global.PageSize - 1) / global.PageSize}, nil
}

// UpdateGarden update
func UpdateGarden(b *db.Garden) error {
	return db.UpdateGarden(b)
}

// GetGardenByID get by id
func GetGardenByID(id int) (*db.Garden, error) {
	return db.GetGardenByID(id)
}

// GetGardenHelpByID get by id
func GetGardenHelpByID(id int) (*db.GardenHelp, error) {
	return db.GetGardenHelpByID(id)
}

// GetGardenHelpTitles 货源帮助标题列表
func GetGardenHelpTitles() ([]*db.GardenHelp, error) {
	return db.GetGardenHelpTitles()
}
