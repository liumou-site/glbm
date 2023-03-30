package glbm

import (
	"log"
	"os/user"
)

// GetUserInfo 获取当前系统环境信息(display: 是否显示信息)
func GetUserInfo() (*user.User, error) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return currentUser, nil
}
