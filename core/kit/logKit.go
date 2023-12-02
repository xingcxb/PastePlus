package kit

import (
	"PastePlus/core/basic/common"
	"context"
	"github.com/xingcxb/goKit/core/dateKit"
	"github.com/xingcxb/goKit/core/strKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLog 初始化日志服务
func InitLog() {
	// 获取当前系统的主目录
	homeDir, _ := HomeDir(context.Background())
	filePath := strKit.Splicing(homeDir, "/log/PastePlus-", dateKit.Today(), ".log")
	// 配置日志编码器
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	// 创建日志核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(openLogFile(filePath)),
		atomicLevel,
	)
	// 构建日志记录器
	common.Logger = zap.New(core)
	defer common.Logger.Sync()
	// 将日志记录器替换为全局日志记录器
	zap.ReplaceGlobals(common.Logger)
}

func openLogFile(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
