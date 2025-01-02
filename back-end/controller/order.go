package controller

import (
	"back-end/model"
	"back-end/service"

	"github.com/gin-gonic/gin"
)

func HandleCreateOrder(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	var parm model.Order
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	parm.UserId = user.ID
	err = service.CreateOrder(parm)
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

func HandleDeleteOrder(c *gin.Context) {
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
	err = service.DeleteOrder(parm.ID)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
}

func HandleUpdateOrder(c *gin.Context) {
	var parm model.Order
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UpdateOrder(parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
	}
}
