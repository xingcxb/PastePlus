package customEvents

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/kit/boot"
	"PastePlus/core/plugin/db"
	"PastePlus/core/plugin/dialogKit"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/xingcxb/goKit/core/numKit"
	"github.com/xingcxb/goKit/core/strKit"
	"go.uber.org/zap"
	"golang.design/x/clipboard"
	"os/exec"
	"runtime"
)

type CustomApp struct {
	App *application.App
}

// LoadPasteConfig 加载配置文件
func (a *CustomApp) LoadPasteConfig() {
	a.App.Events.On(common.EventsHandLoadPasteConfigToCore, func(e *application.WailsEvent) {
		// 获取配置信息
		configs, err := db.FindAllConfig()
		if err != nil {
			return
		}
		configData := make(map[string]string, 0)
		for _, config := range configs {
			configData[config.Key] = config.Value
		}
		common.Logger.Info("加载配置文件", zap.String("configData", fmt.Sprintf("%v", configData)))
		//marshal, _ := json.Marshal(configs)
		a.App.Events.Emit(&application.WailsEvent{
			Name: common.EventsHandLoadPasteConfigToFrontend,
			Data: configData,
		})
	})
}

// HandleCleanAllHistoryData 清空所有历史剪贴板自定义事件
func HandleCleanAllHistoryData(app *application.App) {
	app.Events.On(common.EventsHandleCleanAllPasteHistoryToCore, func(e *application.WailsEvent) {
		_, err := db.ResetPasteHistoryData()
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "清空历史剪贴板数据失败")
			return
		}
		dialogKit.PackageTipsDialog(dialogKit.Info, "成功", "清空历史剪贴板数据成功")
	})
}

// SetBootUp 设置开机启动自定义事件
/*
 * @param app app基础
 */
func SetBootUp(app *application.App) {
	app.Events.On(common.EventsHandleBootUpToCore, func(e *application.WailsEvent) {
		// 设置开机启动
		value, err := boot.SetAppBootUp()
		app.Events.Emit(&application.WailsEvent{
			Name: common.EventsHandleBootUpToFrontend,
			Data: value,
		})
		if err != nil {
			common.Logger.Error("设置开机启动失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "设置开机启动失败")
		} else {
			dialogKit.PackageTipsDialog(dialogKit.Info, "成功", "设置开机启动成功")
		}
	})
}

// FindPasteHistory 查找历史剪贴板自定义事件
/*
 * @param app app基础
 */
func FindPasteHistory(app *application.App) {
	app.Events.On(common.EventsFindPasteHistoryToCore, func(e *application.WailsEvent) {
		// 从数据库中获取所有的历史记录
		pasteGoList := db.FindPasteDateByAll()
		// 将go的数据转换为vue的数据
		pasteVueList := make([]common.PasteHistoryVue, 0)
		for _, pasteGo := range pasteGoList {
			// 构建类似IM工具显示时间格式
			spacingTime := kit.FormatWeChatTimeStr(pasteGo.CreatedAt)
			// 格式化输出
			pasteContent := formatContent(pasteGo.Content, pasteGo.Type)
			// 填充结构体数据
			pasteVue := common.PasteHistoryVue{
				Id:          pasteGo.Id,        // 剪贴板id
				Content:     pasteContent,      // 剪贴板内容
				Type:        pasteGo.Type,      // 剪贴板类型
				FromApp:     pasteGo.FromApp,   // 剪贴板来源
				SpacingTime: spacingTime,       // 与现在间隔时间，人可阅读模式
				CreatedAt:   pasteGo.CreatedAt, // 剪贴板创建时间
			}
			pasteVueList = append(pasteVueList, pasteVue)
		}
		pastesJsonByte, err := json.Marshal(pasteVueList)
		if err != nil {
			common.Logger.Error("序列化历史剪贴板数据失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "序列化历史剪贴板数据失败")
			return
		}
		app.Events.Emit(&application.WailsEvent{
			Name: common.EventsFindPasteHistoryToFrontend,
			Data: string(pastesJsonByte),
		})
	})
}

// HandleCardClick 单击卡片绑定事件
/*
 * @param app app基础
 * @param window 当前窗口
 */
func HandleCardClick(app *application.App, window *application.WebviewWindow) {
	app.Events.On(common.EventsHandleCardClickToCore, func(e *application.WailsEvent) {
		// 查询数据
		pasteData := db.FindPasteById(func() int {
			return int(e.Data.(float64))
		}())
		// 如果没有获取到数据，就直接返回
		if pasteData.Id == 0 {
			return
		}
		// 将数据写入到粘贴板中
		clipboard.Write(func() clipboard.Format {
			switch pasteData.Type {
			case "image":
				return clipboard.FmtImage
			default:
				// 默认文本格式
				return clipboard.FmtText
			}
		}(), pasteData.Content)
	})
}

// HandleCardDoubleClick 双击卡片绑定事件
/*
 * @param app app基础
 * @param window 当前窗口
 */
func HandleCardDoubleClick(app *application.App, window *application.WebviewWindow, actionPid int) {
	app.Events.On(common.EventsHandleCardDoubleClickToCore, func(e *application.WailsEvent) {
		// 查询数据
		pasteData := db.FindPasteById(func() int {
			return int(e.Data.(float64))
		}())
		// 如果没有获取到数据，就直接返回
		if pasteData.Id == 0 {
			return
		}
		// 关闭窗口
		window.Close()
		// 格式化输出
		pasteContent := formatContent(pasteData.Content, pasteData.Type)
		if actionPid != 0 {
			pidExists, err := robotgo.PidExists(actionPid)
			if err != nil || !pidExists {
				common.Logger.Error("查询程序是否存在失败", zap.Error(err))
				return
			}
			// 如果上一个应用的pid不是程序自身，那么切换到下一个程序聚焦
			err = activeWindow(actionPid) // robotgo.ActivePid(actionPid)
			if err != nil {
				common.Logger.Error("激活程序失败", zap.Error(err))
				dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", strKit.Splicing("激活程序失败", err.Error()))
				return
			}
			if pasteData.Type == "text" {
				robotgo.MilliSleep(300)
				robotgo.TypeStr(pasteContent)
			}
			// 窗口关闭时，关闭当前窗口上所有的绑定事件
			app.Events.Off(common.EventsHandleCardDoubleClickToCore)
			app.Events.Off(common.EventsHandleCardClickToCore)
			app.Events.Off(common.EventsFindPasteHistoryToCore)
		}
	})
}

func HandOpenUpdate() {

}

// formatContent 格式化输出数据
/*
 * @param contentByte 原始数据
 * @param typeStr 数据类型
 * @return string 格式化后的数据
 */
func formatContent(contentByte []byte, typeStr string) string {
	pasteContent := ""
	if typeStr == "image" {
		pasteContent = base64.StdEncoding.EncodeToString(contentByte)
		pasteContent = strKit.Splicing("data:image/png;base64,", pasteContent)
	} else {
		pasteContent = string(contentByte)
	}
	return pasteContent
}

// activeWindow 激活窗口
/*
 * @param pid int 程序pid
 * @return error 错误信息
 */
func activeWindow(pid int) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		//cmd = exec.Command("osascript", "-e", strKit.Splicing("tell application \"System Events\" to set frontmost of the first process whose unix id is ", numKit.Int2Str(pid)), "to true")
		cmd = exec.Command("osascript", "-e", fmt.Sprintf(`tell application "System Events" to set frontmost of every process whose unix id is %d to true`, pid))
	case "windows":
		cmd = exec.Command("cmd", "/c", strKit.Splicing("wmic process where processid=", numKit.Int2Str(pid), " call setforeground"))
	case "linux":
		cmd = exec.Command("xdotool", "windowactivate", strKit.Splicing("--sync ", numKit.Int2Str(pid)))
	}
	return cmd.Run()
}
