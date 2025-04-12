package entity_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Create(c *gin.Context) {
	req := &createEntityTypesReq{}

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

	entityTypes, err := h.Params.EntityTypeService.CreateMultiply(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newEntityTypeViews(entityTypes)
	response.Response(c, view)
}
