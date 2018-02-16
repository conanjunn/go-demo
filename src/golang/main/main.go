package main

import (
	"github.com/gin-gonic/gin"
	"golang/service"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		str := service.MainFn()
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.Run()
}
