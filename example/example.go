package main

import (
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(limit.Limit(200)) // limit the number of current requests

	r.GET("/", func(c *gin.Context) {
		// your code
	})

	r.Run()
}
