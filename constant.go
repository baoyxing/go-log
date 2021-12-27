package go_log

const DefalutLogPath = "app.log"

type EnvType int

const (
	//开发环境
	EnvDev EnvType = iota
	//测试环境
	EnvTest
	//预生产环境
	EnvPrePro
	//生产环境
	EnvPro
)
