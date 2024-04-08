package routers

import (
	"file_store/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter() *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())

	engine.Use(gin.Recovery())

	r := engine.Group("/manager")
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		go func() {
			r.POST("/upload", controllers.UploadImage)
		}()

		go func() {
			r.StaticFS("/shader/huythang", http.Dir("shader"))
		}()
	}

	return &ApiRouter{
		Engine: engine,
	}
}
