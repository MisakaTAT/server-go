package models

type User struct {
	Model
	Username string `json:"username" gorm:"unique;not null;comment:用户名" binding:"required"`
	Password string `json:"password" gorm:"not null;comment:密码" binding:"required"`
	Email    string `json:"email" gorm:"not null;comment:邮箱地址" binding:"required"`
	Avatar   string `json:"avatar,omitempty" gorm:"comment:头像链接" binding:"omitempty"`
}
