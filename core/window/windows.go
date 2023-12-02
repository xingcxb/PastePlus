// Package window 主窗口
package window

import (
	"PastePlus/core/api/customEvents"
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/kit/pid"
	"PastePlus/core/plugin/dialogKit"
	"PastePlus/core/window/hook"
	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/xingcxb/goKit/core/strKit"
	"runtime"
	"strconv"
)

var (
	// 默认主窗口宽度
	mainWindowWidth = 1920
	// 默认主窗口高度
	mainWindowHeight = 258
)

// MainWindow 主窗口
func MainWindow(app *application.App) {
	actionPid, err := pid.GetPid()
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "获取当前激活的窗口的Pid失败")
		return
	}
	if w, ok := app.GetWindowByName(common.MainWindowName).(*application.WebviewWindow); ok {
		// 判断如果当前的窗口已经存在，则显示
		actionPid = kit.CheckPid(robotgo.GetPid(), app.GetPID())
		w.Show().Focus()
		return
	}
	common.Logger.Info(strKit.Splicing("当前激活的窗口的Pid：", strconv.Itoa(actionPid)))
	actionPid = kit.CheckPid(robotgo.GetPid(), app.GetPID())
	// 获取屏幕，该函数调用必须在app.Run()之后
	screen, _ := app.GetPrimaryScreen()
	// 覆盖默认屏幕的宽度
	mainWindowWidth = screen.Size.Width
	// 获取当前主窗口的名字
	mainWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		// 设置窗口名称
		Name: common.MainWindowName,
		// todo 暂时只对Mac进行优化
		Mac: application.MacWindow{
			Backdrop:      application.MacBackdropTranslucent,
			DisableShadow: false,
			TitleBar: application.MacTitleBar{
				AppearsTransparent:   true,
				Hide:                 true,
				HideTitle:            true,
				FullSizeContent:      true,
				UseToolbar:           true,
				HideToolbarSeparator: true,
				ToolbarStyle:         application.MacToolbarStyleUnified,
			},
		},
		BackgroundType: application.BackgroundTypeTransparent, // 窗口透明
		BackgroundColour: application.RGBA{
			Red:   0,
			Green: 0,
			Blue:  0,
			Alpha: 0,
		},
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"escape": func(window *application.WebviewWindow) {
				window.Close()
			},
		},
		//AlwaysOnTop:   true,                        // 窗口置顶，该操作不会让其它程序失去焦点
		Focused:       true,                        // 窗口失去焦点
		DisableResize: true,                        // 禁止窗口缩放
		URL:           common.MainWindowContentUrl, // 设置窗口内容
		Width:         mainWindowWidth,             // 设置宽度
		Height:        mainWindowHeight,            // 设置高度
	})
	// 自定义事件，查询历史数据
	customEvents.FindPasteHistory(app)
	// 自定义事件，单击卡片事件
	customEvents.HandleCardClick(app, mainWindow)
	// 自定义事件，双击卡片事件
	customEvents.HandleCardDoubleClick(app, mainWindow, actionPid)

	// 设置窗口位置
	if runtime.GOOS == "darwin" {
		// (0,0)为Mac窗口左下角，默认位置(0,604)
		// 注意，设置窗口未知尽量在声明了窗口后，立即设置，不然会出现窗口闪烁的情况
		mainWindow.SetAbsolutePosition(0, 0)
	} else {
		// 如果是非macOS系统，假定原点坐标为左上角
		mainWindow.SetAbsolutePosition(0, -mainWindowHeight)
	}
	// 设置为显示聚焦
	mainWindow.Show().Focus()
	mainWindow.SetAlwaysOnTop(true)
	// 强制刷新
	mainWindow.ForceReload()
	// 设置窗口大小
	mainWindow.SetSize(mainWindowWidth, mainWindowHeight)
	// 窗口失去焦点时关闭窗口
	hook.WindowLostFocusClose(mainWindow)
	return
}

// SettingsWindow 设置窗口
/*
 * @Description: 设置窗口
 * @param app *application.App 应用程序
 * @param url string 窗口内容地址
 */
func SettingsWindow(app *application.App, url string) {
	if url == "" {
		url = common.SettingsWindowContentUrl
	}
	if w, ok := app.GetWindowByName(common.SettingsWindowName).(*application.WebviewWindow); ok {
		// 判断如果当前的窗口已经存在，则显示并聚焦
		w.Show().Focus()
		return
	}
	// 获取当前主窗口的名字
	settingsWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		// 设置窗口名称
		Name: common.SettingsWindowName,
		// todo 暂时只对Mac进行优化
		Mac: application.MacWindow{
			DisableShadow:           false,
			InvisibleTitleBarHeight: 50,
			TitleBar: application.MacTitleBar{
				AppearsTransparent:   false,
				Hide:                 false,
				HideTitle:            true,
				FullSizeContent:      true,
				UseToolbar:           false,
				HideToolbarSeparator: true,
			},
		},

		DisableResize: true, // 禁止窗口缩放
		URL:           url,  // 设置窗口内容
		Width:         600,  // 设置宽度
		Height:        500,  // 设置高度
	})
	// 绑定自定义事件
	// 绑定设置程序自动启动
	customEvents.SetBootUp(app)
	// 绑定清理粘贴板数据
	customEvents.HandleCleanAllHistoryData(app)
	// 绑定加载配置文件
	customEvents.LoadPasteConfig(app)

	// 设置窗口位置居中
	settingsWindow.Center()
	// 设置为显示聚焦
	settingsWindow.Show().Focus()
	// 窗口失去焦点时关闭窗口
	//hook.WindowLostFocusClose(settingsWindow)
	// 窗口关闭时关闭窗口
	hook.WindowClose(settingsWindow)
	return
}
