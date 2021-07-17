package service

/*
date:2021-07-17 14:20:14
*/

import "magic/db"

// AddUserLevel add
func AddUserLevel(req *db.UserLevel) (*db.UserLevel, error) {
	return db.AddUserLevel(req)
}

// UpdateUserLevel update
func UpdateUserLevel(req *db.UserLevel) (*db.UserLevel, error) {
	return db.UpdateUserLevel(req)
}

// GetUserLevelByID get by id
func GetUserLevelByID(id int) (*db.UserLevel, error) {
	return db.GetUserLevelByID(id)
}

// ListUserLevel  page by condition
func ListUserLevel(req *db.UserLevel) (*db.DataStore, error) {
	list, err := db.ListUserLevel(req)
	if err != nil {
		return nil, err
	}
	total, err := db.CountUserLevel(req)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + req.PageSize - 1) / req.PageSize}, nil
}

// DeleteUserLevel delete
func DeleteUserLevel(id int) error {
	return db.DeleteUserLevel(id)
}
