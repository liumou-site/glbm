package glbm

import (
	"fmt"
	"gitee.com/liumou_site/gf"

	"gitee.com/liumou_site/gcs"
	"github.com/spf13/cast"
)

// WgetToSave 使用wget下载文件到指定路径
func WgetToSave(url, filename string) bool {
	if gcs.CheckCmd("wget") {
		run := gcs.NewShell()
		run.RunShell("wget -c -O ", filename, " ", url)
		if run.Err == nil {
			logs.Info("Download Succeeded", filename)
			return true
		} else {
			logs.Error("Download Failed", filename)
		}
	} else {
		logs.Error("Wget Command does not exist")
	}
	return false
}

// CurlToSave 使用Curl下载文件到指定路径(cover是否覆盖已有文件)
func CurlToSave(url, filename string, cover bool) bool {
	if gcs.CheckCmd("curl") {
		run := gcs.NewShell()
		f := gf.NewFile(filename)
		f.IsFile()
		if f.IsFiles {
			if cover {
				f.DeleteFile() // 删除文件
			}
		}
		cmd := cast.ToString(fmt.Sprintf("curl -o %s %s", filename, url))
		logs.Debug(cmd)
		run.RunShell(cmd)
		if run.Err == nil {
			logs.Info("Download Succeeded", filename)
			return true
		} else {
			logs.Error("Download Failed", filename)
		}
	} else {
		logs.Error("Curl Command does not exist")
	}
	return false
}
