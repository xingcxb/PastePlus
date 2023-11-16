//go:build windows
// +build windows

package pid

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetPid 获取pid
func GetPid() int {
	psScript := `(Get-Process -id (Get-WindowThreadProcessId (Get-ForegroundWindow)))[0].Id`
	cmd := exec.Command("powershell", "-Command", psScript)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// 转换输出为字符串并去除可能的换行符
	pidStr := strings.TrimSpace(string(output))

	// 将字符串转换为 int
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println("Error converting PID to integer:", err)
		return 0
	}

	fmt.Printf("Active Window PID: %d\n", pid)
	return pid
}
