package db

import (
	"errors"
	"magic/global"

	"github.com/jinzhu/gorm"
)

/*
CREATE TABLE `garden_flower_knapsack` (
  `_id` int(11) DEFAULT NULL COMMENT 'pk',
  `garden_id` int(11) DEFAULT NULL COMMENT '花园id',
  `seed_id` int(11) DEFAULT NULL COMMENT '种子id',
  `seed_num` varchar(255) DEFAULT '0' COMMENT '种子数量',
  `cate` int(11) DEFAULT NULL COMMENT '分类 1种子 2道具',
  `prop_id` int(11) DEFAULT NULL COMMENT '道具id',
  `prop_num` varchar(255) DEFAULT NULL COMMENT '道具数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='背包'
*/

// GardenFlowerKnapsack  背包
type GardenFlowerKnapsack struct {
	ID       int    `gorm:"column:_id" json:"_id" form:"_id"`
	GardenID int    `gorm:"column:garden_id" json:"garden_id" form:"garden_id"`
	SeedID   int    `gorm:"column:seed_id" json:"seed_id" form:"seed_id"`
	SeedNum  int    `gorm:"column:seed_num" json:"seed_num" form:"seed_num"`
	Cate     int    `gorm:"column:cate" json:"cate" form:"cate"`
	PropID   int    `gorm:"column:prop_id" json:"prop_id" form:"prop_id"`
	PropNum  int    `gorm:"column:prop_num" json:"prop_num" form:"prop_num"`
	Page     int    `json:"page" form:"page" gorm:"-" `
	SeedName string `json:"seed_name" form:"seed_name" `
	PropName string `json:"prop_name" form:"prop_name" `
}

// TableName 表名
func (o *GardenFlowerKnapsack) TableName() string {
	return "garden_flower_knapsack"
}

// AddGardenFlowerKnapsack 新增
func AddGardenFlowerKnapsack(o *GardenFlowerKnapsack, tx ...*gorm.DB) error {
	db := global.MYSQL
	if len(tx) != 0 {
		db = tx[0]
	}
	return db.Create(o).Error
}

// UpdateGardenFlowerKnapsack 修改会涉及到修改数量
func UpdateGardenFlowerKnapsack(o *GardenFlowerKnapsack, tx ...*gorm.DB) error {
	db := global.MYSQL
	if len(tx) != 0 {
		db = tx[0]
	}
	return db.Table("garden_flower_knapsack").Where("_id=?", o.ID).Update(o).Error
}

// UpdateGardenFlowerKnapsackHandleSeedNumZelo 修改种子数量
func UpdateGardenFlowerKnapsackHandleSeedNumZelo(o *GardenFlowerKnapsack, tx ...*gorm.DB) (*GardenFlowerKnapsack, error) {
	db := global.MYSQL
	if len(tx) != 0 {
		db = tx[0]
	}
	sql := " UPDATE `garden_flower_knapsack` SET seed_num = ?  WHERE (_id=?) "
	// return db.Debug().Table("garden_flower_knapsack").Where("_id=?", o.ID).Update(o).Error

	err := db.Debug().Exec(sql, o.SeedNum, o.ID).Error
	// fmt.Println("db,,,", o.SeedNum)
	return o, err
}

// ListGardenFlowerKnapsack2 查询 根据花园id查询出所有的 花 需要改进 连表之类的 TODO
func ListGardenFlowerKnapsack2(gardenid int, page int) ([]*GardenFlowerKnapsack, error) {
	res := make([]*GardenFlowerKnapsack, 0)
	db := global.MYSQL
	sql := "select * from garden_flower_knapsack where garden_id = ? order by atlas_id  limit ?,?"
	err := db.Raw(sql, gardenid, global.PageSize*(page-1), global.PageSize).Scan(&res).Error
	return res, err
}

// ListGardenFlowerKnapsack 分页条件查询
func ListGardenFlowerKnapsack(o *GardenFlowerKnapsack) ([]*GardenFlowerKnapsack, error) {
	res := make([]*GardenFlowerKnapsack, 0)
	sql := ""
	if o.Cate == 1 {
		sql = "select aa.*,bb.seed_name from garden_flower_knapsack aa  inner join garden_seeds bb on aa.seed_id = bb._id where aa.cate = 1 and garden_id = ? order by aa.seed_id limit ?,?"
	} else if o.Cate == 2 {
		sql = "select aa.*,bb.p_name prop_name from garden_flower_knapsack aa  inner join garden_props bb on aa.prop_id = bb._id where aa.cate = 2 and garden_id = ? ORDER BY aa.prop_id limit ?,? "

	} else {
		return res, errors.New("请正确传入cate参数")
	}
	// err := global.MYSQL.Table("garden_flower_knapsack").Where(o).Offset((o.Page - 1) * global.PageSize).Limit(global.PageSize).Find(&res).Error
	err := global.MYSQL.Raw(sql, o.GardenID, (o.Page-1)*global.PageSize, global.PageSize).Scan(&res).Error
	return res, err
}

// CountGardenFlowerKnapsack 条件数量
func CountGardenFlowerKnapsack(o *GardenFlowerKnapsack) (int64, error) {
	var count int64
	err := global.MYSQL.Table("garden_flower_knapsack").Where(o).Count(&count).Error
	return count, err
}

// CountGardenFlowerKnapsackV2 查询出某个花园下的背包下某个种子的总的数量
func CountGardenFlowerKnapsackV2(gardenID int, seedID int) (*GardenFlowerKnapsack, error) {
	var tmp GardenFlowerKnapsack
	err := global.MYSQL.Table("garden_flower_knapsack").Where("garden_id = ? and seed_id = ?", gardenID, seedID).First(&tmp).Error
	return &tmp, err
}

// IsExistGardenFlowerKnapsackSeed 背包是否存在种子
func IsExistGardenFlowerKnapsackSeed(gardenID int, seedID int) (*GardenFlowerKnapsack, error) {
	var res GardenFlowerKnapsack
	err := global.MYSQL.Table("garden_flower_knapsack").Where("garden_id = ? and seed_id = ?", gardenID, seedID).First(&res).Error
	return &res, err
}

// IsExistGardenFlowerKnapsackProp 背包是否存在道具
func IsExistGardenFlowerKnapsackProp(gardenID int, propID int) (*GardenFlowerKnapsack, error) {
	var res GardenFlowerKnapsack
	err := global.MYSQL.Table("garden_flower_knapsack").Where("garden_id = ? and prop_id = ?", gardenID, propID).First(&res).Error
	return &res, err
}
