package service

import (
	"magic/db"
	"magic/global"
	utils "magic/utils"
)

// UserAddGames x
func UserAddGames(req *global.UserAddGamesParams) (interface{}, error) {
	_, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return nil, err
	}
	_, err = db.GetGamesByID(req.GameID)
	if err != nil {
		return nil, err
	}
	// 先查询是否之前添加过了
	userGames, err := db.GetUserGamesByUserAndGame(req.UserID, req.GameID)
	if err != nil {
		return nil, err
	}
	if len(userGames) > 0 {
		return "您已经添加过该游戏", nil
	}

	// 1.1否则再创建一条新的记录
	err = db.AddUserGames(&db.UserGames{
		UserID:   req.UserID,
		GameID:   req.GameID,
		IsDelete: 0,
		AddTime:  utils.GetNowTimeString(),
	})
	if err != nil {
		return nil, err
	}
	return "添加成功", err
}

// UserOrderGames x
func UserOrderGames(req *global.UserAddGamesParams) (interface{}, error) {
	// user, err := db.GetUsersByID(req.UserID)
	// if err != nil {
	// 	return nil, err
	// }

	err := db.UserOrderGames(req.OrderIndex, req.UserGameID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UserGamesList x
func UserGamesList(req *global.UserAddGamesParams) (interface{}, error) {
	user, err := db.GetUsersByID(req.UserID)
	if err != nil {
		return nil, err
	}

	userGames, err := db.GetGamesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return userGames, nil
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
	err = db.DeleteUserGames(user.ID, game.ID)
	if err != nil {
		return err
	}
	return nil

}
