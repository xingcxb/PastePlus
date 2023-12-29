package kit

import (
	"PastePlus/core/basic/common"
	"context"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/xingcxb/goKit/core/fileKit"
	"github.com/xingcxb/goKit/core/strKit"
	"go.uber.org/zap"
	"os"
	"os/user"
	"path/filepath"
	"time"
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

// InitOto 初始化oto上下文
func InitOto() {
	c, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   44100, // 假设所有音频文件都是 44100 Hz
		ChannelCount: 2,     // 假设所有音频文件都是立体声
		Format:       2,     // 假设所有音频文件都是16位
		BufferSize:   8192,
	})
	if err != nil {
		common.Logger.Error("准备Oto上下文失败", zap.Error(err))
		return
	}
	common.AudioContent = c
	<-ready
}

// PlaySound 播放声音
/*
 * @param soundType string 声音类型
 */
func PlaySound(soundType string) {
	// 获取当前运行目录
	runPath := fileKit.GetCurrentAbPath()
	// 向上回一层，资源文件目录和可执行文件目录在同一个目录下
	runPath = filepath.Dir(runPath)
	soundFilePath := ""
	switch soundType {
	case "copy":
		// 播放复制声音
		soundFilePath = strKit.Splicing(runPath, "/Resources/Copy.mp3")
	case "paste":
		// 播放粘贴声音
		soundFilePath = strKit.Splicing(runPath, "/Resources/Paste.mp3")
	}
	// 打开 MP3 文件
	f, err := os.Open(soundFilePath)
	if err != nil {
		common.Logger.Error("打开MP3文件失败", zap.Error(err))
		return
	}
	defer f.Close()
	// 解码 MP3 文件
	d, err := mp3.NewDecoder(f)
	if err != nil {
		common.Logger.Error("解码MP3文件失败", zap.Error(err))
		return
	}
	// 创建播放器
	p := common.AudioContent.NewPlayer(d)
	// 播放音频
	p.Play()
	time.Sleep(1 * time.Second)
	// 等待音频播放完成（可根据需要调整）
	defer p.Close()
}
