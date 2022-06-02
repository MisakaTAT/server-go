package initialize

import (
	"fmt"
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
		panic(fmt.Errorf("viper init failed: %v", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		global.Zap.Infof("config file changed: %s", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			global.Zap.Errorf("config unmarshal failed: %v", err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		global.Zap.Errorf("config unmarshal failed: %v", err)
	}

	return v
}
