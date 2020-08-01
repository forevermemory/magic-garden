package db

import "magic/global"

/*
date:2020-08-01 13:07:54
*/

// GardenHelp GardenHelp
type GardenHelp struct {
	ID       int    `json:"_id" form:"_id" gorm:"column:_id;primary_key;auto_increment;comment:'主键'"`
	HTitle   string `json:"h_title" form:"h_title"`
	HContent string `json:"h_content" form:"h_content"`
}

// TableName 表名
func (o *GardenHelp) TableName() string {
	return "garden_help"
}

// GetGardenHelpByID 根据id查询一个
func GetGardenHelpByID(id int) (*GardenHelp, error) {
	db := global.MYSQL
	o := &GardenHelp{}
	err := db.Table("garden_help").Where("_id = ?", id).First(o).Error
	return o, err
}

// AddGardenHelp 新增
func AddGardenHelp(o *GardenHelp) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// UpdateGardenHelp 修改
func UpdateGardenHelp(o *GardenHelp) error {
	db := global.MYSQL
	return db.Table("garden_help").Where("_id=?", o.ID).Update(o).Error
}

// AllGardenHelp 查询所有
func AllGardenHelp() ([]*GardenHelp, error) {
	db := global.MYSQL
	res := make([]*GardenHelp, 0)
	err := db.Table("garden_help").Find(&res).Error
	return res, err
}

// GetGardenHelpTitles 查询所有标题
func GetGardenHelpTitles() ([]*GardenHelp, error) {
	db := global.MYSQL
	res := make([]*GardenHelp, 0)
	sql := "select _id,h_title from garden_help"
	err := db.Raw(sql).Find(&res).Error
	return res, err
}

// CountGardenHelp 条件数量
func CountGardenHelp(o *GardenHelp) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("garden_help").Count(&count).Error
	return count, err
}
