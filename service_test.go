package glbm

import (
	"testing"
)

// 服务管理测试
func TestServiceApi_ReLoad(t *testing.T) {
	s := NewService("docker.service", "1")
	s.ReLoad().ReStart()       // 重新加载配置并重启服务
	s.Stop()                   // 停止服务
	s.Start()                  // 启动服务
	s.ReLoadDaemon().ReStart() // 重载全部守护进程配置并重启服务
	if s.Exists() {
		t.Error("service exits")
	}
	s.Start()

}

// 服务管理测试
func TestServiceApi_Exists(t *testing.T) {
	s := NewService("docker.service", "1")
	if !s.Exists() {
		logs.Error("service not exists")
	}
}
