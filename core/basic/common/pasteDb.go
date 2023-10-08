// Package common 剪贴板数据库结构体
package common

// PasteHistory 剪贴板历史记录结构体
type PasteHistory struct {
	Id        int64  `json:"id"`         // ID 历史记录ID
	FromApp   string `json:"from_app"`   // FromApp 历史记录来源应用
	Content   string `json:"content"`    // Content 历史记录内容
	Type      string `json:"type"`       // Type 历史记录类型
	CreatedAt string `json:"created_at"` // CreatedAt 历史记录创建时间
}

// PasteConfig 基础配置信息
type PasteConfig struct {
	Id    int64 `json:"id"`    // ID 配置ID
	Key   bool  `json:"key"`   // 键
	Value bool  `json:"value"` // 值
}
