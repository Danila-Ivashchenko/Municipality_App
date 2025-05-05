package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Update(c *gin.Context) {
	req := &updateUserRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	user := context_paylod_parser.GetUserFromContext(ctx)
	if user == nil {
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	data := req.Convert(user.ID)

	user, err = h.Params.UserService.UpdateUser(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := newUserView(user)
	response.Response(c, v)
}
