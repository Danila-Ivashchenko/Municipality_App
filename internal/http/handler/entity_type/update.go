package entity_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Update(c *gin.Context) {
	req := &updateEntityTypeReq{}

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

	entityType, err := h.Params.EntityTypeService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newEntityTypeView(entityType)
	response.Response(c, view)
}
