// Package cron 定时任务，用于清理数据
package cron

import (
	"PastePlus/core/kit"
	"PastePlus/core/plugin/db"
	"github.com/robfig/cron/v3"
)

// CreateCron 创建定时任务
func CreateCron() {
	c := cron.New()
	// 每隔半个小时清理一次过期的历史记录
	_, err := c.AddFunc("@every 30m", func() {
		ClearPasteHistory()
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
	// 获取配置信息
	config, err := db.FindConfigByKey("historyCapacity")
	if err != nil {
		return
	}
	startDate := kit.GetDateByTimeUnit(config.Value)
	cleanData := db.FindPasteListByLTDate(startDate)
	for _, v := range cleanData {
		db.DeletePasteById(v.Id)
	}
}
