package controller

import (
	"back-end/service"

	"github.com/gin-gonic/gin"
)

func HandleUploadBanner(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  err.Error(),
		})
		return
	}
	err = service.CreateBanner(file)
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
		ID uint
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.DeleteBanner(parm.ID)
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
