package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ResponseSuccess uint = 0
var ResponseServerError uint = 1
var ResponseParmError uint = 2
var ResponseAuthError uint = 3

type Resp struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SetResponse(c *gin.Context, r Resp) {
	c.JSON(http.StatusOK, r)
}
