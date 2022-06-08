package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ok   = 0
	fail = -1
)

func result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	result(ok, "", nil, c)
}

func OkWithMsg(msg string, c *gin.Context) {
	result(ok, msg, nil, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	result(ok, "", data, c)
}

func OkWithDetailed(msg string, data interface{}, c *gin.Context) {
	result(ok, msg, data, c)
}

func Fail(c *gin.Context) {
	result(fail, "", nil, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	result(fail, msg, nil, c)
}

func FailWithDetailed(msg string, data interface{}, c *gin.Context) {
	result(fail, msg, data, c)
}
