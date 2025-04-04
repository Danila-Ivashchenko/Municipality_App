package object_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetAll(c *gin.Context) {
	ctx := parser.Context(c)

	objectTypes, err := h.Params.ObjectTypeService.GetAll(ctx)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newObjectTypeViews(objectTypes)
	response.Response(c, view)
}
