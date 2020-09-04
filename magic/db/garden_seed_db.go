package db

import "magic/global"

/*
CREATE TABLE `garden_seeds` (
  `_id` int(11) NOT NULL COMMENT 'pk',
  `raw_id` int(11) DEFAULT '0' COMMENT '合成的种子对应的原种子id',
  `img_url` varchar(255) DEFAULT NULL,
  `level_num` int(11) DEFAULT NULL COMMENT '等级',
  `raw_price` int(11) DEFAULT NULL COMMENT '种子价格',
  `vip_price` int(11) DEFAULT NULL COMMENT 'vip价格',
  `level_str` varchar(255) DEFAULT NULL COMMENT '等级str',
  `seed_name` varchar(255) DEFAULT NULL COMMENT '种子名称',
  `rarity` int(11) DEFAULT NULL COMMENT '稀有度',
  `is_change_color` int(11) DEFAULT '0' COMMENT '是否支持使用染色剂',
  `forecast_num` int(11) DEFAULT NULL COMMENT '预计成花(不去浇水除虫会减少产量)',
  `forecast_time` int(11) DEFAULT NULL COMMENT '预计时间(小时)(有肥料可以缩短时间)',
  `meaning` varchar(255) DEFAULT NULL COMMENT '花语',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='种子'
*/

// GardenSeeds 种子
type GardenSeeds struct {
	ID            int    `gorm:"column:_id" json:"_id" form:"_id"`
	RawID         int    `gorm:"column:raw_id" json:"raw_id" form:"raw_id"`
	ImgURL        string `gorm:"column:img_url" json:"img_url" form:"img_url"`
	LevelNum      int    `gorm:"column:level_num" json:"level_num" form:"level_num"`
	RawPrice      int    `gorm:"column:raw_price" json:"raw_price" form:"raw_price"`
	VipPrice      int    `gorm:"column:vip_price" json:"vip_price" form:"vip_price"`
	LevelStr      string `gorm:"column:level_str" json:"level_str" form:"level_str"`
	SeedName      string `gorm:"column:seed_name" json:"seed_name" form:"seed_name"`
	Rarity        int    `gorm:"column:rarity" json:"rarity" form:"rarity"`
	IsChangeColor int    `gorm:"column:is_change_color" json:"is_change_color" form:"is_change_color"`
	ForecastNum   int    `gorm:"column:forecast_num" json:"forecast_num" form:"forecast_num"`
	ForecastTime  int    `gorm:"column:forecast_time" json:"forecast_time" form:"forecast_time"`
	Meaning       string `gorm:"column:meaning" json:"meaning" form:"meaning"`

	PageNo   int `json:"page" form:"page" gorm:"-"`
	PageSize int `json:"page_size" form:"page_size" gorm:"-"`
}

// TableName 表名
func (o *GardenSeeds) TableName() string {
	return "garden_seeds"
}

// GetGardenSeedsByID get one
func GetGardenSeedsByID(seedid int) (*GardenSeeds, error) {
	db := global.MYSQL
	var res GardenSeeds
	err := db.Table("garden_seeds").Where("_id = ? ", seedid).First(&res).Error
	return &res, err
}

// ListGardenMagician 魔法屋可以合成的种子列表
func ListGardenMagician(o *GardenSeeds) ([]GardenSeeds, error) {
	db := global.MYSQL
	var res []GardenSeeds
	sql := "select * from garden_seeds where rarity = 2 limit ?,?"
	err := db.Raw(sql, o.PageNo*global.PageSize, global.PageSize).Scan(&res).Error
	return res, err
}

// CountGardenMagician 条件数量
func CountGardenMagician(o *GardenSeeds) (int64, error) {
	var count int64
	err := global.MYSQL.Table("garden_seeds").Where("rarity = 2").Count(&count).Error
	return count, err
}

// MagicianSeedSynthesisMethods 合成方式
type MagicianSeedSynthesisMethods struct {
	SeedID         int    `json:"seed_id"`
	AtlasID        int    `json:"atlas_id"`
	Num            int    `json:"num"`
	FlowerCateName string `json:"flower_cate_name"`
	FlowerImage    string `json:"flower_image"`
	Rarity         int    `json:"rarity"`
	SeedName       string `json:"seed_name"`
	GardenID       string `json:"garden_id"`
	TotalNum       int    `json:"total_num"`
}

// GardenMagicianDetail  魔法屋查询一个种子合成所需材料
func GardenMagicianDetail(seedID int, gardenID string) ([]MagicianSeedSynthesisMethods, error) {
	var res []MagicianSeedSynthesisMethods
	sql := `
		SELECT aa.*,bb.flower_cate_name,bb.flower_image,bb.rarity,cc.seed_name,dd.garden_id,IFNULL(dd.num,0) total_num 
		FROM garden_synthesis_methods aa 
		left join garden_atlas bb on aa.atlas_id = bb._id 
		left join garden_seeds cc on aa.seed_id = cc._id 
		left join  garden_flower_house dd on dd.atlas_id = aa.atlas_id
		WHERE aa.seed_id = ?  and ( dd.garden_id = ? or dd.garden_id is null)
		ORDER BY atlas_id
	`
	err := global.MYSQL.Raw(sql, seedID, gardenID).Scan(&res).Error
	return res, err
}
