//go:build darwin

package hook

import (
	"PastePlus/core/plugin/db"
	"context"
	"fmt"
	"github.com/xingcxb/goKit/core/dateKit"
	"golang.design/x/clipboard"
)

// RegexListen 注册监听
func RegexListen() {
	hookClipboard()
}

func hookClipboard() {
	// 初始化剪切板包
	err := clipboard.Init()
	if err != nil {
		fmt.Println("剪切板初始化失败", err)
		return
	}
	// 读取剪切板数据
	//clipboardData := clipboard.Read(clipboard.FmtText)
	//fmt.Println("剪切板数据：", string(clipboardData))
	// 监听剪切板的变化
	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	for dataByte := range ch {
		// 获取该内容是否存在于数据库中
		paste := db.FindDataByContent(dataByte)
		if paste.Id == 0 {
			// 如果不存在那么就需要重新构建数据
			paste.FromApp = ""
			paste.Content = dataByte
			paste.Type = ""
			paste.CreatedAt = dateKit.Now()
		}
		db.SaveOrUpdatePaste(paste)
		fmt.Println("Clipboard changed:", string(dataByte))
	}
}
