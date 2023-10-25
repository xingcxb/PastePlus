package customEvents

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/plugin/db"
	"PastePlus/core/plugin/dialogKit"
	"encoding/json"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// FindPasteHistory 查找历史剪贴板自定义事件
func FindPasteHistory(app *application.App) {
	app.Events.On(common.EventsFindPasteHistoryToCore, func(e *application.WailsEvent) {
		// 从数据库中获取所有的历史记录
		pastes := db.FindAllPaste()
		pastesJsonByte, err := json.Marshal(pastes)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "序列化历史剪贴板数据失败")
			return
		}
		app.Events.Emit(&application.WailsEvent{
			Name: common.EventsFindPasteHistoryToFrontend,
			Data: string(pastesJsonByte),
		})
	})
}
