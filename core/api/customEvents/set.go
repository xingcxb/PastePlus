package customEvents

import (
	"PastePlus/core/basic/common"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func FindPasteHistory(app *application.App) {
	app.Events.On(common.EventsFindPasteHistoryToCore, func(e *application.WailsEvent) {
		pastes := make([]common.PasteHistory, 0)
		//todo 从数据库中获取历史剪贴板数据
		//此处是假数据
		pastes = append(pastes, common.PasteHistory{
			Id:        1,
			FromApp:   "QQ",
			Content:   "这是一条来自QQ数据",
			Type:      "text",
			CreatedAt: "2021-01-01 00:00:00",
		})
		pastes = append(pastes, common.PasteHistory{
			Id:        2,
			FromApp:   "微信",
			Content:   "这是一条来自微信数据",
			Type:      "text",
			CreatedAt: "2021-01-01 00:00:00",
		})
		fmt.Println("12312===>", e.Data, e.Name)
		pastesJsonByte, err := json.Marshal(pastes)
		if err != nil {
			fmt.Println("==!!!!==>", err)
			return
		}
		fmt.Println(string(pastesJsonByte))
		app.Events.Emit(&application.WailsEvent{
			Name: common.EventsFindPasteHistoryToFrontend,
			Data: string(pastesJsonByte),
		})
	})

}

func Test() string {
	return "你好啊2"
}
