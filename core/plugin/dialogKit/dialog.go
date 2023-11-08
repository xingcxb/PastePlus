package dialogKit

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type DialogType string

const (
	Info     DialogType = "info"     // 普通消息对话框
	Question DialogType = "question" // 询问消息对话框
	Warning  DialogType = "warning"  // 警告消息对话框
	Error    DialogType = "error"    // 错误消息对话框
)

// PackageTipsDialog 封包提示对话框
/*
 * @param dialogType 对话框类型
 * @param title 对话框标题
 * @param message 对话框消息
 * @param buttonLeft 左侧按钮
 * @param buttonRight 右侧按钮
 * @param fn 回调函数
 */
func PackageTipsDialog(dialogType DialogType, title, message string) {
	var dialogModel *application.MessageDialog
	switch dialogType {
	case Info:
		// 信息对话框
		dialogModel = application.InfoDialog()
	case Warning:
		// 警告对话框
		dialogModel = application.WarningDialog()
	default:
		// 错误对话框
		dialogModel = application.ErrorDialog()
	}
	// 设置对话框消息标题
	dialogModel.SetTitle(title)
	// 设置对话框消息内容
	dialogModel.SetMessage(message)
	// 显示对话框
	dialogModel.Show()
}

// PackageTipsMutualDialog 封包提示对话框
/*
 * @param title 对话框标题
 * @param message 对话框消息
 * @param buttonLeft 左侧按钮
 * @param buttonRight 右侧按钮
 * @param fn 回调函数
 */
func PackageTipsMutualDialog(title, message, buttonLeftStr, buttonRightStr string, fn func()) {
	dialogModel := application.QuestionDialog()
	var leftButton *application.Button
	if fn != nil {
		// 设置对话框左侧按钮
		leftButton = dialogModel.AddButton(buttonLeftStr).OnClick(fn)
	} else {
		leftButton = dialogModel.AddButton(buttonLeftStr)
	}
	// 设置左侧为默认按钮
	dialogModel.SetDefaultButton(leftButton)
	// 设置对话框右侧按钮
	dialogModel.AddButton(buttonRightStr)
	// 设置对话框消息标题
	dialogModel.SetTitle(title)
	// 设置对话框消息内容
	dialogModel.SetMessage(message)
	// 显示对话框
	dialogModel.Show()
}
