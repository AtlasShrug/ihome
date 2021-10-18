package main

import (
	"github.com/gin-gonic/gin"
	"ihome/web/controller"
)

func main() {
	router := gin.Default()
	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)
	router.GET("api/v1.0/imagecode/:uuid", controller.GetImageCd)

	router.Run(":8888")
}
