package db

// DataStore list的结构体
type DataStore struct {
	Total     int64       `json:"total"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}
