package utils

import "time"

// GetNowTimeString 获取当前时间
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
