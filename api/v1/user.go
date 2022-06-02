package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/resp"
	"server/common/utils"
	"server/models"
	"server/service"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Result(resp.Failed, utils.Translate(err), nil, c)
		return
	}
	// 检查用户名是否被注册
	if service.CheckUserExist(user.Username) {
		resp.Result(resp.Failed, "用户名已存在", nil, c)
		return
	}
	// 添加用户到数据库
	if err := service.Register(&user); err != nil {
		resp.Result(resp.Failed, fmt.Sprintf("注册失败：%s", err.Error()), nil, c)
		return
	}
	resp.Result(resp.Succeed, "注册成功", nil, c)
}
