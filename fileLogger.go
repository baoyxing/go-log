package go_log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogger struct {
}

func (*FileLogger) NewLogComponent(config ZapLoggerConfig) *zap.Logger {
	var logger *zap.Logger
	switch config.Env {
	case EnvTest:
		fileLogger := &testFileLogger{}
		logger = initFileLogger(fileLogger, config, zapcore.InfoLevel)
	case EnvPrePro, EnvPro:
		fileLogger := &proFileLogger{}
		logger = initFileLogger(fileLogger, config, zapcore.ErrorLevel)

	}
	return logger
}

type fileLoggerInterface interface {
	newCore(config ZapLoggerConfig, Level zapcore.Level) zapcore.Core
	newLoggger(core zapcore.Core) *zap.Logger
}

func (*FileLogger) newCore(config ZapLoggerConfig, Level zapcore.Level) zapcore.Core {
	if config.LogFilePath == "" {
		config.LogFilePath = DefalutLogPath
	}
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   config.LogFilePath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,                 // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,                 // 保留30个备份，默认不限
		MaxAge:     7,                  // 保留7天，默认不限
		Compress:   true,               // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
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
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		Level,
	)
	return core
}

type proFileLogger struct {
	FileLogger
}

func (p *proFileLogger) newLoggger(core zapcore.Core) *zap.Logger {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.ErrorLevel)
	core.Enabled(zap.ErrorLevel)
	return zap.New(core)
}

type testFileLogger struct {
	FileLogger
}

func (t *testFileLogger) newLoggger(core zapcore.Core) *zap.Logger {
	//Enable stack trace
	caller := zap.AddCaller()
	//Open file and line number
	development := zap.Development()
	return zap.New(core, caller, development)
}
func initFileLogger(f fileLoggerInterface, config ZapLoggerConfig, Level zapcore.Level) *zap.Logger {
	return f.newLoggger(f.newCore(config, Level))
}
