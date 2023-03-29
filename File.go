package glbm

import "gitee.com/liumou_site/gf"

// Copy 复制文件/文件夹
func (api *ApiFile) Copy(dst string) *ApiFile {
	api.Dst = dst
	api.fileMan.Src = api.Src
	api.fileMan.Dst = api.Dst
	api.fileMan.Copy()
	api.Err = api.fileMan.Err
	return api
}

// Move 移动文件/文件夹
func (api *ApiFile) Move(dst string) *ApiFile {
	api.Dst = dst
	api.fileMan.Src = api.Src
	api.fileMan.Dst = api.Dst
	api.fileMan.Move()
	api.Err = api.fileMan.Err
	return api
}

// Delete 删除文件/文件夹,默认使用Src文件
func (api *ApiFile) Delete() *ApiFile {
	fs := gf.NewFile(api.Src)
	fs.Delete()
	fs.Close()
	api.Err = fs.Err
	return api
}

// DeleteFile 删除文件,默认使用Src文件
func (api *ApiFile) DeleteFile() *ApiFile {
	fs := gf.NewFile(api.Src)
	fs.DeleteFile()
	fs.Close()
	api.Err = fs.Err
	return api
}
