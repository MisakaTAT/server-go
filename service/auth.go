package service

import (
	"server/common/utils"
	"server/global"
	"server/models"
)

// LoginCheck 登录检查
func LoginCheck(username string, password string) (user models.User, ok bool) {
	global.DB.Where("username = ?", username).First(&user)
	if scryptPassword, err := utils.ScryptPassword(password); err == nil {
		if user.Username == username && user.Password == scryptPassword {
			return user, true
		}
	}
	return user, false
}
