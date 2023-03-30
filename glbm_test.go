package glbm

import (
	"fmt"
	"testing"

	"gitee.com/liumou_site/logger"
)

// 权限功能测试
func TestJurisdiction(t *testing.T) {
	g := CheckSudo("1")
	if g {
		fmt.Println("密码检验正确")
	} else {
		fmt.Println("密码错误或无权限")
	}
	d := Developer()
	if d {
		fmt.Println("已开启开发者")
	} else {
		fmt.Println("未开启开发者")
	}
}

// 用户信息测试
func TestGetUserInfo(t *testing.T) {
	get, err := GetUserInfo()
	if err == nil {
		logger.Info("username: ", get.Username)
		logger.Info("uid: ", get.Uid)
		logger.Info("u_home: ", get.HomeDir)
	}
}

func TestVersion(t *testing.T) {
	Version()
}
func TestOsInfo(t *testing.T) {
	info, err := GetOsInfo()
	if err == nil {
		logger.Info("系统类型: ", info.Name)
		logger.Info("系统架构: ", info.Arch)
		logger.Info("系统版本: ", info.Version)
		logger.Info("系统代号: ", info.CodeName)
	} else {
		logger.Error(err.Error())
	}

}
