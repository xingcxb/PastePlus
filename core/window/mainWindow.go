// Package window 主窗口
package window

import (
	"PastePlus/core/window/hook"
	"github.com/wailsapp/wails/v3/pkg/application"
)

const (
	mainWindowName = "MainWindow"
)

var (
	// 默认主窗口宽度
	mainWindowWidth = 1920
	// 默认主窗口高度
	mainWindowHeight = 250
)

// MainWindow 主窗口
func MainWindow(app *application.App) {
	if w, ok := app.GetWindowByName(mainWindowName).(*application.WebviewWindow); ok {
		// 判断如果当前的窗口已经存在，则显示并聚焦
		w.Show().Focus()
		return
	}
	// 获取当前主窗口的名字
	mainWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		// 设置窗口名称
		Name: mainWindowName,
		Mac: application.MacWindow{
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
		// 设置窗口内容
		URL:    "#/home",
		Width:  mainWindowWidth,
		Height: mainWindowHeight,
	})
	// 设置窗口位置
	// (0,0)为窗口左下角，默认位置(0,604)
	// 注意，设置窗口未知尽量在声明了窗口后，立即设置，不然会出现窗口闪烁的情况
	mainWindow.SetAbsolutePosition(0, 0)
	// 设置为显示聚焦
	mainWindow.Show().Focus()
	// 获取屏幕，该函数调用必须在app.Run()之后
	screen, _ := app.GetPrimaryScreen()
	// 覆盖默认屏幕的宽度
	mainWindowWidth = screen.Size.Width
	// 设置窗口大小
	mainWindow.SetSize(mainWindowWidth, mainWindowHeight)
	// 窗口失去焦点时隐藏窗口
	hook.WindowLostFocus(mainWindow)
	//hook.WindowLostFocusTest(app)
	return
}

// SetWindow 设置窗口
func SetWindow(app *application.App) *application.WebviewWindow {

	return nil
}
