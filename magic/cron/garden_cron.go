package cron

import (
	"magic/db"
	"magic/global"
	mylog "magic/utils/log"

	"github.com/robfig/cron"
)

// InitCronTask 定时任务
func InitCronTask() {
	go signin()
	go computerFower()
	go setDisaster()
}

// 设置自然灾害
func setDisaster() {
	mylog.Info("并设置自然灾害定时任务启动")
	c := cron.New()
	c.AddFunc("0 0 */1 * *", func() {
		// c.AddFunc("*/3 * * * *", func() {
		mylog.Info("开始并设置自然灾害")
		length, _ := db.SetGardenFlowerpotDisaster()
		mylog.Success("设置自然灾害结束,累计处理%v条数据", length)
	})
	c.Start()
	select {}
}

// 计算成花时间
func computerFower() {
	mylog.Info("计算成花时间定时任务启动")
	c := cron.New()
	c.AddFunc("0 */1 * * *", func() {
		// c.AddFunc("0 0 0 * *", func() {
		mylog.Info("开始计算成花时间")
		length, _ := db.AllGardenFlowerpot()
		mylog.Success("计算成花时间结束,累计处理%v条数据", length)
	})
	c.Start()
	select {}
}

// 每天零点重置签到 TODO 中间是否有断签到？
func signin() {
	mylog.Info("更新每日签到,清除获得经验限制,定时任务启动")
	c := cron.New()
	// c.AddFunc("*/3 * * * * ?", func() {
	c.AddFunc("0 0 0 * *", func() {
		mylog.Warning("开始更新每日签到")
		tx := global.MYSQL.Begin()
		sql := "update garden set is_signin = 1,g_sow_exp = 0,g_handle_exp = 0,g_current_ex = 0;"
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			mylog.Error("更新每日签到错误::err", err)
			return
		}
		mylog.Success("更新每日签到成功")
	})
	c.Start()
	select {}
}
