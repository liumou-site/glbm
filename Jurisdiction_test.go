package glbm

import (
	"testing"

	"gitee.com/liumou_site/logger"
)

// 检查开发者模式
func TestDeveloper(t *testing.T) {
	if Developer() {
		logger.Info("已开启开发者模式")
	} else {
		logger.Error("未开启开发者模式")
	}
}

// 检查Sudo
func TestCheckSudo(t *testing.T) {
	if CheckSudo("1") {
		logger.Info("Sudo权限获取成功")
	} else {
		logger.Error("SUdo权限获取失败")
	}
	if CheckSudo("2") {
		logger.Info("Sudo权限获取成功")
	} else {
		logger.Error("SUdo权限获取失败")
	}
}
