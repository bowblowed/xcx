package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnautz horized, gin.H{"error": "未提供 token"})
			c.Abort()
			return
		}

		if !validateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
