package system

import (
	"server/tackle/response"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.Ok(c)
}
