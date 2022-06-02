package service

import (
	"server/global"
	"server/models"
	"server/utils"
)

// LoginCheck 登录检查
func LoginCheck(username string, password string) (user models.User, ok bool, error error) {
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, false, err
	}
	if scryptPassword, err := utils.ScryptPassword(password); err == nil {
		if user.Username == username && user.Password == scryptPassword {
			return user, true, nil
		}
	}
	return user, false, nil
}
