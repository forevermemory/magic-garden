package db

import (
	"magic/global"
	"magic/utils"

	"github.com/jinzhu/gorm"
)

/*
CREATE TABLE `garden_flower_house` (
  `_id` int(11) DEFAULT NULL COMMENT 'pk',
  `garden_id` int(11) DEFAULT NULL COMMENT '花园id',
  `atlas_id` int(11) DEFAULT NULL COMMENT '图谱id',
  `num` varchar(255) DEFAULT '0' COMMENT '数量',
  `cate` int(11) DEFAULT 1 COMMENT '分类 1花篮 2花瓶'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='花房'
*/

// GardenFlowerHouse 花房
type GardenFlowerHouse struct {
	ID             int    `gorm:"column:_id" json:"_id" form:"_id"`
	GardenID       string `gorm:"column:garden_id" json:"garden_id" form:"garden_id"`
	AtlasID        int    `gorm:"column:atlas_id" json:"atlas_id" form:"atlas_id"`
	Num            int    `gorm:"column:num" json:"num" form:"num"`
	Cate           int    `gorm:"column:cate" json:"cate" form:"cate"`
	SeedID         int    `gorm:"-" json:"seed_id" form:"seed_id"`
	FlowerCateName string `gorm:"-" json:"flower_cate_name" form:"flower_cate_name"`
	FlowerImage    string `gorm:"-" json:"flower_image" form:"flower_image"`
	Rarity         int    `gorm:"-" json:"rarity" form:"rarity"`
}

// TableName 表名
func (o *GardenFlowerHouse) TableName() string {
	return "garden_flower_house"
}

// IsLightupAtlas 是否点亮图谱
func IsLightupAtlas(result int, gardenid string) (*GardenFlowerHouse, error) {
	var o GardenFlowerHouse
	err := global.MYSQL.Table("garden_flower_house").Where("cate = 1").Where("atlas_id = ? and garden_id = ?", result, gardenid).First(&o).Error
	return &o, err
}

// AddGardenFlowerHouse 新增
func AddGardenFlowerHouse(o *GardenFlowerHouse, tx ...*gorm.DB) (*GardenFlowerHouse, error) {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	err := db.Create(o).Error
	return o, err
}

// UpdateGardenFlowerHouse 修改会涉及到修改数量
func UpdateGardenFlowerHouse(o *GardenFlowerHouse, tx ...*gorm.DB) (*GardenFlowerHouse, error) {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	err := db.Table("garden_flower_house").Where("_id=?", o.ID).Update(o).First(o).Error
	return o, err
}

// UpdateGardenFlowerHouseNumber 修改数量
func UpdateGardenFlowerHouseNumber(num int, gardenID string, atlasID int, tx ...*gorm.DB) error {
	db := global.MYSQL
	if tx != nil {
		db = tx[0]
	}
	sql := "update garden_flower_house set num = ? where garden_id = ? and atlas_id = ?"
	return db.Exec(sql, num, gardenID, atlasID).Error
}

// ListGardenFlowerHouse 查询 根据花园id查询出所有的 花 需要改进 连表之类的 TODO
func ListGardenFlowerHouse(gardenid string, cate int, page int) ([]*GardenFlowerHouse, error) {
	res := make([]*GardenFlowerHouse, 0)
	db := global.MYSQL
	sql := "select aa.*,bb.* from garden_flower_house aa left join garden_atlas bb on aa.atlas_id = bb._id  where aa.garden_id = ? and aa.cate = ? and aa.num >0 order by atlas_id  limit ?,?"
	err := db.Raw(sql, gardenid, cate, global.PageSize*(page-1), global.PageSize).Scan(&res).Error
	return res, err
}

// CountGardenFlowerHouse count
func CountGardenFlowerHouse(gardenid string, cate int) (int64, error) {
	db := global.MYSQL
	var count int64
	err := db.Table("garden_flower_house").Where("cate = ? and garden_id = ? and num >0", cate, gardenid).Count(&count).Error
	return count, err
}

// GardenHouseStatistics 统计
func GardenHouseStatistics(gardenid string, cate int) (interface{}, error) {
	sql := `SELECT count(1) total_zhong,sum(num) total_duo FROM garden_flower_house WHERE garden_id = ?  and cate = ?
	GROUP BY cate`
	rows, err := global.MYSQL.Raw(sql, gardenid, cate).Rows()
	if err != nil {
		return nil, err
	}
	res, err := utils.SQLMap(rows)
	if err != nil {
		return nil, err
	}
	return res, nil
}
