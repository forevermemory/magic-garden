package service

/*
date:2020-07-30 20:20:24
*/

import "magic/db"

// AddGames add
func AddGames(b *db.Games) (*db.Games, error) {
	return db.AddGames(b)
}

// UpdateGames update
func UpdateGames(b *db.Games) (*db.Games, error) {
	return db.UpdateGames(b)
}

// GetGamesByID get by id
func GetGamesByID(id int) (*db.Games, error) {
	return db.GetGamesByID(id)
}

// ListGames  page by condition
func ListGames(b *db.Games) (*db.DataStore, error) {
	list, err := db.ListGames(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountGames(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + b.PageSize - 1) / b.PageSize}, nil
}

// DeleteGames delete
func DeleteGames(id int) error {
	return db.DeleteGames(id)
}
