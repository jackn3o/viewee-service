package api

import (
	"log"
	"time"

	"viewee-service/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() *gin.Engine {
	log.Println("configuring server...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	return initializeRoutes(r)
}

func initializeRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	jwtMw := middleware.JWTAuth()
	test := r.Group("/test")
	test.Use(jwtMw)

	return r
}
