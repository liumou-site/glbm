package glbm

import (
	"fmt"
	"strings"

	"gitee.com/liumou_site/gbm"
)

// InstallListLocal 使用 Dpkg安装本地文件(列表传入)
func (api *ApiDpkg) InstallListLocal(fileList []string, name string) {
	installStr := gbm.SliceToString(fileList, " ")
	api.InstallFile(installStr, name)
}

// InstallFile 使用 Dpkg安装本地文件
func (api *ApiDpkg) InstallFile(pac string, name string) {
	api.Result = false // 初始化状态
	api.Sudo.RunSudo("dpkg -i", pac)
	api.Err = api.Sudo.Err
	if api.Err == nil {
		api.Result = true
		if api.Info {
			logs.Info("[ %s ] Installation succeeded", name)
		}
	} else {
		logs.Error("[ %s ] Installation failed", name)
	}
}

// UninstallSlice 使用 Dpkg卸载包(切片传入)
func (api *ApiDpkg) UninstallSlice(pacList []string) {
	uninstallStr := gbm.SliceToString(pacList, " ")
	api.Uninstall(uninstallStr)
}

// Uninstall 使用 Dpkg卸载单个包
func (api *ApiDpkg) Uninstall(Package string) {
	api.Result = false // 初始化状态
	api.Sudo.RunSudo("dpkg", "-P", Package)
	if api.Sudo.Err == nil {
		api.Result = true
		if api.Info {
			logs.Info("[ %s ] Uninstall succeeded", Package)
		}
	} else {
		logs.Error("[ %s ] Uninstall Failed", Package)
	}
	api.Err = api.Sudo.Err
}

// GetPackageStatus 使用 Dpkg查询包状态, 通过res返回字典,通过status返回查询状态,字典key(status/Name/version)
func (api *ApiDpkg) GetPackageStatus(pacPackage string) (m map[string]string) {
	cmd := fmt.Sprintf("dpkg -l | grep %s |  awk '{print $1,$2,$3}'", pacPackage)
	api.Sudo.RunScript(cmd)
	m = map[string]string{
		"status":  "Query failed",
		"Name":    "Query failed",
		"version": "Query failed",
	}
	api.Err = api.Sudo.Err
	if api.Err == nil {
		strSp := strings.Split(api.Sudo.Strings, " ")
		strSp = gbm.SliceRemoveNull(strSp)
		_name := fmt.Sprintf(strSp[1])
		_status := fmt.Sprintf(strSp[0])
		_ver := fmt.Sprintf(strSp[2])
		// 定义一个字典用于存储数据
		m = map[string]string{
			"status":  _status,
			"Name":    _name,
			"version": _ver,
		}
		if api.Info {
			logs.Info("query was successful: %s", pacPackage)
		}
	}
	return
}

// Installed 使用Dpkg查询是否已安装(ii)
func (api *ApiDpkg) Installed(pac string) bool {
	m := api.GetPackageStatus(pac)
	if api.Err == nil {
		status := m["status"]
		if status == "ii" {
			return true
		}
	}
	return false
}

// CheckVersion 使用dpkg检查本地安装版本与标准版本是否一致
func (api *ApiDpkg) CheckVersion(pac string, version string) (status_ bool, ver_ string) {
	info := api.GetPackageStatus(pac)
	ver := info["version"]
	name := info["Name"]
	status := info["status"]
	if api.Err == nil {
		if name != pac {
			mess := fmt.Sprintf("Name mismatch: %s != %s", name, pac)
			logs.Warn(mess)
			return false, name
		}
		if status == "ii" {
			if api.Info {
				logs.Info("Normal status")
			}
			if ver == version {
				if api.Info {
					logs.Info("Version consistency")
				}
				return true, ver
			} else {
				logs.Warn("Inconsistent version: %s != %s", ver, version)
				return false, ver
			}
		} else {
			logs.Warn("Abnormal status")
			return false, ver
		}
	}
	return false, ver
}

// CheckPacKey 使用两个关键词查询本地是否已安装某个软件包并返回最终包名
func (api *ApiDpkg) CheckPacKey(pac1, pac2 string) (bool, string) {
	var result bool
	cmd := fmt.Sprintf("dpkg -l | grep %s | grep %s | awk '{print $2}' | sed -n 1p ", pac1, pac2)
	api.Sudo.RunScript(cmd) // 判断是否存在
	ou := strings.Split(api.Sudo.Strings, "\n")
	ou = gbm.SliceRemoveNull(ou)
	api.Sudo.Strings = ou[0]
	if len(api.Sudo.Strings) >= 2 {
		result = true
	} else {
		result = false
	}
	if api.Sudo.Err != nil {
		result = false
	}
	if !result {
		if api.Debug {
			logs.Debug(api.Sudo.Strings)
		}
	}
	if api.Sudo.Strings == "0" {
		result = false
	}
	return result, api.Sudo.Strings
}

// ConfigureAll 使用 Dpkg --configure -a继续配置
func (api *ApiDpkg) ConfigureAll() bool {
	api.Sudo.RunShell("dpkg --configure -a")
	if api.Sudo.Err == nil && api.Info {
		logs.Info("Configure succeeded")
	} else {
		logs.Error("Configure Failed")
	}
	return api.Sudo.Err == nil
}
