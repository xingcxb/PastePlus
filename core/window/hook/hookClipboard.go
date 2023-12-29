//go:build darwin

package hook

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/plugin/db"
	"context"
	"github.com/xingcxb/goKit/core/dateKit"
	"go.uber.org/zap"
	"golang.design/x/clipboard"
)

// RegexListen 注册监听
func RegexListen() {
	hookClipboard(context.Background())
}

// 钩子处理粘贴板
func hookClipboard(ctx context.Context) {
	// 初始化剪切板包
	err := clipboard.Init()
	if err != nil {
		common.Logger.Error("剪切板初始化失败", zap.Error(err))
		return
	}
	// 读取剪切板数据
	//clipboardData := clipboard.Read(clipboard.FmtText)
	//fmt.Println("剪切板数据：", string(clipboardData))
	// 监听图片格式的数据
	chImage := clipboard.Watch(ctx, clipboard.FmtImage)
	// 监听文本剪切板的数据
	chText := clipboard.Watch(ctx, clipboard.FmtText)
	// 初始化粘贴数据
	var pasteData common.PasteHistoryGo
	for {
		select {
		case <-ctx.Done():
			return
		case dataImage, ok := <-chImage:
			if !ok {
				common.Logger.Info("剪切板图片监听失败")
				return
			}
			// 目前这里存在一个bug，如果使用复制快捷键的话，并不认可
			// win：Shift+Win+s
			// mac：Ctrl+Shift+Cmd+4
			// 剥离数据
			pasteData.Type = "image"
			pasteData.Content = dataImage
		case dataText, ok := <-chText:
			if !ok {
				common.Logger.Info("剪切板文本监听失败")
				return
			}
			// 剥离数据
			pasteData.Type = "text"
			pasteData.Content = dataText
		}
		// 将id置为0，目前这里出现了一个问题，莫名其妙id为3，所以这里强制置为0
		pasteData.Id = 0
		// 获取该内容是否存在于数据库中
		pasteDb := db.FindPasteDataByContent(pasteData.Content)
		if pasteDb.Id != 0 {
			pasteData.Id = pasteDb.Id
			pasteData.FromApp = pasteDb.FromApp
		}
		// 如果不存在那么就需要重新构建数据
		pasteData.CreatedAt = dateKit.Now()
		db.SaveOrUpdatePaste(pasteData)
		kit.PlaySound("copy")
		//common.Logger.Info("Clipboard changed:", zap.String("%v", string(pasteData.Content)))
	}
}
