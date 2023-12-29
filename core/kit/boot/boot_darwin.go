//go:build darwin

package boot

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"errors"
	"github.com/xingcxb/goKit/core/strKit"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
)

const launchctl = "/bin/launchctl"

// SetAppBootUp 设置开机启动
func SetAppBootUp() (bool, error) {
	// 获取运行程序的位置
	path, err := os.Executable()
	common.Logger.Info("获取运行程序的位置", zap.String("path:", path))
	if err != nil {
		return false, err
	}
	// 从后面查找PastePlus.app的位置
	idx := strings.LastIndex(path, "PastePlus.app")
	appPath := ""
	// 如果找到了
	if idx != -1 {
		appPath = path[:idx]
	} else {
		return false, errors.New("未找到PastePlus.app的位置")
	}
	common.Logger.Info("获取PastePlus.app的位置", zap.String("appPath:", appPath))
	status, err := checkStartupStatus()
	if err != nil {
		return false, err
	}
	if status {
		// 禁用开机启动
		return disableStartup()
	} else {
		// 启用开机启动
		return enableStartup(appPath)
	}

}

// enableStartup 启用开机启动
func enableStartup(appPath string) (bool, error) {
	cmd := exec.Command(launchctl, "enable", "gui/"+os.Getenv("UID"), strKit.Splicing(appPath, "/PastePlus.app"))
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}

// DisableStartup 取消开机启动
func disableStartup() (bool, error) {
	cmd := exec.Command(launchctl, "disable", "gui/"+os.Getenv("UID"), "user/"+os.Getenv("UID")+"/"+basic.AppName)
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}

// 检查应用程序是否已添加到开机启动项
func checkStartupStatus() (bool, error) {
	cmd := exec.Command(launchctl, "list")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return strings.Contains(string(output), basic.AppName), nil
}
