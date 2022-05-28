package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	DB     *gorm.DB
	Viper  *viper.Viper
	Zap    *zap.SugaredLogger
	CONFIG *config.Config
)
