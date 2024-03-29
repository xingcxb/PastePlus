//go:build darwin

package boot

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/logic"
	"errors"
	"fmt"
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
	if err != nil {
		return false, err
	}
	// 从后面查找PastePlus.app的位置
	idx := strings.LastIndex(path, basic.AppName+".app")
	appPath := ""
	// 如果找到了
	if idx != -1 {
		appPath = path[:idx]
	} else {
		return false, errors.New("未找到PastePlus.app的位置")
	}
	appPath += "/" + basic.AppName + ".app"
	// 检查是否已添加到开机启动项
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
	fmt.Println("appPath:", appPath)
	applescript := `
tell application "System Events"
    make login item at end with properties {path:"` + appPath + `", hidden:false}
end tell
`
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	if err != nil {
		common.Logger.Error("设置开机启动失败", zap.Error(err))
		return false, err
	}
	logic.HandConfigValue(common.ConfigKeyBootUp, "true")
	return true, nil
}

// DisableStartup 取消开机启动
func disableStartup() (bool, error) {
	applescript := `
tell application "System Events"
    if exists login item "` + basic.AppName + `" then
        delete login item "` + basic.AppName + `"
    end if
end tell
`
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	if err != nil {
		common.Logger.Error("取消开机启动失败", zap.Error(err))
		return false, err
	}
	isOk := logic.HandConfigValue(common.ConfigKeyBootUp, "false")
	common.Logger.Info("see change db status", zap.Bool("isOk:", isOk))
	return true, nil
}

// 检查应用程序是否已添加到开机启动项
func checkStartupStatus() (bool, error) {
	applescript := `
tell application "System Events"
    get the name of every login item
end tell
`
	cmd := exec.Command("osascript", "-e", applescript)
	output, err := cmd.Output()
	if err != nil {
		common.Logger.Error("检查开机启动失败", zap.Error(err))
		return false, err
	}
	common.Logger.Info("检查开机启动", zap.String("output:", string(output)))
	loginItems := strings.Split(string(output), ", ")
	for _, item := range loginItems {
		item = strings.ReplaceAll(item, "\n", "")
		if item == basic.AppName {
			return true, nil
		}
	}
	return false, nil
}
