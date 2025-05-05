package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, err.Error())
}

func Unauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, err.Error())
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, err.Error())
}

func Response(c *gin.Context, v any) {
	c.JSON(http.StatusOK, v)
}

func ResponseNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
