package config

import (
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

// Config 配置参数的结构体映射
type Config struct {
	// 应用服务配置
	System System `json:"system" yaml:"system"`
	// 数据库配置
	Mysql Mysql `json:"mysql" json:"mysql"`
}

// GetConfig 获取config.yaml文件里的配置参数
func GetConfig(path string) (*Config, error) {
	//1.读取文件
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//2.映射到结构体
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	//3.返回映射结构体和错误
	return &config, nil
}

// DB 数据库连接
var DB *gorm.DB = nil

// InitMysqlConnect 连接数据库
func InitMysqlConnect(config Config) error {
	// 数据库连接串
	connectString := config.Mysql.UserName + ":" + config.Mysql.Password +
		"@tcp(" + config.Mysql.Host + ":" + strconv.Itoa(config.Mysql.Port) + ")/" + config.Mysql.DbName +
		"?" + "charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	dbConnect, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return err
	}
	// 实例化数据库连接
	DB = dbConnect
	return nil
}
