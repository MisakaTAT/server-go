package config

// Jwt 配置文件结构体
type Jwt struct {
	Exp     int64  `yaml:"exp"`
	Iss     string `yaml:"iss"`
	SignKey string `yaml:"signKey"`
}
