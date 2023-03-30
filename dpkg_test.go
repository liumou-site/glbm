package glbm

import (
	"fmt"
	"testing"

	"gitee.com/liumou_site/logger"
)

func TestDpkg(t *testing.T) {
	dpkg := NewDpkg("1", false)

	if dpkg.Installed("docker-ce") {
		logger.Info("docker.io IS Installed")
	} else {
		logger.Error("docker.io not installed")
	}
	if dpkg.Installed("openssh-server") {
		logger.Info("SSh IS Installed")
	} else {
		logger.Error("SSh not installed")
	}
}

func TestCheckIed(t *testing.T) {
	dpkg := NewDpkg("1", false)
	res, data := dpkg.CheckPacKey("wps", "code")
	if res {
		logger.Info("查询成功: ", data)
	} else {
		logger.Emer("查询失败")
	}
	res1, data1 := dpkg.CheckPacKey("dmi", "code")
	if res1 {
		logger.Info("查询成功: ", data1)
	} else {
		logger.Emer("查询失败")
	}
}

func TestUninstall(t *testing.T) {
	dpkg := NewDpkg("Liumou17?", true)
	dpkg.Uninstall("vsftpd")
	if dpkg.Result {
		logger.Info("卸载 成功")
	} else {
		logger.Emer("卸载 失败")
	}
	dpkg.Uninstall("ftp")
	if dpkg.Result {
		logger.Info("卸载 成功")
	} else {
		logger.Emer("卸载 失败")
	}
}

func TestApiDpkg_GetPackageStatus(t *testing.T) {
	dpkg := NewDpkg("1", false)
	info := dpkg.GetPackageStatus("docker.io")
	fmt.Println(info)
}

func TestApiDpkg_ConfigureAll(t *testing.T) {
	dpkg := NewDpkg("Liumou17?", false)
	dpkg.ConfigureAll()
	if dpkg.Err == nil {
		logger.Info("配置成功")
	} else {
		logger.Error("配置失败")
	}
}
