package user_admin

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) Update(c *gin.Context) {
	req := &reqUpdateUser{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	data := req.Convert()

	user, err := h.Params.UserAdminService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewUserExView(user)
	response.Response(c, v)
}
