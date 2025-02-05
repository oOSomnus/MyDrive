package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET(
		"/ping", func(c *gin.Context) {
			c.JSON(
				200, gin.H{
					"message": "pong",
				},
			)
		},
	)
	r.Run(":8080")
}
