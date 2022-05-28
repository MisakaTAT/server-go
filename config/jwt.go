package config

type Jwt struct {
	Exp     int64  `yaml:"exp"`
	Iss     string `yaml:"iss"`
	SignKey string `yaml:"signKey"`
}
