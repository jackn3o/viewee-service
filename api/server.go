package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() *gin.Engine {
	log.Println("configuring server...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return initializeRoutes(r)
}

func initializeRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
