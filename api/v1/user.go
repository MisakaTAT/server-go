package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"server/middleware"
	"server/models"
	"server/service"
	"server/utils"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Result(utils.Failed, utils.Translate(err), nil, c)
		return
	}
	// 检查用户名是否被注册
	if service.HasRegister(user.Username) {
		utils.Result(utils.Failed, "用户名已存在", nil, c)
		return
	}
	// 添加用户到数据库
	if err := service.Register(&user); err != nil {
		utils.Result(utils.Failed, fmt.Sprintf("注册失败：%s", err.Error()), nil, c)
		return
	}
	utils.Result(utils.Succeed, "注册成功", nil, c)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	if uid := middleware.ParseUUID(c); uid != uuid.Nil {
		user, err := service.GetUserInfo(uid)
		if err != nil {
			utils.Result(utils.Failed, "用户信息获取失败", nil, c)
			return
		}
		utils.Result(utils.Succeed, "", user, c)
		return
	}
	utils.Result(utils.Failed, "UUID解析失败", nil, c)
}
