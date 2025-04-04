package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Logout(c *gin.Context) {
	ctx := parser.Context(c)

	token := parser.ParseUserAuthToken(c)
	if token == nil {
		return
	}

	err := h.Params.UserService.Logout(ctx, token)
	if err != nil {
		response.Error(c, err)
		return
	}
}
