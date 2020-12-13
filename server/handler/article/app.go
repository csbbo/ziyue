package article

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine) {
	article := Article{}
	r := app.Group("/api")
	{
		r.POST("/article", article.Create)
		r.PUT("/article", article.Update)
		r.GET("/article", article.Get)
		r.DELETE("/article", article.Delete)
		r.GET("/articlelist", article.GetList)
		r.POST("/articleparse", article.ParseArticleFile)
		r.GET("showarticles", article.ShowArticles)
	}
}
