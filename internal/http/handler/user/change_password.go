package user

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) ChangePassword(c *gin.Context) {
	req := &changeUserPasswordRequest{}

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

	err = h.Params.UserService.ChangeUserPassword(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
