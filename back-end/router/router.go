package router

import (
	"back-end/config"
	"back-end/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.POST("/wxLogin", controller.HandleWxLogin)
	banner := Router.Group("/banner")
	{
		banner.POST("/upload", controller.HandleUploadBanner)
		banner.POST("/delete", controller.HandleDeleteBanner)
		banner.POST("/list", controller.HandleListBanner)
		banner.POST("/get", controller.HandleGetBanner)
	}
	category := Router.Group("/category")
	{
		category.POST("/create", controller.HandleCreateCategory)
		category.POST("/delete", controller.HandleDeleteCategory)
		category.POST("/list", controller.HandleListCategory)
		category.POST("/get", controller.HandleGetCategory)
		category.POST("/upload", controller.HandleUploadBanner)
	}
	user := Router.Group("/user")
	{
		user.POST("/create", controller.HandleCreateUser)
		user.POST("/get", controller.HandleGetUser)
		user.POST("/getByOpenId", controller.HandleGetUserByOpenId)
		user.POST("/list", controller.HandleListUser)
		user.POST("/update", controller.HandleUpdateUser)
	}
}

func RunRouter() {
	Router.Run(fmt.Sprintf(":%d", config.Server.Port))
}
