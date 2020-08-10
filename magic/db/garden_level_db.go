package db

import "magic/global"

/*
CREATE TABLE `garden_flower_level` (
  `level_num` int(11) NOT NULL AUTO_INCREMENT COMMENT '级别',
  `level_name` varchar(255) DEFAULT NULL COMMENT '等级名称',
  `up_level` varchar(255) DEFAULT NULL COMMENT '升级所需要经验',
  PRIMARY KEY (`level_num`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COMMENT='等级表'
*/

// GardenFlowerLevel 等级
type GardenFlowerLevel struct {
	LevelNum  int    `gorm:"column:level_num" json:"level_num" form:"level_num"`
	LevelName string `gorm:"column:level_name" json:"level_name" form:"level_name"`
	UpLevel   string `gorm:"column:up_level" json:"up_level" form:"up_level"`
}

// TableName 表名
func (o *GardenFlowerLevel) TableName() string {
	return "garden_flower_level"
}

// GetGardenFlowerLevelByLevelNum 根据等级查询
func GetGardenFlowerLevelByLevelNum(num int) (*GardenFlowerLevel, error) {
	var res GardenFlowerLevel
	err := global.MYSQL.Table("garden_flower_level").Where("level_num = ?", num).First(&res).Error
	return &res, err
}
