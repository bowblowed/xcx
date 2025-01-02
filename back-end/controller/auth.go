package controller

import (
	"back-end/model"
	"back-end/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			SetResponse(c, Resp{
				Code: ResponseAuthError,
				Msg:  "no token",
			})
			c.Abort()
			return
		}
		user, err := service.Auth(authHeader)
		if err != nil {
			SetResponse(c, Resp{
				Code: ResponseAuthError,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user")
		if user.(model.User).Type != model.UserTypeAdmin {
			SetResponse(c, Resp{
				Code: ResponseAuthError,
				Msg:  "no admin",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func HandleGetUserInfo(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	SetResponse(c, Resp{
		Code: ResponseSuccess,
		Data: user,
	})
}
