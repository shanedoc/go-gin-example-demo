package routers

import (
	"go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	return r
}
