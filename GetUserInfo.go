package glbm

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
)

// GetUserInfo 获取当前系统环境信息(display: 是否显示信息)
func GetUserInfo(display bool) (ok bool, username string, userid int, UserHome string) {
	ok = false
	username = "None"
	userid = 10000
	UserHome = "None"
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	username = currentUser.Name
	userid, err = strconv.Atoi(currentUser.Uid)
	if err != nil {
		userid = 10000
	}
	UserHome = currentUser.HomeDir
	if display {
		fmt.Println("UserName is: ", username)
		fmt.Println("UserId is: ", userid)
		fmt.Println("UserHome : ", UserHome)
	}
	return
}
