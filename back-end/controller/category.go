package controller

import (
	"back-end/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetCategory(c *gin.Context) {
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
	category, err := service.GetCategory(uint(parm.ID))
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: category,
	})
}

func HandleListCategory(c *gin.Context) {
	categories, err := service.ListCategories()
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: categories,
	})
}

func HandleDeleteCategory(c *gin.Context) {
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
	err = service.DeleteCategory(uint(parm.ID))
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

func HandleCreateCategory(c *gin.Context) {
	file, err := c.FormFile("file")
	Name := c.PostForm("Name")
	if Name == "" {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	if err == nil {
		err = service.CreateCategory(Name, file)
	} else {
		err = service.CreateCategory(Name)
	}
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
func HandleUploadCategoryPic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	idstr := c.PostForm("ID")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UploadCategoryPic(file, uint(id))
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

func HandleListProductByCategoryId(c *gin.Context) {
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
	products, err := service.ListProductsByCategory(uint(parm.ID))
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: products,
	})
}
