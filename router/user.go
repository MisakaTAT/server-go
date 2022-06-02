package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	UserRouter.Use(middleware.JWTAuth())
	{

	}
	UserRouter.POST("/register", v1.Register)
}
