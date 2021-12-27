package main

import (
	go_log "github.com/baoyxing/go-log"
)

func main() {
	cofig := go_log.ZapLoggerConfig{LogFilePath: go_log.DefalutLogPath, Env: go_log.EnvTest}
	logger := cofig.InitLogger()
	logger.Info("test")

}
