package glbm

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	dpkg := NewConnect(true)
	logs.Info("开始新建")
	err := dpkg.AddConnect()
	if err != nil {
		return
	}
}

func TestConnectDns(t *testing.T) {
	dpkg := NewConnect(true)
	logs.Info("开始添加")
	err := dpkg.AddDns()
	if err != nil {
		logs.Error("配置错误")
	}
}

func TestApiConnection_GetUseCon(t *testing.T) {
	co := NewConnect(true)
	co.GetUseCon()
	if co.Err != nil {
		logs.Error("配置错误")
	}
}

func TestApiConnection_GetConList(t *testing.T) {
	co := NewConnect(true)
	co.GetConList()
	if co.Err != nil {
		logs.Error("配置错误")
	} else {
		logs.Info("连接列表如下")
		fmt.Println(co.ConList)
	}
}
