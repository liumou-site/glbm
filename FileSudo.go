package glbm

// CopySudo 使用Sudo权限复制文件/文件夹
func (api *ApiFileSudo) CopySudo(dst string) *ApiFileSudo {
	api.Dst = dst
	api.sudo.RunSudo("cp", "-rf", api.Src, api.Dst)
	api.Err = api.sudo.Err
	return api
}

// MoveSudo 使用Sudo权限移动文件/文件夹
func (api *ApiFileSudo) MoveSudo(dst string) *ApiFileSudo {
	api.Dst = dst
	api.sudo.RunSudo("mv", api.Src, api.Dst)
	api.Err = api.sudo.Err
	return api
}

// DeleteSudo 使用Sudo权限删除文件/文件夹
func (api *ApiFileSudo) DeleteSudo(filename string) *ApiFileSudo {
	api.sudo.RunSudo("rm", "-rf", filename)
	api.Err = api.sudo.Err
	return api
}

// DeleteFileSudo 使用Sudo权限删除文件
func (api *ApiFileSudo) DeleteFileSudo(filename string) *ApiFileSudo {
	api.sudo.RunSudo("rm", "-f", filename)
	api.Err = api.sudo.Err
	return api
}
