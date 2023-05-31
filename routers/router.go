package routers

import (
	"go-gin-example/pkg/setting"
	v1 "go-gin-example/routers/api/v1"

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
	//v1 路由
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.DELETE("tags/:id", v1.DeleteTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
	}

	return r
}
