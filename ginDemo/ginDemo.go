package ginDemo

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var ServePort string

func RunGinDemo() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		log.Println("get /ping request")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", ServePort))
}
