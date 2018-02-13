package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	authorized := r.Group("/")
	{
		authorized.GET("/get", getFn)
	}

	r.Run(":9999")
}

func getFn() {
}
