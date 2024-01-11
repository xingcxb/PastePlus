// Package common 剪贴板数据库结构体
package common

const (
	ConfigKeyBootUp          = "bootUp"          // ConfigKeyBootUp 开机启动配置键
	ConfigKeySound           = "sound"           // ConfigKeySound 声音配置键
	ConfigKeyMenuIcon        = "menuIcon"        // ConfigKeyMenuIcon 菜单栏图标配置键
	ConfigKeyHistoryCapacity = "historyCapacity" // ConfigKeyHistoryCapacity 历史记录容量配置键
)

// PasteHistoryGo 剪贴板历史记录结构体Go使用
type PasteHistoryGo struct {
	Id        int64  `json:"id"`         // ID 历史记录ID
	FromApp   string `json:"from_app"`   // FromApp 历史记录来源应用
	Content   []byte `json:"content"`    // Content 历史记录内容
	Type      string `json:"type"`       // Type 历史记录类型
	CreatedAt string `json:"created_at"` // CreatedAt 历史记录创建时间
}

// PasteHistoryVue 剪贴板历史记录结构体提供给vue使用
type PasteHistoryVue struct {
	Id          int64  `json:"id"`           // ID 历史记录ID
	FromApp     string `json:"from_app"`     // FromApp 历史记录来源应用
	Content     string `json:"content"`      // Content 历史记录内容
	Type        string `json:"type"`         // Type 历史记录类型
	SpacingTime string `json:"spacing_time"` // SpacingTime 历史记录和当前间隔时间
	CreatedAt   string `json:"created_at"`   // CreatedAt 历史记录创建时间
}

// PasteConfigVue 基础配置信息结构体提供给vue使用
type PasteConfigVue struct {
	Key   string `json:"key"`   // Key 配置键
	Value string `json:"value"` // Value 配置值
}

// PasteConfig 基础配置信息
type PasteConfig struct {
	Id    int64  `json:"id"`    // ID 配置ID
	Key   string `json:"key"`   // 键
	Value string `json:"value"` // 值
}
