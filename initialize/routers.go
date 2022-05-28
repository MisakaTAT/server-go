package initialize

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/router"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		router.InitAuthRouter(v1)
		router.InitUserRouter(v1)
	}
	return r
}
