package glbm

import (
	"strings"

	"gitee.com/liumou_site/gcs"
	"gitee.com/liumou_site/gf"
)

// NewNmcli 返回一个网卡管理实例
func NewNmcli() *ApiNmcli {
	net := new(ApiNmcli)
	return net
}

// NewDpkg 初始化(realtime: 是否开启实时刷新)
func NewDpkg(password string, realtime bool) *ApiDpkg {
	d := new(ApiDpkg)
	d.password = password
	d.Sudo = gcs.NewSudo(password)
	d.Sudo.Realtime = realtime
	d.Info = false
	d.Debug = false
	return d
}

// NewApt 管理初始化
func NewApt(password string, debug, realtime bool) *ApiApt {
	c := new(ApiApt)
	c.password = password
	c.Sudo = gcs.NewSudo(password)
	c.Sudo.Realtime = realtime
	c.Debug = debug
	c.Info = false
	c.dpkg = NewDpkg(password, false)
	return c
}

// NewConnect 连接管理
func NewConnect(debug bool) *ApiConnection {
	c := new(ApiConnection)
	net := NewNmcli()
	err := net.GetDefaultRouteInfo()
	if err != nil {
		return nil
	}
	eth, err := net.GetAllEthInfo()
	if err != nil {
		logs.Error("本机网络信息获取失败")
		return nil
	}
	for _, i := range eth {
		logs.Debug(i.Name)
	}
	c.Dns = strings.Split("119.29.29.29 180.76.76.76", " ")
	c.Address = net.address    // 设置当前默认IP
	c.Mask = net.mask          // 设置当前默认掩码
	c.Dev = net.DefaultDevices // 设置当前默认网卡设备
	c.Gw = net.DefaultGw       // 设置当前默认网关
	if debug {
		logs.Debug("(c.Address)设置当前默认IP: %s", c.Address)
		logs.Debug("(c.Mask)设置当前默认掩码: %d", c.Mask)
		logs.Debug("(c.Dev)设置当前默认网卡: %s", c.Dev)
		logs.Debug("(c.Gw)设置当前默认网关: %s", c.Gw)
		logs.Debug("(c.Dns)设置当前默认DNS: %s", c.Dns)
		logs.Debug("如需修改以上值请自行通过实例变量赋值")
	}
	return c
}

func NewService(name, password string) *ApiService {
	s := new(ApiService)
	sudo := gcs.NewSudo(password)
	s.Name = name
	s.sudo = sudo
	s.sudo.Realtime = false
	s.Password = password
	return s
}

// NewFile 文件管理模块
func NewFile(src string) *ApiFile {
	s := new(ApiFile)
	s.Src = src
	s.shell = gcs.NewShell()
	s.fileMan = gf.NewFileMan(src, src)
	return s
}

// NewFileSudo 文件管理模块
func NewFileSudo(src, password string) *ApiFileSudo {
	s := new(ApiFileSudo)
	sudo := gcs.NewSudo(password)
	s.sudo = sudo
	s.Src = src
	s.sudo.Realtime = false
	s.Password = password
	s.shell = gcs.NewShell()
	s.fileMan = gf.NewFileMan(src, src)
	return s
}
