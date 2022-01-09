package system

import (
	v1 "server/internal/api/v1/system"

	"github.com/gin-gonic/gin"
)

func System(r *gin.RouterGroup, engine *gin.Engine) {
	r.GET("/ping", v1.Ping)
}
