package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Dokku and Oracle Cloud is working",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Caldwell",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
