package logic

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/plugin/db"
	"go.uber.org/zap"
)

// HandConfigValue 操作配置信息
/*
 * @param key 配置信息的key
 * @param value 要更新的value
 */
func HandConfigValue(key, value string) bool {
	c, err := db.FindConfigByKey(key)
	if err != nil {
		common.Logger.Error("查询配置信息失败", zap.Error(err))
		return false
	}
	if c.Id == 0 {
		common.Logger.Error("未查询到配置信息", zap.Error(err))
		return false
	}
	c.Value = value
	return db.UpConfig(c)
}
