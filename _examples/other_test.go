package _examples

import (
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"context"
	"fmt"
	"github.com/xingcxb/goKit/core/fileKit"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOsFilePath(t *testing.T) {
	fmt.Println(os.Executable())
}

func TestFilePath(t *testing.T) {
	path := "/Users/symbol/Desktop/goproject/v3fanyi/bin/PastePlus.app/Contents/MacOS/qoq"

	// 从后面查找PastePlus.app的位置
	idx := strings.LastIndex(path, "PastePlus.app")
	// 如果找到了
	if idx != -1 {
		// 取到PastePlus.app之前的部分
		result := path[:idx]
		fmt.Println(result)
	}
}

func TestPlaySound(t *testing.T) {
	kit.PlaySound("copy")
}

func TestPath(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Replace(dir, "\\", "/", -1))
	fmt.Println(fileKit.GetCurrentAbPath())
	fmt.Println(fileKit.HomeDir(context.Background()))
	kit.InitLog()
}

func TestLogFile(t *testing.T) {
	kit.InitLog()
	common.Logger.Info("测试日志")
	common.Logger.Debug("测试日志")
	common.Logger.Error("测试日志")
	zap.L().Info("全局日志测试")
}
