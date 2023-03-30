package glbm

import (
	"fmt"
	"runtime"
	"strings"

	"gitee.com/liumou_site/gcs"
)

// GetOsInfo 获取系统信息
func GetOsInfo() (info *OsInfo, err error) {
	info = new(OsInfo)
	gc := gcs.NewShell()
	osType := runtime.GOOS
	info.Arch = runtime.GOARCH
	info.Version = "10"
	if osType == "linux" {
		gc.RunShell("grep ^ID /etc/os-release")
		text := gc.Strings
		osType = strings.Split(text, "=")[1]
		osType = strings.Fields(osType)[0] // 去除：换行、空等
		info.Name = osType
		if gc.Err != nil {
			return info, fmt.Errorf("查询系统发行版失败")
		}
		switch osType {
		case "kylin":
			{
				gc.RunShell("sed -n 2p /etc/kylin-build")
				info.Version = strings.Split(gc.Strings, " ")[1]
				//gc.RunShell("grep CODENAME /etc/os-release")
				//info.CodeName = strings.Replace(gc.Strings, "\"", "", 2)
			}
		case "kali":
			{
				gc.RunShell("sed -n 3p /etc/os-release")
				info.Version = strings.Split(gc.Strings, "=")[1]
				info.Version = strings.Replace(info.Version, "\"", "", 2)
				//gc.RunShell("grep CODENAME /etc/os-release")
				//info.CodeName = strings.Replace(gc.Strings, "\"", "", 2)
			}
		case "uos":
			{
				gc.RunShell("grep ^Min /etc/os-version")
				info.Version = strings.Split(gc.Strings, "=")[1]
				//gc.RunShell("grep CODENAME /etc/os-release")
				//info.CodeName = strings.Replace(gc.Strings, "\"", "", 2)
			}
		default:
			{
				gc.RunShell("sed -n 3p /etc/os-release")
				info.Version = strings.Split(gc.Strings, "=")[1]
				info.Version = strings.Replace(info.Version, "\"", "", 2)
			}
		}
		gc.RunShell("grep CODENAME /etc/os-release")
		info.CodeName = strings.Replace(gc.Strings, "\"", "", 2)
		info.CodeName = strings.Split(info.CodeName, "=")[1]
		info.CodeName = strings.Fields(info.CodeName)[0]
		err = gc.Err
		info.Version = strings.Fields(info.Version)[0]
		return info, nil
	} else {
		err = fmt.Errorf("暂不支持WIndows系统")
		logs.Warn(err.Error())
		return info, err
	}
}
