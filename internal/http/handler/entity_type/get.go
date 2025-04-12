package entity_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetAll(c *gin.Context) {
	ctx := parser.Context(c)

	entityTypes, err := h.Params.EntityTypeService.GetAll(ctx)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newEntityTypeViews(entityTypes)
	response.Response(c, view)
}
