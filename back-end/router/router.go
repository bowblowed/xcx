package router

import (
	"back-end/config"
	"back-end/controller"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(cors.New(cors.Config{
		// 允许所有来源访问
		AllowOrigins: []string{"*"},
		// 允许的请求方法
		AllowMethods: []string{"*"},
		// 允许的请求头
		AllowHeaders: []string{"*"},
	}))
	Router.POST("/wxLogin", controller.HandleWxLogin)
	Router.GET("/file/:file", controller.HandleFileContent)
	pic := Router.Group("/pic")
	{
		pic.POST("getSrc", controller.HandleGetFile)
	}

	banner := Router.Group("/banner")
	{
		banner.POST("/create", controller.HandleUploadBanner)
		banner.POST("/delete", controller.HandleDeleteBanner)
		banner.POST("/list", controller.HandleListBanner)
	}
	category := Router.Group("/category")
	{
		category.POST("/create", controller.HandleCreateCategory)
		category.POST("/delete", controller.HandleDeleteCategory)
		category.POST("/list", controller.HandleListCategory)
		category.POST("/get", controller.HandleGetCategory)
		category.POST("/uploadPic", controller.HandleUploadCategoryPic)

	}
	user := Router.Group("/user")
	{
		user.POST("create", controller.HandleCreateUser)
		user.POST("get", controller.HandleGetUser)
		user.POST("getInfo", controller.AuthMiddleware(), controller.HandleGetUserInfo)
		user.POST("getByOpenId", controller.HandleGetUserByOpenId)
		user.POST("list", controller.HandleListUser)
		user.POST("update", controller.HandleUpdateUser)
		shopCart := user.Group("shopCart", controller.AuthMiddleware())
		{
			shopCart.POST("get", controller.HandleGetUserShopCart)
			shopCart.POST("insertItems", controller.HandleUserShopCartInsertItems)
			shopCart.POST("deleteItem", controller.HandleUserShopCartDeleteItem)
			shopCart.POST("changeCount", controller.HandleUserShopCartChangeCount)
		}
	}
	product := Router.Group("product")
	{
		tag := product.Group("tag")
		{
			tag.POST("create", controller.HandleCreatePriceTag)
			tag.POST("list", controller.HandleListPriceTag)
			tag.POST("delete", controller.HandleDeletePriceTag)
			tag.POST("update", controller.HandleUpdatePriceTag)
		}
		product.POST("create", controller.HandleCreateProduct)
		product.POST("update", controller.HandleUpdateProduct)
		product.POST("get", controller.HandleGetProduct)
		product.POST("delete", controller.HandleDeleteProduct)
		product.POST("list", controller.HandleListProducts)
		product.POST("caculatePrice", controller.HandleCaculatePrice)
		product.POST("listById", controller.HandleProductListById)
		product.POST("listByCategoryId", controller.HandleListProductByCategoryId)
	}
	order := Router.Group("/order")
	{
		order.POST("/create", controller.AuthMiddleware(), controller.HandleCreateOrder)
		order.POST("/delete", controller.HandleDeleteOrder)
		order.POST("/update", controller.HandleUpdateOrder)
	}
}

func RunRouter() {

	Router.RunTLS(fmt.Sprintf(":%d", config.Server.Port), "./mydomain.crt", "./mydomain.key")
}
