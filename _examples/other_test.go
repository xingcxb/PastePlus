package _examples

import (
	"fmt"
	"os"
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
