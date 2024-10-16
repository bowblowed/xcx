package controller

import (
	"back-end/service"

	"github.com/gin-gonic/gin"
)

func HandleWxLogin(c *gin.Context) {
	var parm struct {
		Code string `form:"code"`
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	token, err := service.WxLogin(parm.Code)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: token,
	})
}
