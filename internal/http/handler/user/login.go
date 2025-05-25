package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Login(c *gin.Context) {

	req := &loginUserRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	data := req.Convert()

	user, err := h.Params.UserService.Login(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newUserTokenView(user)
	response.Response(c, view)
}
