package glbm

import (
	"gitee.com/liumou_site/logger"
	"testing"
)

func TestWget(t *testing.T) {
	if WgetToSave("http://down.liumou.site/upload/Summary.py", "s.py") {
		logger.Info("下载成功")
	} else {
		logger.Error("下载失败")
	}
}

func TestCurl(t *testing.T) {
	if CurlToSave("http://down.liumou.site/upload/Summary.py", "s.py", true) {
		logger.Info("下载成功")
	} else {
		logger.Error("下载失败")
	}
}
