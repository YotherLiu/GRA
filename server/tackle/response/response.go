package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一处理响应信息

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "ok",
	})
}

func Success() {

}

func Fail() {

}

func FailForMessage() {

}
