package db

import "magic/global"

/*
CREATE TABLE `garden_flower_house` (
  `_id` int(11) DEFAULT NULL COMMENT 'pk',
  `garden_id` int(11) DEFAULT NULL COMMENT '花园id',
  `atlas_id` int(11) DEFAULT NULL COMMENT '图谱id',
  `num` varchar(255) DEFAULT '0' COMMENT '数量',
  `cate` int(11) DEFAULT '1' COMMENT '分类 1花篮 2花瓶'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='花房'
*/

// GardenFlowerHouse 花房
type GardenFlowerHouse struct {
	ID       int    `gorm:"column:_id" json:"_id" form:"_id"`
	GardenID int    `gorm:"column:garden_id" json:"garden_id" form:"garden_id"`
	AtlasID  int    `gorm:"column:atlas_id" json:"atlas_id" form:"atlas_id"`
	Num      string `gorm:"column:num" json:"num" form:"num"`
	Cate     int    `gorm:"column:cate" json:"cate" form:"cate"`
}

// TableName 表名
func (o *GardenFlowerHouse) TableName() string {
	return "garden_flower_house"
}

// AddGardenFlowerHouse 新增
func AddGardenFlowerHouse(o *GardenFlowerHouse) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// UpdateGardenFlowerHouse 修改会涉及到修改数量
func UpdateGardenFlowerHouse(o *GardenFlowerHouse) error {
	db := global.MYSQL
	return db.Table("garden_flower_house").Where("_id=?", o.ID).Update(o).Error
}

// ListGardenFlowerHouse 查询 根据花园id查询出所有的 花 需要改进 连表之类的 TODO
func ListGardenFlowerHouse(gardenid int, page int) ([]*GardenFlowerHouse, error) {
	res := make([]*GardenFlowerHouse, 0)
	db := global.MYSQL
	sql := "select * from garden_flower_house where garden_id = ? order by atlas_id  limit ?,?"
	err := db.Raw(sql, gardenid, global.PageSize*(page-1), global.PageSize).Scan(&res).Error
	return res, err
}
