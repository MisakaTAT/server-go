package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"server/utils"
)

type User struct {
	Model
	Username string `json:"username" gorm:"unique;not null;comment:用户名" binding:"required"`
	Nickname string `json:"nickname" gorm:"not null;comment:昵称" binding:"required"`
	Password string `json:"password" gorm:"not null;comment:密码" binding:"required"`
	Email    string `json:"email" gorm:"not null;comment:邮箱地址" binding:"required"`
	Avatar   string `json:"avatar" gorm:"comment:头像链接" binding:"omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	password, err := utils.ScryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.UUID = uuid.NewV4()
	u.Password = password
	return
}
