package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Me(c *gin.Context) {
	ctx := parser.Context(c)

	user := context_paylod_parser.GetUserFromContext(ctx)

	if user == nil {
		return
	}

	view := newUserView(user)
	response.Response(c, view)
}
