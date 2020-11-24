package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ip2region/router"
	"github.com/sirupsen/logrus"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	// route
	r := router.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		logrus.Panicf("Server Start Panic", err)
	}
}
