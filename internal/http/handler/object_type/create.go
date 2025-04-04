package object_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Create(c *gin.Context) {
	req := &createObjectTypesReq{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	data := req.Convert()

	objectTypes, err := h.Params.ObjectTypeService.CreateMultiply(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newObjectTypeViews(objectTypes)
	response.Response(c, view)
}
