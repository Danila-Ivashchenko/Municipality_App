package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Register(c *gin.Context) {

	req := &registerUserRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	data := req.Convert()

	user, err := h.Params.UserService.RegisterUser(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newUserView(user)
	response.Response(c, view)
}
