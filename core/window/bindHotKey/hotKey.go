// Package bindHotKey 绑定快捷键
package bindHotKey

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/window"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
	"runtime"
	"time"
)

var (
	HotKeysCVHook chan hook.Event
)

// BindingGlobalHotkey 绑定主窗口快捷键
func BindingGlobalHotkey(app *application.App) {
	// 注册command+shift+v快捷键，临时使用command+shift+z避免和我本地的软件冲突
	hotKeysCSV := []string{"command", "shift", "v"}
	if runtime.GOOS != "darwin" {
		// 如果不是Mac，则使用ctrl+shift+v
		hotKeysCSV[0] = "ctrl"
	}
	hook.Register(hook.KeyDown, hotKeysCSV, func(e hook.Event) {
		common.Logger.Info("按下快捷键")
		// 启动窗口
		window.MainWindow(app)
	})
	HotKeysCVHook = hook.Start()
	go func() {
		for {
			<-hook.Process(HotKeysCVHook)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
