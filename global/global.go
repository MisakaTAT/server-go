package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	DB     *gorm.DB
	ZAP    *zap.SugaredLogger
	CONFIG *config.Config
)
