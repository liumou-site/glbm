package glbm

import (
	"fmt"
	"strings"

	"gitee.com/liumou_site/gf"
)

func (api *ApiService) Start() {
	api.sudo.RunSudo("systemctl start ", api.Name)
	api.Err = api.sudo.Err
}

func (api *ApiService) ReStart() {
	api.sudo.RunSudo("systemctl restart", api.Name)
	api.Err = api.sudo.Err
}

func (api *ApiService) Stop() {
	api.sudo.RunSudo("systemctl stop", api.Name)
	api.Err = api.sudo.Err
}

func (api *ApiService) ReLoad() *ApiService {
	api.sudo.RunSudo("systemctl reload", api.Name)
	api.Err = api.sudo.Err
	return api
}

func (api *ApiService) Enable() *ApiService {
	api.sudo.RunSudo("systemctl enable", api.Name)
	api.Err = api.sudo.Err
	return api
}

func (api *ApiService) Disable() *ApiService {
	api.sudo.RunSudo("systemctl disable", api.Name)
	api.Err = api.sudo.Err
	return api
}

func (api *ApiService) ReLoadDaemon() *ApiService {
	api.sudo.RunSudo("systemctl reload-daemon")
	api.Err = api.sudo.Err
	return api
}

func (api *ApiService) Exists() bool {
	api.sudo.RunScriptSudo("systemctl -all| awk '{print $1}'")
	r := gf.NewReadFile("demo")
	r.Text = api.sudo.Strings
	services := strings.Split(r.Text, "\n")
	fmt.Println(services)
	for _, s := range services {
		if strings.Contains(s, api.Name) {
			return true
		}
	}
	return false
}
