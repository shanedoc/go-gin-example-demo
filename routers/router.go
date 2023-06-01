package routers

import (
	"go-gin-example/middleware/jwt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers/api"
	v1 "go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	//v1 路由
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		//标签
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.DELETE("tags/:id", v1.DeleteTag)
		apiV1.PUT("/tags/:id", v1.EditTag)

		//文章
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/article/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
