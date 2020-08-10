package db

import "magic/global"

/*

CREATE TABLE `garden_atlas` (
  `_id` int(11) NOT NULL COMMENT '主键',
  `seed_id` int(11) DEFAULT NULL COMMENT '种子id',
  `flower_cate_name` varchar(255) DEFAULT NULL COMMENT '花种类名称',
  `flower_image` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '图片路径',
  `rarity` int(11) DEFAULT NULL COMMENT '稀有度(0 普通 1独特 2稀有)',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='花之图谱'
*/

// GardenAtlas 花之图谱
type GardenAtlas struct {
	ID             int    `gorm:"column:_id" json:"_id" form:"_id"`
	SeedID         int    `gorm:"column:seed_id" json:"seed_id" form:"seed_id"`
	FlowerCateName string `gorm:"column:flower_cate_name" json:"flower_cate_name" form:"flower_cate_name"`
	FlowerImage    string `gorm:"column:flower_image" json:"flower_image" form:"flower_image"`
	Rarity         int    `gorm:"column:rarity" json:"rarity" form:"rarity"`
}

// TableName 表名
func (o *GardenAtlas) TableName() string {
	return "garden_atlas"
}

// GetGardenAtlasBySeedID 根据种子id查询图谱 1-n
func GetGardenAtlasBySeedID(seedid int) ([]*GardenAtlas, error) {
	db := global.MYSQL
	var results []*GardenAtlas
	err := db.Table("garden_atlas").Where("seed_id = ?", seedid).Scan(&results).Error
	return results, err
}
