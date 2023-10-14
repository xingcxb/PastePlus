// Package bindHotKey 绑定快捷键
package bindHotKey

import (
	"PastePlus/core/window"
	"fmt"
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
		hotKeysCV := []string{"command", "c"}
		if runtime.GOOS != "darwin" {
			// 如果不是Mac，则使用ctrl+c
			hotKeysCV = []string{"ctrl", "c"}
		}
		// 注册command+c快捷键 暂时采用此方法
		hook.Register(hook.KeyDown, hotKeysCV, func(e hook.Event) {
			// todo 暂时只能处理文本内容
			text, b := app.Clipboard().Text()
			if b {
				fmt.Println("=====>", text)
			} else {
				fmt.Println("剪贴板无内容")
			}
		})
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
