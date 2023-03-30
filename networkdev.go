package glbm

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"gitee.com/liumou_site/gbm"

	"gitee.com/liumou_site/logger"
	"github.com/spf13/cast"
)

// 这是管理Linux网卡的模块

// GetConnectionList 获取指定设备及类型的连接列表
func (n *ApiNmcli) GetConnectionList() bool {
	if len(n.devList) == 0 {
		return false
	}
	c := fmt.Sprintf("nmcli connection |  grep %s | grep %s| awk '{print $1,$2}'", n.conType, n.device)
	n.cmd.RunShell(c)
	cs := strings.Split(n.cmd.Strings, "\n")
	cs = gbm.SliceRemoveNull(cs)
	n.connectionList = cs
	return n.cmd.Result
}

// IsIPv4 判断是否属于IPv4
func IsIPv4(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ".")
}

// IsLoopbackV4 判断是否属于回环
func IsLoopbackV4(ipAddr string) bool {
	return ipAddr == "127.0.0.1"
}

// IsDockerDevice 判断是否Docker网卡
func IsDockerDevice(eth string) bool {
	return eth == "docker0"
}

// GetAllEthInfo 获取所有网卡信息(IPV4)
func (n *ApiNmcli) GetAllEthInfo() ([]ApiEth, error) {
	var EthInfo []ApiEth
	interfaceList, err := net.Interfaces() // 获取接口列表信息
	if err != nil {
		logger.Error("获取网卡列表失败: ", err)
		return EthInfo, err
	}
	var byName *net.Interface
	var addrList []net.Addr
	for _, address := range interfaceList {
		byName, err = net.InterfaceByName(address.Name)
		if err != nil {
			fmt.Println(err)
			return EthInfo, err
		}
		addrList, err = byName.Addrs()
		if err != nil {
			logs.Error(err.Error())
			return EthInfo, err
		}
		for _, Addr := range addrList {
			IpInfo := strings.SplitN(Addr.String(), "/", 2)
			mask := cast.ToInt(IpInfo[1]) // 子网掩码
			ip := IpInfo[0]               // IP地址
			//fmt.Println("\n网卡名称: ", address.Name)
			//fmt.Println("IP地址: ", ip)
			//fmt.Println("子网掩码: ", mask)
			n.devList = append(n.devList, address.Name) // 增加网卡切片
			n.mac = cast.ToString(address.HardwareAddr) // 记录网卡Mac地址
			//fmt.Println("网卡地址: ", address.HardwareAddr)
			if IsIPv4(ip) && !IsLoopbackV4(ip) && !IsDockerDevice(address.Name) { // 判断是否属于IPv4格式/回环地址/Docker网卡
				eth := ApiEth{
					Name:  address.Name,
					Index: address.Index,
					Ipv4:  net.ParseIP(ip),
					Mask:  mask,
					Mac:   n.mac}
				EthInfo = append(EthInfo, eth)
			} else {
				logger.Debug("略过非IPV4地址: ", ip)
			}
		}
	}
	n.deviceQuantity = len(EthInfo) // 设备数量
	return EthInfo, nil
}

// GetEthGw 获取指定网卡的网关地址
func (n *ApiNmcli) GetEthGw(eth string) (net.IP, error) {
	n.cmd.RunShell("nmcli -f IP4.GATEWAY dev show", eth)
	if n.cmd.Err != nil {
		return nil, n.cmd.Err
	}
	f := strings.Split(n.cmd.Strings, ":")
	if len(f) != 2 {
		return nil, fmt.Errorf("无法通过 : 截取网关信息: %s", n.cmd.Strings)
	}

	gw := net.ParseIP(strings.Fields(f[1])[0])
	if gw != nil {
		return gw, nil
	}
	return nil, fmt.Errorf("数据匹配失败")
}

// GetEthInfo 获取指定网卡的网络信息
func (n *ApiNmcli) GetEthInfo(eth string) (EthInfo, error) {
	var dnsList []net.IP
	exists := false
	info, err := n.GetAllEthInfo()
	if err != nil {
		return EthInfo{}, err
	}
	for i, v := range info {
		if v.Name == eth {
			exists = true // 证明匹配到了网卡
			n.cmd.RunShell("nmcli device show", eth)
			if n.cmd.Err != nil {
				logs.Warn("设备信息获取失败")
			}
			n.cmd.Grep("IP4.DNS").Column(0, " ")
			if n.cmd.Err == nil {
				ds := strings.Split(n.cmd.Strings, "\n")
				if len(ds) == 0 {
					logger.Warn("DNS第一次获取失败")
				} else {
					for _, v := range ds {
						ipCut := strings.Split(v, ":") // 通过冒号截取数据
						if len(ipCut) == 2 {
							// 如果截取的数据量等于2则进行截取
							ip := ipCut[1]
							IpAddr := net.ParseIP(ip)
							if IpAddr != nil {
								logs.Debug("添加DNS信息: ", IpAddr)
								dnsList = append(dnsList, IpAddr)
							}
						}
					}
				}
			}
			logger.Info("第: [ %d ]个网卡匹配成功", i)
			gw, ge := n.GetEthGw(eth)
			if dnsList == nil {
				n.cmd.RunShell("grep nameserver /etc/resolv.conf")
				if n.cmd.Err == nil {
					dns := strings.Split(n.cmd.Strings, "\n")
					if len(dns) >= 1 {
						for _, v := range dns {
							ips := strings.Split(v, " ")
							if len(ips) == 2 {
								ip := ips[1]
								IpAddr := net.ParseIP(ip)
								if IpAddr != nil {
									dnsList = append(dnsList, IpAddr)
								}
							}
						}
					}
				}
			}
			if ge == nil {
				if dnsList != nil {
					return EthInfo{IP: v.Ipv4, GW: gw, MASK: v.Mask, DNS: dnsList}, nil
				} else {
					return EthInfo{IP: v.Ipv4, GW: gw, MASK: v.Mask, DNS: nil}, nil
				}
			} else {
				return EthInfo{IP: v.Ipv4, GW: gw, MASK: v.Mask, DNS: nil}, ge
			}
		}
	}
	getErr := fmt.Errorf("网卡不存在,无法获取信息")
	if exists {
		getErr = fmt.Errorf("网卡存在，但是无法获取网关")
	}
	return EthInfo{}, getErr
}

// GetDefaultRouteInfo 获取默认网关&网卡设备
func (n *ApiNmcli) GetDefaultRouteInfo() error {
	cmd := exec.Command("ip", "-4", "route", "show", "default")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "default via ") {
			fields := strings.Fields(line)
			fmt.Println(fields)
			if len(fields) >= 3 {
				ip := net.ParseIP(fields[2])
				if ip != nil {
					n.DefaultGw = ip
				}
				logs.Debug("获取IP")
				fmt.Println(n.DefaultGw)
				n.DefaultDevices = fields[4]
				logs.Debug("获取设备")
				fmt.Println(n.DefaultDevices)
				if n.DefaultGw == nil {
					return fmt.Errorf("网关数据获取异常: %s", fields[2])
				}
				return nil
			}
		}
	}
	return fmt.Errorf("找不到默认网关地址")
}
