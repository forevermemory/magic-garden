package service

import (
	"errors"
	"magic/db"
	"magic/global"
	utils "magic/utils"
)

// UserAddGames x
func UserAddGames(req *global.UserAddGamesParams) error {
	user, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return err
	}
	game, err := db.GetGamesByID(req.GameID)
	if err != nil {
		return err
	}
	// 先查询是否之前添加过了
	userGames, err := db.GetUserGamesByUserAndGame(user.ID, game.ID)
	if err != nil {
		return err
	}
	if len(userGames) == 0 {
		// 1.1否则再创建一条新的记录
		return db.AddUserGames(&db.UserGames{
			UserID:   user.ID,
			GameID:   game.ID,
			IsDelete: 1,
			AddTime:  utils.GetNowTimeString(),
		})
	}
	// // 1.2 修改
	userGame := userGames[0]
	userGame.AddTime = utils.GetNowTimeString()
	if err = db.UpdateUserGames(userGame); err != nil {
		return err
	}
	return nil
}

// UserDeleteGames x
func UserDeleteGames(req *global.UserAddGamesParams) error {
	user, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return err
	}
	game, err := db.GetGamesByID(req.GameID)
	if err != nil {
		return err
	}
	userGames, err := db.GetUserGamesByUserAndGame(user.ID, game.ID)
	if err != nil {
		return err
	}
	if len(userGames) > 1 {
		return errors.New("未知错误")
	}
	userGame := userGames[0]
	userGame.DeleteTime = utils.GetNowTimeString()
	userGame.IsDelete = 2

	return db.UpdateUserGames(userGame)
}
