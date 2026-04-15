package config

// System 配置参数的结构体映射
type System struct {
	// 监听地址
	Host string `json:"host" yaml:"host"`
	// 监听端口
	Port int `json:"port" yaml:"port"`
}
