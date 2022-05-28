package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (m *Mysql) Dsn() string {
	const config = "?charset=utf8mb4&parseTime=True&loc=Local"
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBName + config
}
