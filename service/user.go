package service

import (
	"gorm.io/gorm"
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

// DeleteUser 删除用户
func DeleteUser(id int) (error error) {
	var user models.User
	if err := global.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
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

// UpdateUser 更新用户信息
func UpdateUser(id int, u models.User) error {
	if err := global.DB.Model(&models.User{}).Where("id = ?", id).Save(u).Error; err != nil {
		return err
	}
	return nil
}

// QueryUserList TODO: 查询所有用户 此方法需要重写
func QueryUserList(pageSize int, pageNum int) (int64, []models.User, error) {
	var totalRows int64
	var users []models.User
	err := global.DB.Model(&models.User{}).Count(&totalRows).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Select([]string{"id", "username"}).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, nil, err
	}
	return totalRows, users, nil
}
