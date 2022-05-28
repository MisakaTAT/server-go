package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		global.Zap.Panicf("Failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(
		models.User{},
	); err != nil {
		global.Zap.Panicf("Auto migrate failed: %v", err)
	}

	return db
}
