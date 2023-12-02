//go:build darwin

package pid

import (
	"PastePlus/core/basic/common"
	"go.uber.org/zap"
	"os/exec"
	"strconv"
	"strings"
)

// GetPid 获取pid
func GetPid() (int, error) {
	appleScript := `tell application "System Events" to unix id of first process whose frontmost is true`
	cmd := exec.Command("osascript", "-e", appleScript)
	output, err := cmd.Output()
	if err != nil {
		common.Logger.Error("获取pid失败", zap.Error(err))
		return 0, err
	}
	// 转换输出为字符串并去除可能的换行符
	pidStr := strings.TrimSpace(string(output))

	// 将字符串转换为 int
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Logger.Error("Error converting PID to integer", zap.Error(err))
		return 0, err
	}
	return pid, nil
}
