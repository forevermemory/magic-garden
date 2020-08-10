package global

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	// MYSQL 全局连接对象
	MYSQL *gorm.DB
	// REDIS 全局redis连接池
	REDIS *redis.Pool
)

// InitGlobalConn 初始化
func InitGlobalConn() {
	godotenv.Load() // 加载 .env 配置文件
	initMysql()
	initRedis()
}
func initMysql() {
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

func initRedis() {
	if REDIS == nil {
		pool := &redis.Pool{
			// Other pool configuration not shown in this example.
			MaxActive: 1024,
			MaxIdle:   16,
			Wait:      true,
			Dial: func() (redis.Conn, error) {
				count := 0
				for {
					count++
					c, err := redis.Dial("tcp", os.Getenv("redis"))
					if err != nil {
						fmt.Println("连接redis错误,10s后尝试下一次连接,当前连接总的次数为:", count)
						time.Sleep(time.Second * 5)
					} else {
						// 登陆
						// if _, err := c.Do("AUTH", "password"); err != nil {
						// 	c.Close()
						// 	return nil, err
						// }
						// 使用1号存储
						if _, err := c.Do("SELECT", 1); err != nil {
							c.Close()
							fmt.Println("选择一号存储错误,正在重新连接,当前连接总的次数为:", count)
							continue
						}
						fmt.Println("redis 连接成功")
						return c, nil
					}
				}

			},
		}
		REDIS = pool
		pool.Get().Close()
	}
}
