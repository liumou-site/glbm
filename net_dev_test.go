package glbm

import (
	"fmt"
	"testing"

	"gitee.com/liumou_site/logger"
)

// 获取网卡信息
func TestDevInfo(t *testing.T) {
	n := NewNmcli()
	info, err := n.GetAllEthInfo()
	if err == nil {
		logger.Info("获取成功")
		for _, i := range info {
			fmt.Println(i)
		}
	} else {
		logger.Error("获取失败")
	}
}

// 获取默认网关
func TestDefaultGw(t *testing.T) {
	n := NewNmcli()
	res := n.GetDefaultRouteInfo()
	if res == nil {
		logger.Info("获取成功")
	} else {
		logger.Error("获取失败")
	}
	fmt.Println("网卡名称: ", n.DefaultDevices)
	fmt.Println(n.DefaultGw)
}

// 获取指定网卡的网关
func TestApiNmcli_GetEthGw(t *testing.T) {
	n := NewNmcli()
	res, err := n.GetEthGw(n.DefaultDevices)
	if err == nil {
		logger.Info("获取成功")
		fmt.Println(res)
	} else {
		logger.Error("获取失败")
		fmt.Println(err.Error())
	}
}

// 获取指定网卡的信息
func TestApiNmcli_GetEthInfo(t *testing.T) {
	n := NewNmcli()
	res, err := n.GetEthInfo("enp125s0f1")
	if err == nil {
		logger.Info("获取成功")
		fmt.Println(res)
	} else {
		logger.Error("获取失败")
		fmt.Println(err.Error())
	}
	res, err = n.GetEthInfo("wlan0")
	if err == nil {
		logger.Info("获取成功")
		fmt.Println(res)
	} else {
		logger.Error("获取失败")
		fmt.Println(err.Error())
	}
}
