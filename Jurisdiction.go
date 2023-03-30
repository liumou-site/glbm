package glbm

import (
	"container/list"
	"fmt"
	"strings"

	"gitee.com/liumou_site/gf"

	"gitee.com/liumou_site/gcs"
)

// Developer UOS系统下检查是否开启开发者模式
func Developer() bool {
	info, err := GetOsInfo()
	if err != nil {
		logs.Error("无法获取当前系统类型")
		return false
	}
	if info.Name != "uos" {
		logs.Info("当前系统非UOS系统,无需检测开发者模式")
		return true
	}
	gc := gcs.NewShell()
	gc.Debug = true
	fileList := list.New()
	fileList.PushBack("/var/lib/deepin/developer-install_modes/enabled")
	fileList.PushBack("/var/lib/deepin/developer-install_mode/enabled")
	fileList.PushBack("/var/lib/deepin/developer-mode/enabled")
	for i := fileList.Front(); i != nil; i = i.Next() {
		file := fmt.Sprintln(i.Value)
		file = strings.Fields(file)[0]
		gc.RunShell("cat", file)
		if gc.Err == nil {
			logs.Debug("文件内容: %s", gc.Strings)
			return true
		} else {
			logs.Error(gc.Err.Error())
		}
	}
	return false
}

// CheckSudo 判断是否拥有sudo权限
func CheckSudo(password string) bool {
	c := gcs.NewSudo(password)
	f := gf.NewFile("/cmd.testing")
	f.Exists()
	if f.ExIst {
		c.RunSudo("rm -f /cmd.testing")
	} else {
		cmd := "touch /cmd.testing"
		c.RunSudo(cmd)
		if c.Err == nil {
			c.RunSudo("rm -f /cmd.testing")
		}
	}
	return c.Err == nil
}
