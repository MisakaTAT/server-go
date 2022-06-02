package service

import (
	"server/common/utils"
	"server/global"
	"server/models"
)

// Register 用户注册
func Register(user *models.User) (error error) {
	password, err := utils.ScryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (userExist bool) {
	var user models.User
	global.DB.Where("username = ?", username).Take(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
