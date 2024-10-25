package controller

import (
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
		token := authHeader
		openId, err := service.ParseToken(token)
		if err != nil {
			SetResponse(c, Resp{
				Code: ResponseAuthError,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		user, err := service.GetUserByOpenId(openId)
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
