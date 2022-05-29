package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/resp"
	"server/common/utils"
	"server/models"
	"server/service"
	"server/structs"
	"strconv"
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

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteUser(id); err != nil {
		resp.Result(resp.Failed, fmt.Sprintf("用户删除失败：%s", err.Error()), nil, c)
		return
	}
	resp.Result(resp.Succeed, "用户删除成功", nil, c)
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {

}

// QueryUser 查询用户
func QueryUser(c *gin.Context) {

}

// QueryUserList 查询所有用户
func QueryUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	totalRows, userList, err := service.QueryUserList(pageSize, pageNum)
	if err != nil {
		resp.Result(resp.Failed, fmt.Sprintf("用户查询失败：%s", err.Error()), nil, c)
		return
	}
	// 封装分页数据
	data := structs.Pagination{
		TotalRows: totalRows,
		Data:      userList,
	}
	resp.Result(resp.Succeed, "", data, c)
}
