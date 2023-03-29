package glbm

import (
	"fmt"
	"testing"
)

// Apt 测试
func TestApt(t *testing.T) {
	apt := NewApt("1", true, false)
	i := apt.Install("vsftpd")
	if i {
		fmt.Println("安装成功")
	} else {
		fmt.Println("Failed to install")
	}

	// fileMan := "go.deb vsftpd.deb"
	// fileList := strings.Split(fileMan, " ")
	// s := apt.LocalInstallList(fileList, "测试")
	// if s {
	// 	logger.Info("ok")
	// }
	// f := apt.AptLocalInstallFile("/home/liumou/LinuxData/git/golang/modular/glbm/ApiDpkg.go", "测试")
	// if f == nil {
	// 	logger.Info("ok")
	// }
}

// 卸载
func TestAptUninstall(t *testing.T) {
	apt := NewApt("1", true, true)
	r := apt.Uninstall("vsftpd", "ftp")
	if r {
		fmt.Println("Uninstall is succeeded")
	} else {
		fmt.Println("Uninstall is Failed")
	}
	r2 := apt.Uninstall("vsft", "ftp")
	if r2 {
		fmt.Println("Uninstall is succeeded")
	} else {
		fmt.Println("Uninstall is Failed")
	}
}
