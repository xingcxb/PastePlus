package main

import (
	"PastePlus/core/api/bindings"
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/plugin/cron"
	"PastePlus/core/plugin/db"
	"PastePlus/core/window/bindHotKey"
	"PastePlus/core/window/hook"
	"PastePlus/core/window/tray"
	"embed"
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        basic.AppName,
		Description: basic.AppDescription,
		// 设置任务栏icon，默认关于等其它原生对话框所使用的icon
		Icon: func() []byte {
			b, _ := assets.ReadFile("frontend/dist/logoX800.png")
			return b
		}(),
		Bind: []any{
			// 绑定前端的api，目前这里还有bug，暂时不要使用
			&bindings.SetService{},
		},
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
	b, _ := assets.ReadFile("frontend/dist/icon/icons.icns")
	// 设置托盘图标，彩色的
	systemTray.SetIcon(b)
	// 设置托盘图标，单色的
	//systemTray.SetTemplateIcon(b)
	// 创建系统托盘菜单
	trayMenu := tray.CreateSysTray(app)
	// 设置托盘菜单
	systemTray.SetMenu(trayMenu)

	// 用于追加插件
	go func() {
		// 初始化日志
		kit.InitLog()
		// 初始化声音播放oto上下文
		kit.InitOto()
		// 初始化数据库
		db.InitDb()
		if !db.Sqlite3Status {
			// 数据库启动失败关闭程序
			app.Quit()
		}
		// 绑定全局热键
		bindHotKey.BindingGlobalHotkey(app)
		// 启动监听
		hook.RegexListen()
		// 创建定时任务
		cron.CreateCron()
	}()
	err := app.Run()

	if err != nil {
		common.Logger.Error("程序启动失败", zap.Error(err))
	}
}
