package glbm

import (
	"gitee.com/liumou_site/logger"
	"testing"
)

func TestApiFileSudo(t *testing.T) {
	fm := NewFileSudo("demo/s.txt", "1")
	fm.CopySudo("demo/dst.txt")
	if fm.Err == nil {
		logger.Info("复制成功")
	} else {
		logger.Error("复制失败")
	}
	fm.DeleteSudo("demo/s.txt")
}
