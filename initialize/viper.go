package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"server/global"
)

const defaultConfigFile = "config.yaml"

// Viper 初始化配置文件
func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	if err := v.ReadInConfig(); err != nil {
		global.Zap.Panicf("Fatal error config file: %v", err)
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		global.Zap.Infof("Config file changed: %s", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			global.Zap.Errorf("Config unmarshal failed: %v", err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		global.Zap.Errorf("Config unmarshal failed: %v", err)
	}

	return v
}
