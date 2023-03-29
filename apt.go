package glbm

import (
	"container/list"
	"fmt"

	"gitee.com/liumou_site/gbm"
	"gitee.com/liumou_site/gf"
	"github.com/spf13/cast"
)

// LocalInstallList 使用apt安装本地文件(列表传入)
func (api *ApiApt) LocalInstallList(fileList []string, name string) bool {
	// 将切片转换为字符串
	installStr := gbm.SliceToString(fileList, " ")
	status := api.AptLocalInstallStr(installStr, name)
	return status
}

// AptLocalInstallStr 使用apt安装本地文件(字符串),直接安装，不会检测文件是否存在
func (api *ApiApt) AptLocalInstallStr(installStr, name string) bool {
	api.Sudo.RunSudo("apt install -y -f", installStr)
	if api.Sudo.Err == nil && api.Info {
		logs.Info("[ %s ] Installation succeeded", name)
	} else {
		logs.Error("[ %s ] Installation Failed: ", name, api.Sudo.Err.Error())
	}
	api.Err = api.Sudo.Err
	return api.Sudo.Err == nil
}

// AptLocalInstallFile 使用apt安装本地文件(单个),安装之前将会检测文件是否存在
func (api *ApiApt) AptLocalInstallFile(filename, name string) (err error) {
	f := gf.NewFile(filename)
	f.IsFile() // 判断文件是否存在
	if f.IsFiles {
		api.Sudo.RunSudo("apt install -y -f ", filename)
		if api.Sudo.Err == nil && api.Info {
			logs.Info("[ %s ] Installation succeeded", name)
		} else {
			logs.Error("[ %s ] Installation Failed: ", name, api.Sudo.Err.Error())
		}
		api.Err = api.Sudo.Err
	} else {
		msg := fmt.Sprintf("fileMan does not exist: %s", filename)
		logs.Error(msg)
		api.Err = fmt.Errorf(msg)
		fmt.Println(err)
	}
	return
}

// UninstallList 通过apt卸载包(列表传入)
func (api *ApiApt) UninstallList(pacList list.List) (ok, failed list.List) {
	// 卸载成功的列表
	uninstallListOk := list.New()
	// 卸载失败的列表
	uninstallListFailed := list.New()
	for i := pacList.Front(); i != nil; i = i.Next() {
		pac := "" + fmt.Sprint(i.Value)
		un := api.Uninstall(pac, pac)
		if un {
			uninstallListOk.PushBack(pac)
		} else {
			uninstallListFailed.PushBack(pac)
		}
	}
	return *uninstallListOk, *uninstallListFailed
}

// UninstallSlice 通过apt卸载包(列表传入)
func (api *ApiApt) UninstallSlice(pacList []string) (ok, failed list.List) {
	// 卸载成功的列表
	uninstallListOk := list.New()
	// 卸载失败的列表
	uninstallListFailed := list.New()
	for i := range pacList {
		pac := pacList[i]
		un := api.Uninstall(pac, pac)
		if un {
			uninstallListOk.PushBack(pac)
		} else {
			uninstallListFailed.PushBack(pac)
		}
	}
	return *uninstallListOk, *uninstallListFailed
}

// Uninstall 通过apt卸载单个包
func (api *ApiApt) Uninstall(Package string, name string) (res bool) {
	// 判断是否已安装
	ied := api.dpkg.Installed(Package)
	if ied {
		api.Sudo.RunSudo("apt purge -y", Package)
		if api.Sudo.Err == nil && api.Info {
			logs.Info("[ %s ] Uninstallation succeeded", name)
		} else {
			logs.Error("[ %s ] Uninstallation Failed: ", name, api.Sudo.Err.Error())
		}
		api.Err = api.Sudo.Err
		return
	} else {
		msg := fmt.Sprintf("[ %s ] Package is not installed", name)
		logs.Error(msg)
		res = true
		api.Err = fmt.Errorf(msg)
	}
	return
}

// InstallList 通过apt安装包(列表传入)
func (api *ApiApt) InstallList(pacList list.List) (ok, failed list.List) {
	// 安装成功的列表
	installListOk := list.New()
	// 安装失败的列表
	installListFailed := list.New()
	for i := pacList.Front(); i != nil; i = i.Next() {
		pac := "" + fmt.Sprint(i.Value)
		un := api.Install(pac)
		if un {
			installListOk.PushBack(pac)
		} else {
			installListFailed.PushBack(pac)
		}
	}
	return *installListOk, *installListFailed
}

// InstallSlice 通过apt安装包(切片传入) - 推荐
func (api *ApiApt) InstallSlice(pacList []string) (ok, failed list.List) {
	// 安装成功的列表
	installListOk := list.New()
	// 安装失败的列表
	installListFailed := list.New()
	for p := range pacList {
		pac := cast.ToString(p)
		un := api.Install(pac)
		if un {
			installListOk.PushBack(pac)
		} else {
			installListFailed.PushBack(pac)
		}
	}
	return *installListOk, *installListFailed
}

// Install 使用apt在线安装包
func (api *ApiApt) Install(Package string) bool {
	api.Sudo.RunSudo("apt install -y", Package)
	if api.Sudo.Err == nil {
		if api.Info {
			logs.Info("[ %s ] Installation succeeded", Package)
		}
	} else {
		fmt.Println(api.Sudo.Err)
		logs.Error("Installation Failed")
	}
	api.Err = api.Sudo.Err
	return api.Err == nil
}

// UpdateIndex 更新索引
func (api *ApiApt) UpdateIndex() (res bool) {
	if api.Debug {
		logs.Info("Apt Index Update")
	}
	cmd := "apt update"
	api.Sudo.RunSudo(cmd)
	if api.Sudo.Err == nil && api.Info {
		logs.Info("Apt Index Update Succeeded")
	} else {
		logs.Error("Apt Index Update Failed: ")
	}
	api.Err = api.Sudo.Err
	return api.Sudo.Err == nil
}

// Upgrade 更新系统
func (api *ApiApt) Upgrade() (res bool) {
	logs.Info("System Update")
	api.Sudo.RunSudo("apt upgrade")
	if api.Sudo.Err == nil && api.Info {
		logs.Info("System Upgrade Successfully")
	} else {
		logs.Error("System Upgrade  Failed")
	}
	api.Err = api.Sudo.Err
	return api.Sudo.Err == nil
}
