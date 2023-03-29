package glbm

import (
	"gitee.com/liumou_site/gcs"
	"gitee.com/liumou_site/gf"
	"net"
)

// ApiApt 定义一个命令行密码结构体
type ApiApt struct {
	password string       // sudo权限使用的密码
	Sudo     *gcs.ApiSudo // 执行sudo命令
	Err      error        // 错误信息
	dpkg     *ApiDpkg     // Dpkg实例
	Debug    bool         // 是否开启Debug
	Info     bool         // 是否开启Info
}

// ApiDpkg 定义一个Dpkg管理结构
type ApiDpkg struct {
	password string       // 主机密码
	Sudo     *gcs.ApiSudo // 执行sudo命令
	Err      error        // 错误信息
	Result   bool         // 操作结果
	Debug    bool         // 是否开启Debug
	Info     bool         // 是否开启Info
}

// ApiNmcli 网卡管理
type ApiNmcli struct {
	// status         string          // 连接状态
	conType        string   // 连接类型
	device         string   // 设备名称
	devList        []string // 设备列表
	connectionList []string // 连接列表
	deviceQuantity int      // 设备数量
	DefaultGw      net.IP   // 当前默认网关
	DefaultDevices string   // 当前默认网卡设备
	// subnet         string          // 当前网段
	address net.IP      // 当前IP
	mask    int         // 当前子网掩码
	mac     string      // 当前网卡Mac地址
	cmd     gcs.ApiSudo // 设置实例
	Err     error       // 错误信息
}
type EthInfo struct {
	IP   net.IP   // IP地址
	GW   net.IP   // 网关地址
	MASK int      // 子网掩码
	DNS  []net.IP // DNS列表
}

// ApiEth 网卡信息结构体
type ApiEth struct {
	Index int    // 网卡设备索引
	Name  string // 网卡名称
	Ipv4  net.IP // IPV4地址
	Mask  int    //子网掩码
	Mac   string // Mac物理地址
	Err   error  // 错误信息
}

type ApiConnection struct {
	Name       string        // 连接名称(例如: dhcp)
	Types      string        // 连接类型(bridge/wifi/ethernet)
	Method     string        // 连接模式(auto)
	uuid       string        // 连接UUID
	Dns        []string      // DNS列表
	Dev        string        // 设备名称
	Gw         net.IP        // 网关地址
	Address    net.IP        // IP地址
	Mask       int           // 子网掩码
	Err        error         // 错误信息
	cmd        *gcs.ApiShell // 命令实例
	UseConName string        // 正在使用的连接名称
	ConList    []string      // 连接列表
}

type ApiService struct {
	Name     string       // 服务名称
	Err      error        // 错误
	Status   string       // 当前状态
	Password string       // Sudo权限密码
	sudo     *gcs.ApiSudo // Sudo实例
}

type ApiFile struct {
	PathAbs    string         // 操作对象绝对路径
	PathBase   string         // 操作对象基础文件名
	PathFormat string         // 操作对象文件格式
	Src        string         // 源文件
	SrcAbs     string         // 源文件绝对路径
	SrcBase    string         // 源文件基础文件名
	SrcFormat  string         // 源文件格式
	Dst        string         // 目标文件
	DstAbs     string         // 目标文件绝对路径
	DstBase    string         // 目标文件基础文件名
	DstFormat  string         // 目标文件格式
	shell      *gcs.ApiShell  // 命令实例
	Err        error          // 错误
	fileMan    *gf.ApiFileMan // 文件管理
}

type ApiFileSudo struct {
	Password string       // 密码
	sudo     *gcs.ApiSudo // Sudo实例
	ApiFile               // 继承文件结构
}
