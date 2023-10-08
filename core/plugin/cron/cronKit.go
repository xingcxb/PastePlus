// Package cron 定时任务，用于清理数据
package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

// CreateCron 创建定时任务
func CreateCron() {
	c := cron.New()
	_, err := c.AddFunc("@every 1m", func() {
		fmt.Println("tick every 1 second")
	})
	if err != nil {
		return
	}
	c.Start()
}

// StopCron 停止定时任务
func StopCron() {
	c := cron.New()
	c.Stop()
}

// ClearPasteHistory 清理过期的历史记录
func ClearPasteHistory() {

}
