package go_log

import (
	"go.uber.org/zap"
)

type ZapLoggerConfig struct {
	LogFilePath string
	Env         EnvType
}

func (config *ZapLoggerConfig) InitLogger() *zap.Logger {
	var logger *zap.Logger
	if config.Env == EnvDev {
		log := &ConsoleLogger{}
		logger = newLog(log, *config)
	} else {
		log := &FileLogger{}
		logger = newLog(log, *config)
	}
	return logger
}

func newLog(log logger, config ZapLoggerConfig) *zap.Logger {
	return log.NewLogComponent(config)
}

type logger interface {
	NewLogComponent(config ZapLoggerConfig) *zap.Logger
}
