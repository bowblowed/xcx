package controller

import (
	"back-end/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleFileContent(c *gin.Context) {
	parm := c.Param("file")
	c.File("./file/" + parm)
}

func HandleGetFile(c *gin.Context) {
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
	pic, err := service.GetFile(parm.ID)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: pic,
	})
}

func HandleUpdatePic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  err.Error(),
		})
		return
	}
	ID := c.PostForm("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UpdateFile(uint(id), file)
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
