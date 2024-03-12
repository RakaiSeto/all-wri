package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("HELLO WORLD")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
