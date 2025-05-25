package user_admin

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) GetAll(c *gin.Context) {
	ctx := parser.Context(c)

	users, err := h.Params.UserAdminService.GetAll(ctx)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewUserExViews(users)
	response.Response(c, v)
}
