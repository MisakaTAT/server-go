package config

// Config 配置文件结构体
type Config struct {
	Server Server `json:"server" yaml:"server"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Jwt    Jwt    `json:"jwt" yaml:"jwt"`
}
