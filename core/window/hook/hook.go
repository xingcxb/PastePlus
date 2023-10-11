// Package hook 钩子操作
package hook

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// WindowLostFocus 窗口失去焦点
func WindowLostFocus(window *application.WebviewWindow) {
	window.RegisterHook(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		// 目前这里出现了卡顿的情况
		go func() {
			window.Hide()
		}()
	})
}
