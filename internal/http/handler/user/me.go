package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) Me(c *gin.Context) {
	ctx := parser.Context(c)

	user := context_paylod_parser.GetUserFromContext(ctx)

	if user == nil {
		return
	}

	userEx, err := h.Params.UserService.Me(ctx, user.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewUserExView(userEx)
	response.Response(c, v)
}
