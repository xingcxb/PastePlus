// Package tray 系统托盘
package tray

import (
	"PastePlus/core/basic"
	"fmt"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// CreateSysTray 创建系统托盘
func CreateSysTray(app *application.App) *application.Menu {
	trayMenu := app.NewMenu()
	// 打开主面板
	trayMenu.Add("Paste+").OnClick(func(ctx *application.Context) {
		app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
			Title: "打开",
			Mac: application.MacWindow{
				InvisibleTitleBarHeight: 50,
				Backdrop:                application.MacBackdropTranslucent,
				TitleBar:                application.MacTitleBarHiddenInset,
			},
			URL: "#/home",
		})
	})
	// 设置分割线
	trayMenu.AddSeparator()
	// 设置
	trayMenu.Add("设置").OnClick(func(ctx *application.Context) {
		app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
			Title: "打开",
			Mac: application.MacWindow{
				InvisibleTitleBarHeight: 50,
				Backdrop:                application.MacBackdropTranslucent,
				TitleBar:                application.MacTitleBarHiddenInset,
			},
			URL: "#/settings",
		})
	})
	// 设置分割线
	trayMenu.AddSeparator()
	// 帮助
	subMenu := trayMenu.AddSubmenu("帮助")
	subMenu.Add("作者主页").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppAuthorSite)
	})
	subMenu.Add("帮助中心").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppAuthorSite)
	})
	subMenu.Add("功能请求").OnClick(func(ctx *application.Context) {
		browser.OpenURL(basic.AppGithubIssues)
	})
	// 设置分割线
	trayMenu.AddSeparator()
	// 关于
	trayMenu.Add("⏸️暂停").OnClick(func(ctx *application.Context) {
		fmt.Print("======关于点击")
	})
	trayMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	})
	return trayMenu
}
