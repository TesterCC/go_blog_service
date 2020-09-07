package main

import (
	"github.com/gin-gonic/gin"
)

// curl http://127.0.0.1:8080/ping

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":7777")  // default port 8080
}
