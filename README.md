# glbm

## 介绍

`glbm`的全称是(`Go Linux Basic module`), 即使用Golang编写的Linux基础模块

[Gitee 项目地址](https://gitee.com/liumou_site/glbm)

## 特色

* 使用全中文注释，即使小白也能轻松上手
* 完全开源、永久免费

## 功能清单

详细功能清单请访问[pkg.go.Dev](https://pkg.go.Dev/gitee.com/liumou_site/glbm)

## 安装

### 关闭校验

设置环境变量

```shell
export GOSUMDB="off"
```

### 开始安装

首先进入你的项目根目录，然后执行下面的命令

```shell
go get -u gitee.com/liumou_site/glbm
```

效果如下

```shell
➜  go get -u gitee.com/liumou_site/glbm
go: downloading gitee.com/liumou_site/glbm v1.3.4
go: gitee.com/liumou_site/glbm upgrade => v1.3.4
go: downloading gitee.com/liumou_site/gcs v1.2.3
go: downloading gitee.com/liumou_site/logger v1.1.1
go: downloading github.com/spf13/cast v1.5.0
➜  demo
```

## 使用


### Demo

> 最终请看单元测试案例

* [Dpkg](dpkg_test.go)
* [Apt](apt_test.go)
* [glbm](glbm_test.go)

### 其他演示

```golang
package main

import (
	"fmt"
	"gitee.com/liumou_site/glbm"
)

// Ju 权限验证
func Ju() {
	g := glbm.CheckSudo("1")
	if g {
		fmt.Println("密码检验正确")
	} else {
		fmt.Println("密码错误或无权限")
	}
	d := glbm.Developer()
	if d {
		fmt.Println("已开启开发者")
	} else {
		fmt.Println("未开启开发者")
	}
}

// GetUser 获取用户信息
func GetUser() {
	get, username, uid, uHome := glbm.GetUserInfo(false)
	if get {
		fmt.Println(username)
		fmt.Println(uid)
		fmt.Println(uHome)
	}
}

// OsInfo 获取系统信息
func OsInfo() {
	osType, osArch, ov, err := glbm.GetOsInfo()
	if err != nil{
		fmt.Println("获取失败")
    }else{
		fmt.Println(osType)
		fmt.Println(osArch)
		fmt.Println(ov)
		GetUser()
    }
}
```

# 问题反馈

点击链接加入QQ群聊【[坐公交也用券](https://jq.qq.com/?_wv=1027&k=FEeLQ6tz)】