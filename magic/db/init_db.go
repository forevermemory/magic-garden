package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	// MYSQL 全局连接对象
	MYSQL *gorm.DB
)

func init() {
	godotenv.Load() // 加载 .env 配置文件
	if MYSQL == nil {
		var err error
		count := 1
		for {
			MYSQL, err = gorm.Open("mysql", os.Getenv("mysql"))
			if err != nil {
				if count == 1 {
					fmt.Println("数据库连接失败,开始连接时间为:", time.Now().Format("2006-01-02 15:04:05"))
				}
				fmt.Println("数据库连接失败,10s后尝试下一次连接,当前连接总的次数为:", count)
				count++
				time.Sleep(time.Second * 10)
			} else {
				fmt.Println("连接数据库成功")
				MYSQL.DB().SetMaxIdleConns(10)
				MYSQL.DB().SetMaxOpenConns(100)
				// MYSQL.LogMode(true) // true 打印sql日志
				break
			}
		}
	}
}
