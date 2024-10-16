package router

import (
	"back-end/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.GET("/wxLogin", controller.HandleWxLogin)
}

func RunRouter() {
	Router.Run()
}
