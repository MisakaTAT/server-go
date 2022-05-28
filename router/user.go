package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	// UserRouter.Use(middleware.JWTAuth())
	{
		UserRouter.GET("/:id", v1.QueryUser)
		UserRouter.GET("/list", v1.QueryUserList)
		UserRouter.DELETE("/:id", v1.DeleteUser)
		UserRouter.PATCH("/update", v1.UpdateUser)
	}
	UserRouter.POST("/register", v1.Register)
}
