// Package hook 钩子操作
package hook

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

func WindowClose(window *application.WebviewWindow) {
	window.RegisterHook(events.Common.WindowClosing, func(e *application.WindowEvent) {
		go func() {
			window.Close()
		}()
	})
}

// WindowLostFocusHide 窗口失去焦点窗口隐藏
func WindowLostFocusHide(window *application.WebviewWindow) {
	window.RegisterHook(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		// 为了避免隐藏窗体的时候出现卡顿的情况，使用协程处理
		go func() {
			window.Hide()
		}()
	})
}

// WindowLostFocusClose 窗口失去焦点窗口关闭
func WindowLostFocusClose(window *application.WebviewWindow) {
	window.RegisterHook(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		go func() {
			window.Close()
		}()
	})
}
