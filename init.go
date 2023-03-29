package glbm

import "gitee.com/liumou_site/logger"

var logs *logger.LocalLogger // 日志打印

func init() {
	logs = logger.NewLogger(3)
	logs.Version = Version()
	logs.Modular = "glbm"
}
