package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

// 路由总入口，注册所有微服务的 路由
func Register(router *gin.Engine)  {
	//配置跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.HandleMethodNotAllowed = true
	// 使用gin-swagger 中间件
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册 user router
	RegisterUser(router)

}
