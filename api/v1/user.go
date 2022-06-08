package v1

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/middleware"
	"server/models"
	"server/service"
	"server/utils"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.FailWithMsg(utils.Translate(err), c)
		return
	}
	// 检查用户名是否被注册
	if service.HasRegister(user.Username) {
		utils.FailWithMsg("用户名已存在", c)
		return
	}
	// 添加用户到数据库
	if err := service.Register(&user); err != nil {
		utils.FailWithMsg("注册失败", c)
		global.Zap.Errorf("user register failed: %v", err)
		return
	}
	utils.OkWithMsg("注册成功", c)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	if uid := middleware.ParseUUID(c); uid != uuid.Nil {
		user, err := service.GetUserInfo(uid)
		if err != nil {
			utils.FailWithMsg("用户信息获取失败", c)
			return
		}
		utils.OkWithData(user, c)
		return
	}
	utils.FailWithMsg("UUID解析失败", c)
}
