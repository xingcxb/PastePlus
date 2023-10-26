// Package bindHotKey 绑定快捷键
package bindHotKey

import (
	"PastePlus/core/window"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
	"runtime"
)

var (
	HotKeysCVHook chan hook.Event
)

// MainWindowHotKey 绑定主窗口快捷键
func MainWindowHotKey(app *application.App) {
	go func() {
		// 注册command+shift+v快捷键
		hotKeysCSV := []string{"command", "shift", "z"}
		if runtime.GOOS != "darwin" {
			// 如果不是Mac，则使用ctrl+shift+v
			hotKeysCSV[0] = "ctrl"
		}
		hook.Register(hook.KeyDown, hotKeysCSV, func(e hook.Event) {
			// 启动窗口
			window.MainWindow(app)
		})
		HotKeysCVHook = hook.Start()
		<-hook.Process(HotKeysCVHook)
	}()
}
