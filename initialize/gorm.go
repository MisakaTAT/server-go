package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}

	if err = db.AutoMigrate(
		models.User{},
	); err != nil {
		panic(fmt.Errorf("auto migrate failed: %v", err))
	}

	return db
}
