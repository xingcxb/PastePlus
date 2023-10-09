package main

import (
	"PastePlus/core/basic"
	"PastePlus/core/plugin/cron"
	"PastePlus/core/window/tray"
	"embed"
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        basic.AppName,
		Description: basic.AppDescription,
		Assets: application.AssetOptions{
			FS: assets,
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// 创建主窗口，该窗口需要保持，不能退出，不然程序就退出了
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Mac: application.MacWindow{
			DisableShadow:           false,
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
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
		ShouldClose: func(window *application.WebviewWindow) bool {
			window.Hide()
			return false
		},
		Windows: application.WindowsWindow{
			HiddenOnTaskbar: true,
		},
		URL: "#/home",
	})
	//创建托盘
	systemTray := app.NewSystemTray()
	b, _ := assets.ReadFile("frontend/dist/logoX64.png")
	systemTray.SetTemplateIcon(b)
	// 创建系统托盘菜单
	trayMenu := tray.CreateSysTray(app)
	// 设置托盘菜单
	systemTray.SetMenu(trayMenu)
	//systemTray.AttachWindow(coreWindow).WindowOffset(5)

	go func() {
		// 创建定时任务
		cron.CreateCron()
	}()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
