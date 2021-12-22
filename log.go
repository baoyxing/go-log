package go_log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(filePath string, env EnvType) {
	if env == EnvDev {
		Logger, _ = zap.NewDevelopment()
	} else {
		initZapLogger(filePath, env)
	}
}

func initZapLogger(filePath string, env EnvType) {
	if filePath == "" {
		filePath = DefalutLogPath
	}
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   filePath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,       // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,       // 保留30个备份，默认不限
		MaxAge:     7,        // 保留7天，默认不限
		Compress:   true,     // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
	var level zapcore.Level
	switch env {
	case envTest:
		level = zap.InfoLevel
	case envPrePro:
	case envPro:
		level = zap.ErrorLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		level,
	)
	switch env {
	case envTest:
		{
			//Enable stack trace
			caller := zap.AddCaller()
			//Open file and line number
			development := zap.Development()
			Logger = zap.New(core, caller, development)
		}
	case envPrePro:
	case envPro:
		{
			Logger = zap.New(core)
		}
	}
}
