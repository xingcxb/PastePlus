// Package bindings 设置部分交互接口
package bindings

import (
	"PastePlus/core/basic/common"
	"encoding/json"
)

type Set struct {
	name string
}

// SetService 设置对外服务
type SetService struct {
}

// LoadingPasteData 加载历史剪贴板数据
func (*SetService) LoadingPasteData() string {
	pastes := make([]common.PasteHistoryGo, 0)
	// todo 从数据库中获取历史剪贴板数据
	// 此处是假数据
	pastes = append(pastes, common.PasteHistoryGo{
		Id:        1,
		FromApp:   "QQ",
		Content:   []byte("这是一条来自QQ的剪贴板数据"),
		Type:      "text",
		CreatedAt: "2021-01-01 00:00:00",
	})
	pastes = append(pastes, common.PasteHistoryGo{
		Id:        2,
		FromApp:   "微信",
		Content:   []byte("这是一条来自微信的剪贴板数据"),
		Type:      "text",
		CreatedAt: "2021-01-01 00:00:00",
	})
	pastesJsonByte, err := json.Marshal(pastes)
	if err != nil {
		return ""
	}
	return string(pastesJsonByte)
}
