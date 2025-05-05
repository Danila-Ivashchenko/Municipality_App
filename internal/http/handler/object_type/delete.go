package object_type

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Delete(c *gin.Context) {
	req := &deleteObjectTypeReq{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if req.ID == 0 {
		response.Error(c, errors.New("id is required"))
		return
	}

	err = h.Params.ObjectTypeService.Delete(ctx, []int64{req.ID})
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
