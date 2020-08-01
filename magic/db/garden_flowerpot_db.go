package db

import "magic/global"

/*
CREATE TABLE `garden_flowerpot` (
CREATE TABLE `garden_flowerpot` (
  `_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `garden_id` int(11) DEFAULT NULL COMMENT '花园id和用户id相同',
  `number` int(11) DEFAULT NULL COMMENT '花盆编号 1-10',
  `is_lock` int(11) DEFAULT '0' COMMENT '是否解锁该花盆 1 未解锁 2 已经解锁',
  `is_sow` int(11) DEFAULT NULL COMMENT '空还是播种 1 空 2 播种',
  `seed_id` int(11) DEFAULT NULL COMMENT '种的种子',
  `status` int(11) DEFAULT '0' COMMENT '1 正常 2 干旱 3 有虫 4 有草',
  `seed_result` int(11) DEFAULT '0' COMMENT '种子开花结果 开出的花是图谱id',
  `flower_num` int(11) DEFAULT '0' COMMENT '最后成花数量 珍惜的只会一朵 其它的0种子 不处理开原数量 浇水或者除虫、除草会翻倍产量',
  `stage` int(11) DEFAULT NULL COMMENT '花的成长阶段 1 花苗 2 花蕾 3 开花',
  `stage_str` varchar(255) DEFAULT NULL COMMENT '成长剩余时间',
  `sow_time` varchar(255) DEFAULT NULL COMMENT '播种时间',
  `is_change_color` int(11) DEFAULT '0' COMMENT '是否使用染色剂 1 未使用 2使用了',
  `change_result` varchar(255) DEFAULT NULL COMMENT '染色结果string',
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='花盆'
*/

// GardenFlowerpot 花盆
type GardenFlowerpot struct {
	ID            int    `gorm:"column:_id" json:"_id" form:"_id"`
	UserID        int    `gorm:"column:user_id" json:"user_id" form:"user_id"`
	GardenID      int    `gorm:"column:garden_id" json:"garden_id" form:"garden_id"`
	Number        int    `gorm:"column:number" json:"number" form:"number"`
	IsLock        int    `gorm:"column:is_lock" json:"is_lock" form:"is_lock"`
	IsSow         int    `gorm:"column:is_sow" json:"is_sow" form:"is_sow"`
	SeedID        int    `gorm:"column:seed_id" json:"seed_id" form:"seed_id"`
	Status        int    `gorm:"column:status" json:"status" form:"status"`
	SeedResult    int    `gorm:"column:seed_result" json:"seed_result" form:"seed_result"`
	FlowerNum     int    `gorm:"column:flower_num" json:"flower_num" form:"flower_num"`
	Stage         int    `gorm:"column:stage" json:"stage" form:"stage"`
	StageStr      string `gorm:"column:stage_str" json:"stage_str" form:"stage_str"`
	SowTime       string `gorm:"column:sow_time" json:"sow_time" form:"sow_time"`
	IsChangeColor int    `gorm:"column:is_change_color" json:"is_change_color" form:"is_change_color"`
	ChangeResult  string `gorm:"column:change_result" json:"change_result" form:"change_result"`
}

// TableName 表名
func (o *GardenFlowerpot) TableName() string {
	return "garden_flowerpot"
}

// AddGardenFlowerpot 新增
func AddGardenFlowerpot(o *GardenFlowerpot) error {
	db := global.MYSQL
	return db.Create(o).Error
}

// UpdateGardenFlowerpot 修改
func UpdateGardenFlowerpot(o *GardenFlowerpot) error {
	db := global.MYSQL
	return db.Table("garden_flowerpot").Where("_id=?", o.ID).Update(o).Error
}

// ListGardenFlowerpot 查询某个花园的花盆列表
func ListGardenFlowerpot(gardenid int) ([]*GardenFlowerpot, error) {
	res := make([]*GardenFlowerpot, 0)
	db := global.MYSQL
	err := db.Table("garden_flowerpot").Where("garden_id=?", gardenid).Scan(&res).Error
	return res, err
}
