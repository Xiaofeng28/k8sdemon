package config

// Mysql 数据库配置参数映射结构体
type Mysql struct {
	// 数据库地址
	Host string `json:"host" yaml:"host"`
	// 数据库端口
	Port int `json:"port" yaml:"port"`
	// 连接的用户名
	UserName string `json:"username" yaml:"username"`
	// 连接的用户密码
	Password string `json:"password" yaml:"password"`
	// 连接的数据库名
	DbName string `json:"dbname" yaml:"dbname"`
}
