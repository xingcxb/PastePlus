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
		//Icon: func() []byte {
		//	b, _ := assets.ReadFile("frontend/dist/logoX800.png")
		//	return b
		//}(),
		Assets: application.AssetOptions{
			FS: assets,
		},
		Mac: application.MacOptions{
			// 如果要让macOS的Dock不显示这里必须指定
			// 目前还未知晓ActivationPolicyAccessory和ActivationPolicyProhibited有什么区别
			ActivationPolicy: application.ActivationPolicyAccessory,
			// 允许所有的窗口关闭，但是程序依旧保持运行
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
	})

	//创建托盘
	systemTray := app.NewSystemTray()
	b, _ := assets.ReadFile("frontend/dist/logoX32.png")
	systemTray.SetTemplateIcon(b)
	// 创建系统托盘菜单
	trayMenu := tray.CreateSysTray(app)
	// 设置托盘菜单
	systemTray.SetMenu(trayMenu)

	//window.MainWindow(app)
	//systemTray.AttachWindow(window.MainWindow(app)).WindowOffset(5)

	go func() {
		// 创建定时任务
		cron.CreateCron()
	}()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}