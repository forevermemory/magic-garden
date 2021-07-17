package cron

import (
	"magic/db"
	"magic/global"
	mylog "magic/utils/log"
	"time"

	"github.com/robfig/cron"
)

// */5 * * * * *    -- 每隔5s
// 1,11,21,31,41,51 * * * * *    -- 1,11,21,31,41,51 s执行
// 1-5,10-20,30-50 * * * * *    -- 1-5,10-20,30-50 该区间段执行执行

// * 1 * * * *   -- 每小时中1分钟每秒钟执行一次
// * 17-19 * * * *   -- 17-19min 每秒钟执行一次
// * */28 * * * *  -- 每小时分28分中每一秒执行一次
// 00 17-19 * * * *   -- 17-19min 每分钟执行一次
// 00 10,30,40 * * * *   --  每小时中10 30 40 分执行一次
// 00 */1 * * * *   --  每分钟
// 00 */10 * * * *   --  每十分钟
// 00 */30 * * * *   --  每三十分钟

// 同理类推
// * * 20 * * *  -- 每天20点中每一秒执行一次
// 00 * 20 * * *  -- 每天20点中每一分执行一次
// 00 00 20 * * *  -- 每天20点执行一次

// 同理类推...

// TestCronisaster TestCronisaster
func TestCronisaster() {
	mylog.Info("并设置自然灾害定时任务启动")
	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() { // 1h执行一次
		// c.AddFunc("*/3 * * * *", func() {
		mylog.Success("设置自然灾害")
	})
	c.Start()
	select {}
}

// 0 */1 * * * 1min
//  0 0 * * * * 这是 1h
// Second: second,
// Minute: minute,
// Hour:   hour,
// Dom:    dayofmonth,
// Month:  month,
// Dow:    dayofweek,

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
	c.AddFunc("0 0 * * *", func() { // 1h执行一次
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
	c.AddFunc("0 0 * * *", func() { //  这是 1min 处理一次
		// c.AddFunc("0 0 0 * *", func() {
		mylog.Info("开始计算成花时间")
		length, _ := db.AllGardenFlowerpot()
		mylog.Success("计算成花时间结束,累计处理%v条数据", length)
	})
	c.Start()
	select {}
}

// 每天零点重置签到,重置每日获得的经验值
func signin() {
	mylog.Info("更新每日签到,清除获得经验限制,定时任务启动")
	c := cron.New()
	// c.AddFunc("*/3 * * * * ?", func() {
	c.AddFunc("0 0 0 * *", func() { // 0 0 * * * 这是 1h
		mylog.Warning("开始更新每日签到")
		tx := global.MYSQL.Begin()
		sql := "update garden set is_signin = 1,g_sow_exp = 0,g_handle_exp = 0,g_current_ex = 0"
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			mylog.Error("更新每日签到错误::err", err)
			return
		}
		//  中间是否有断签到
		gardens, _ := db.ListGarden()
		for _, garden := range gardens {
			timeL, _ := time.ParseInLocation("2006-01-02 15:04:05", garden.GSigninTime, time.Local)
			if time.Now().Day()-timeL.Day() > 1 {
				// 中间有一天没有签到呢
				garden.SignDays = 1
				_, err := db.UpdateGarden(garden, tx)
				if err != nil {
					tx.Rollback()
					mylog.Error("UpdateGarden::err", err)
					return
				}
			}

		}

		// ok
		tx.Commit()
		mylog.Success("更新每日签到成功")
	})
	c.Start()
	select {}
}
