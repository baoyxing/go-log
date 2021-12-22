package go_log

const DefalutLogPath = "app.log"
type EnvType int
const (
	//开发环境
	EnvDev EnvType = iota
	//测试环境
	envTest
	//预生产环境
	envPrePro
	//生产环境
	envPro

)