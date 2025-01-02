package controller

import (
	"back-end/service"
	"encoding/json"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleDeleteProduct(c *gin.Context) {
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
	err = service.DeleteProduct(parm.ID)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
}

func HandleCreateProduct(c *gin.Context) {
	Name := c.PostForm("Name")
	CategoryIdStr := c.PostForm("CategoryId")
	CategoryId, err := strconv.Atoi(CategoryIdStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	DefaultPriceStr := c.PostForm("DefaultPrice")
	DefaultPrice, err := strconv.Atoi(DefaultPriceStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	CountStr := c.PostForm("Count")
	Count, err := strconv.Atoi(CountStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	Description := c.PostForm("Description")
	picFile, err := c.FormFile("PicFile")
	if err != nil {
		picFile = nil
	}
	var DescriptionFiles []*multipart.FileHeader = []*multipart.FileHeader{}
	DescriptionFileLenStr := c.PostForm("DescriptionFilesLen")
	if DescriptionFileLenStr != "" {
		DescriptionFileLen, err := strconv.Atoi(DescriptionFileLenStr)
		if err != nil {
			SetResponse(c, Resp{
				Code: ResponseParmError,
				Msg:  "parm error",
			})
			return
		}
		for i := 0; i < DescriptionFileLen; i++ {
			tmp, err := c.FormFile("DescriptionFiles_" + strconv.Itoa(i))
			if err != nil {
				continue
			}
			DescriptionFiles = append(DescriptionFiles, tmp)
		}
	}
	propsStr := c.PostForm("Props")
	var props = make(map[string][]string)
	err = json.Unmarshal([]byte(propsStr), &props)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.CreateProduct(Name, uint(CategoryId), uint(DefaultPrice), uint32(Count), Description, picFile, DescriptionFiles, props)
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

func HandleUpdateProduct(c *gin.Context) {
	Idstr := c.PostForm("ID")
	Id, err := strconv.Atoi(Idstr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	Name := c.PostForm("Name")
	CategoryIdStr := c.PostForm("CategoryId")
	CategoryId, err := strconv.Atoi(CategoryIdStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	DefaultPriceStr := c.PostForm("DefaultPrice")
	DefaultPrice, err := strconv.Atoi(DefaultPriceStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	CountStr := c.PostForm("Count")
	Count, err := strconv.Atoi(CountStr)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	Description := c.PostForm("Description")
	picFile, err := c.FormFile("PicFile")
	if err != nil {
		picFile = nil
	}
	var DescriptionFiles []*multipart.FileHeader = []*multipart.FileHeader{}
	DescriptionFileLenStr := c.PostForm("DescriptionFilesLen")
	if DescriptionFileLenStr != "" {
		DescriptionFileLen, err := strconv.Atoi(DescriptionFileLenStr)
		if err != nil {
			SetResponse(c, Resp{
				Code: ResponseParmError,
				Msg:  "parm error",
			})
			return
		}
		for i := 0; i < DescriptionFileLen; i++ {
			tmp, err := c.FormFile("DescriptionFiles_" + strconv.Itoa(i))
			if err != nil {
				continue
			}
			DescriptionFiles = append(DescriptionFiles, tmp)
		}
	}
	propsStr := c.PostForm("Props")
	var props = make(map[string][]string)
	err = json.Unmarshal([]byte(propsStr), &props)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	err = service.UpdateProduct(uint(Id), Name, uint(CategoryId), uint(DefaultPrice), uint32(Count), Description, picFile, DescriptionFiles, props)
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

func HandleListProducts(c *gin.Context) {
	list, err := service.ListProducts()
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: list,
	})
}

func HandleGetProduct(c *gin.Context) {
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
	product, err := service.GetProduct(uint(parm.ID))
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: product,
	})
}

func HandleCaculatePrice(c *gin.Context) {
	var parm struct {
		ProductId uint
		Propm     map[string]string
		Count     uint
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	price, err := service.ProductCaculatePrice(parm.ProductId, parm.Propm, parm.Count)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: price,
	})
}

func HandleProductListById(c *gin.Context) {
	var parm struct {
		IDs []uint
	}
	err := c.ShouldBindBodyWithJSON(&parm)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseParmError,
			Msg:  "parm error",
		})
		return
	}
	productList, err := service.ListProductsByIds(parm.IDs)
	if err != nil {
		SetResponse(c, Resp{
			Code: ResponseServerError,
			Msg:  err.Error(),
		})
		return
	}
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: productList,
	})
}
