package go_log

import (
	"go.uber.org/zap"
)

type ConsoleLogger struct {
}

func (*ConsoleLogger) NewLogComponent(config ZapLoggerConfig) *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
