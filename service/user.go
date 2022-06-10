package service

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/structs"
)

// Register 用户注册
func Register(user *models.User) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
}

// HasRegister 检查用户是否存在
func HasRegister(username string) bool {
	var user models.User
	global.DB.Where("username = ?", username).First(&user)
	return user.UUID != uuid.Nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(uuid uuid.UUID) (info structs.UserInfoResp, error error) {
	if err := global.DB.Model(&models.User{}).Where("uuid = ?", uuid).First(&info).Error; err != nil {
		return info, err
	}
	return info, nil
}
