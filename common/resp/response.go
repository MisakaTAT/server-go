package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	Succeed      = 0   // 成功
	Failed       = -1  // 失败
	Unauthorized = 101 // 未授权
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
