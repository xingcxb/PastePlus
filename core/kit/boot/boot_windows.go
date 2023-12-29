//go:build windows

package boot

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"github.com/xingcxb/goKit/core/fileKit"
	"go.uber.org/zap"
	"golang.org/x/sys/windows/registry"
	"path/filepath"
)

// SetAppBootUp 设置开机启动
func SetAppBootUp() (bool, error) {
	_appRunPath := fileKit.GetCurrentAbPath()
	appRunPath := filepath.Clean(_appRunPath)
	// 打开注册表的 Run 键
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		common.Logger.Error("Failed to open Run key", zap.Error(err))
		return false, err
	}
	defer k.Close()
	// 在 Run 键下创建一个新的字符串值
	err = k.SetStringValue(basic.AppName, appRunPath)
	if err != nil {
		common.Logger.Error("Failed to set registry value", zap.Error(err))
		return false, err
	}
	return true, nil
}
