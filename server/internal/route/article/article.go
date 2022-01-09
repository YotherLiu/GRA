package article

import (
	v1 "server/internal/api/v1/article"

	"github.com/gin-gonic/gin"
)

func ArticleRoute(r *gin.RouterGroup, engine *gin.Engine) {
	article := r.Group("/article")

	article.POST("/add", v1.InsertArticle)
	article.DELETE("/del", v1.DeleteArticle)
	article.PUT("/update", v1.UpdateArticle)
	article.GET("/get", v1.SelectArticle)
}
