package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ip2region/middleware"
	"github.com/go-ip2region/src"
	"github.com/sirupsen/logrus"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件(全局)
	// cors 跨域
	r.Use(middleware.CORS())
	IP := r.Group("/")
	{
		IP.GET("/ip2region", src.Ip2Region)
	}
	logrus.Info("Listening on ===> ",":8080")
	return r
}
