package service

import (
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/models"
	"server/structs"
)

// Register 用户注册
func Register(user *models.User) error {
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// HasRegister 检查用户是否存在
func HasRegister(username string) bool {
	var user models.User
	global.DB.Where("username = ?", username).First(&user)
	if user.UUID != uuid.Nil {
		return true
	}
	return false
}

// GetUserInfo 获取用户信息
func GetUserInfo(uuid uuid.UUID) (info structs.UserInfoResp, error error) {
	if err := global.DB.Model(&models.User{}).Where("uuid = ?", uuid).First(&info).Error; err != nil {
		return info, err
	}
	return info, nil
}
