package admin_user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetAll(c *gin.Context) {
	ctx := parser.Context(c)

	users, err := h.Params.UserService.Get

	view := newUserView(user)
	response.Response(c, view)
}
