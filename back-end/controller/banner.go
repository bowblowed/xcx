package controller

import (
	"back-end/service"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleUploadBanner(c *gin.Context) {
	file, err := c.FormFile("banner")
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	name := "banner_" + fmt.Sprint(time.Now().Unix())
	err = c.SaveUploadedFile(file, "./file/"+name)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	service.CreateBanner(baseUrl + name)
	SetResponse(c, Resp{
		Code: ResponseSuccess,
	})
}

func HandleGetBanner(c *gin.Context) {
	var parm struct {
		Id uint `form:"id"`
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	banner, err := service.GetBanner(parm.Id)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: banner,
	})
}
func HandleListBanner(c *gin.Context) {
	banners, err := service.ListBanners()
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: banners,
	})
}

func HandleDeleteBanner(c *gin.Context) {
	var parm struct {
		Id uint `form:"id"`
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.DeleteBanner(parm.Id)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
	})
}
