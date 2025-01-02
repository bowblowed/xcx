package controller

import (
	"back-end/model"
	"back-end/service"

	"github.com/gin-gonic/gin"
)

func HandleCreateUser(c *gin.Context) {
	var parm model.User
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.CreateUser(parm)
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

func HandleGetUser(c *gin.Context) {
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
	user, err := service.GetUser(parm.ID)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: user,
	})
}

func HandleGetUserByOpenId(c *gin.Context) {
	var parm struct {
		OpenId string `form:"openId"`
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	user, err := service.GetUserByOpenId(parm.OpenId)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: user,
	})
}

func HandleListUser(c *gin.Context) {
	users, err := service.ListUsers()
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: users,
	})
}

func HandleUpdateUser(c *gin.Context) {
	var parm model.User
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UpdateUser(parm.ID, parm)
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

func HandleGetUserShopCart(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	shopCart, err := service.GetUserShopCart(user)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: shopCart,
	})
}

func HandleUserShopCartInsertItems(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	var parm []model.OrderItem

	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UserShopCartInsertItems(user, parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Msg:  "success",
	})
}

func HandleUserShopCartDeleteItem(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	var parm uint
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UserShopCartDeleteItem(user, parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Msg:  "success",
	})
}

func HandleUserShopCartChangeCount(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	var parm struct {
		Idx   uint
		Count uint
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
	}
	err = service.UserShopCartChangeCount(user, parm.Idx, parm.Count)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Msg:  "success",
	})
}
