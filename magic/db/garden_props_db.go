package db

import "magic/global"

/*

CREATE TABLE `garden_props` (
  `_id` int(11) NOT NULL AUTO_INCREMENT,
  `p_name` varchar(255) DEFAULT NULL COMMENT '道具名称',
  `p_price` varchar(255) DEFAULT NULL COMMENT '价格(GB或者元宝)',
  `p_time` varchar(255) DEFAULT NULL COMMENT '减少的时间',
  `p_desc` varchar(255) DEFAULT NULL COMMENT '备注',
  `is_buy` int(11) DEFAULT NULL COMMENT '是否可以用gb购买 1 不可买 2可购买',
  PRIMARY KEY (`_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='道具'
*/

// GardenProps 道具
type GardenProps struct {
	ID     int    `gorm:"column:_id;primary_key" json:"_id" form:"_id"`
	PName  string `gorm:"column:p_name" json:"p_name" form:"p_name"`
	PPrice int    `gorm:"column:p_price" json:"p_price" form:"p_price"`
	PTime  string `gorm:"column:p_time" json:"p_time" form:"p_time"`
	PDesc  string `gorm:"column:p_desc" json:"p_desc" form:"p_desc"`
	ISbuy  int    `gorm:"column:is_buy" json:"is_buy" form:"is_buy"`
}

// TableName 表名
func (o *GardenProps) TableName() string {
	return "garden_props"
}

// GetGardenPropsByID 根据id查道具
func GetGardenPropsByID(propid int) (*GardenProps, error) {
	db := global.MYSQL
	var res GardenProps
	err := db.Table("garden_props").Where("_id = ? ", propid).First(&res).Error
	return &res, err
}
