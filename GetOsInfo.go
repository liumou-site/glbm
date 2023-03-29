package glbm

import (
	"fmt"
	"runtime"
	"strings"

	"gitee.com/liumou_site/gcs"
)

// GetOsInfo 获取系统信息
func GetOsInfo() (osType, osArch, osVersion string, err error) {
	gc := gcs.NewShell()
	osType = runtime.GOOS
	osArch = runtime.GOARCH
	osVersion = "10"
	if osType == "linux" {
		gc.RunShell("grep ^ID /etc/os-release")
		text := gc.Strings
		osType = strings.Split(text, "=")[1]
		osType = strings.Fields(osType)[0] // 去除：换行、空等
		if gc.Err != nil {
			return "", "", "0", fmt.Errorf("查询系统发行版失败")
		}
		logs.Info("系统类型: ", osType)
		switch osType {
		case "kylin":
			{
				gc.RunShell("sed -n 2p /etc/kylin-build")
				osVersion = strings.Split(gc.Strings, " ")[1]
			}
		case "kali":
			{
				gc.RunShell("sed -n 3p /etc/os-release")
				osVersion = strings.Split(gc.Strings, "=")[1]
				osVersion = strings.Replace(osVersion, "\"", "", 2)
			}
		case "uos":
			{
				gc.RunShell("grep ^Min /etc/os-version")
				osVersion = strings.Split(gc.Strings, "=")[1]
			}
		default:
			{
				gc.RunShell("sed -n 3p /etc/os-release")
				osVersion = strings.Split(gc.Strings, "=")[1]
				osVersion = strings.Replace(osVersion, "\"", "", 2)
			}
		}
		err = gc.Err
		osVersion = strings.Fields(osVersion)[0]
		return
	} else {
		err = fmt.Errorf("暂不支持WIndows系统")
		logs.Warn(err.Error())
	}
	return
}
