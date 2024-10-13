package main

import (
	"back-end/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.F()
	r := gin.Default()
	r.Run(":8080")
}
