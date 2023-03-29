package glbm

import (
	"gitee.com/liumou_site/logger"
	"testing"
)

func TestApiFile_Copy(t *testing.T) {
	fm := NewFile("demo/s.txt")
	fm.Copy("demo/dst.txt")
	if fm.Err == nil {
		logger.Info("复制成功")
	} else {
		logger.Error("复制失败")
	}
}

func TestApiFile_Delete(t *testing.T) {
	fm := NewFile("demo/dir")
	fm.Delete()
	if fm.Err == nil {
		logger.Info("文件夹删除成功")
	} else {
		logger.Error("文件夹删除失败")
	}
}

func TestApiFile_DeleteFile(t *testing.T) {
	fm := NewFile("demo/df.txt")
	fm.Delete()
	if fm.Err == nil {
		logger.Info("文件夹删除成功")
	} else {
		logger.Error("文件夹删除失败")
	}
}
