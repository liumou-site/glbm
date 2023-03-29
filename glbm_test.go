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
	get, username, uid, uHome := GetUserInfo(false)
	if get {
		logger.Info("username: ", username)
		logger.Info("uid: ", uid)
		logger.Info("u_home: ", uHome)
	}
}

func TestVersion(t *testing.T) {
	Version()
}
func TestOsInfo(t *testing.T) {
	osType, osArch, ov, err := GetOsInfo()
	if err == nil {
		logger.Info("系统类型: ", osType)
		logger.Info("系统架构: ", osArch)
		logger.Info("系统版本: ", ov)
	} else {
		logger.Error(err.Error())
	}

}
