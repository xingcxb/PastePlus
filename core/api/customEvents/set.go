package customEvents

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/plugin/db"
	"PastePlus/core/plugin/dialogKit"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// FindPasteHistory 查找历史剪贴板自定义事件
func FindPasteHistory(app *application.App) {
	app.Events.On(common.EventsFindPasteHistoryToCore, func(e *application.WailsEvent) {
		// 从数据库中获取所有的历史记录
		pasteGoList := db.FindAllPaste()
		// 将go的数据转换为vue的数据
		pasteVueList := make([]common.PasteHistoryVue, 0)
		for _, pasteGo := range pasteGoList {
			pasteVue := common.PasteHistoryVue{
				Id:        pasteGo.Id,              // 剪贴板id
				Content:   string(pasteGo.Content), // 剪贴板内容
				Type:      pasteGo.Type,            // 剪贴板类型
				FromApp:   pasteGo.FromApp,         // 剪贴板来源
				CreatedAt: pasteGo.CreatedAt,       // 剪贴板创建时间
			}
			pasteVueList = append(pasteVueList, pasteVue)
		}
		pastesJsonByte, err := json.Marshal(pasteVueList)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "序列化历史剪贴板数据失败")
			return
		}
		fmt.Println("=====>", string(pastesJsonByte))
		app.Events.Emit(&application.WailsEvent{
			Name: common.EventsFindPasteHistoryToFrontend,
			Data: string(pastesJsonByte),
		})
	})
}
