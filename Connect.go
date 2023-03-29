package glbm

import (
	"fmt"
	"net"
	"strings"

	"gitee.com/liumou_site/gcs"
	"gitee.com/liumou_site/logger"
)

// AddConnect 新增连接
func (c *ApiConnection) AddConnect() error {
	c.Address = net.ParseIP("127.0.0.1")
	c.Dns = strings.Split("114.114.114.114 8.8.8.8", " ")
	c.Gw = net.ParseIP("10.1.1.1")
	c.Mask = 24
	c.uuid = "as216a5w4d1646"
	c.Method = "auto"
	c.Types = "ethernet"
	c.Name = "test"
	c.Dev = "eth0"
	cmd := fmt.Sprintf("nmcli connection add type %s  con-Name %s ifname %s", c.Types, c.Name, c.Dev)
	fmt.Println(cmd)
	command := gcs.NewShell()
	command.RunShell(cmd)
	return command.Err
}

// AddDns 增加DNS,默认(119.29.29.29 180.76.76.76)
func (c *ApiConnection) AddDns() (err error) {
	shell := gcs.NewShell()

	fmt.Println("增加DNS")
	if len(c.Dns) >= 1 {
		for index, server := range c.Dns {
			if index == 1 {
				cmd := fmt.Sprintf("nmcli connection modify %s ipv4.Dns %s", c.uuid, server)
				shell.RunShell(cmd)
				return shell.Err
			} else {
				cmd := fmt.Sprintf("nmcli connection modify %s +ipv4.Dns %s", c.uuid, server)
				shell.RunShell(cmd)
				return shell.Err
			}
		}
	} else {
		logger.Error("当前实例未配置DNS服务器信息,请先把DNS信息赋值到结构体中")
		err = fmt.Errorf("dns服务器列表为空")
		return err
	}
	return nil
}

// GetUseCon 获取正在使用的连接配置
func (c *ApiConnection) GetUseCon() {
	cmd := "nmcli connection show --active"
	c.cmd.RunShell(cmd)
	if c.cmd.Err != nil {
		c.Err = c.cmd.Err
		logger.Error("无法获取当前连接配置,执行的命令: ", cmd)
		return
	}
	sp := strings.Split(c.cmd.Strings, "\n")
	if len(sp) == 0 {
		logger.Error("命令执行成功,但无法获取内容,请检查nmcli服务")
		c.Err = fmt.Errorf("命令执行成功,但无法获取内容,请检查nmcli服务")
		return
	}
	fmt.Println(sp)
}

// GetConList 获取所有连接配置
func (c *ApiConnection) GetConList() {
	cmd := "nmcli connection show"
	c.cmd.RunShell(cmd)
	if c.cmd.Err != nil {
		c.Err = c.cmd.Err
		logger.Error("无法获取当前连接配置,执行的命令: ", cmd)
		return
	}
	sp := strings.Split(c.cmd.Strings, "\n")
	if len(sp) == 0 {
		logger.Error("命令执行成功,但无法获取内容,请检查nmcli服务")
		c.Err = fmt.Errorf("命令执行成功,但无法获取内容,请检查nmcli服务")
		return
	}
	fmt.Println(sp)
	c.cmd.Column(1, "")
	if c.Err != nil {
		logs.Error("数据截取失败")
		return
	}
	cl := strings.Split(c.cmd.Strings, "\n")
	if len(cl) == 0 {
		c.Err = fmt.Errorf("无法获取连接列表")
		return
	}
	for _, v := range cl {
		if v != "NAME" {
			c.ConList = append(c.ConList, v)
		}
	}

}

// GetConDnsList 获取一个连接配置的DNS服务器列表
func (c *ApiConnection) GetConDnsList(con string) {
	cmd := "nmcli connection show" + con
	c.cmd.RunShell(cmd)
	if c.cmd.Err != nil {
		c.Err = c.cmd.Err
		logger.Error("无法获取当前连接配置,执行的命令: ", cmd)
		return
	}
}
