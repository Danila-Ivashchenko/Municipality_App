package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/domain/core_errors"
	"net/http"
)

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

func Error(c *gin.Context, err error) {
	var coreError *core_errors.DomainError
	if errors.As(err, &coreError) {
		c.JSON(coreError.Code, err.Error())
		return
	}

	c.JSON(http.StatusInternalServerError, "Что-то пошло не так...")
}
