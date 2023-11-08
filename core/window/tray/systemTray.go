// Package tray 系统托盘
package tray

import (
	"PastePlus/core/basic"
	"PastePlus/core/kit"
	"PastePlus/core/window"
	"fmt"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/xingcxb/goKit/core/strKit"
)

// CreateSysTray 创建系统托盘
func CreateSysTray(app *application.App) *application.Menu {
	trayMenu := app.NewMenu()
	// 打开主面板
	trayMenu.Add(strKit.Splicing(basic.AppName, "  ", basic.AppVersion)).OnClick(func(ctx *application.Context) {
		window.MainWindow(app)
	}).SetAccelerator("CmdOrCtrl+Shift+V")
	// 设置分割线
	trayMenu.AddSeparator()
	// 设置
	trayMenu.Add("偏好设置").SetAccelerator("CmdOrCtrl+,").OnClick(func(ctx *application.Context) {
		window.SettingsWindow(app, "")
	})
	// 设置分割线
	trayMenu.AddSeparator()
	// 帮助
	subMenu := trayMenu.AddSubmenu("帮助中心")
	subMenu.Add("作者主页").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppAuthorSite)
	})
	subMenu.Add("帮助中心").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppAuthorSite)
	})
	subMenu.Add("功能请求").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppGithubIssues)
	})
	subMenu.Add("检查更新...").OnClick(func(ctx *application.Context) {
		kit.UpdateAppVersion()
	})
	// 设置分割线
	trayMenu.AddSeparator()
	// 关于
	trayMenu.Add("暂停").SetEnabled(false).OnClick(func(ctx *application.Context) {
		fmt.Println("====>暂停功能")
	})
	trayMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	}).SetAccelerator("CmdOrCtrl+Q")
	return trayMenu
}
