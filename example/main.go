package main

import go_log "github.com/baoyxing/go-log"

var logger = go_log.InitLogger("", go_log.EnvDev)

func main() {
	logger.Info("test")
}
