package route

import (
	"server/internal/route/article"
	"server/internal/route/system"

	"github.com/gin-gonic/gin"
)

func InitRouter(route *gin.RouterGroup, engine *gin.Engine) {
	system.System(route, engine)
	article.ArticleRoute(route, engine)
}
