package kit

import (
	"context"
	"os"
	"os/user"
	"path/filepath"
)

// HomeDir 获取系统当前使用的用户的主目录
/*
 * eg: /Users/symbol
 * @param ctx 上下文
 * @return string,error
 */
func HomeDir(ctx context.Context) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}

// CreateLazyFile 一次性创建好文件和文件夹
/*
 * @param filePath 文件路径
 */
func CreateLazyFile(filePath string) error {
	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// PlaySound 播放声音
/*
 * @param soundType string 声音类型
 */
func PlaySound(soundType string) {
	switch soundType {
	case "copy":
	// 播放复制声音
	case "paste":
		// 播放粘贴声音
	}
}
