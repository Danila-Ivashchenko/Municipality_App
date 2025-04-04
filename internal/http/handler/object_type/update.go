package object_type

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Update(c *gin.Context) {
	req := &updateObjectTypeReq{}

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

	objectType, err := h.Params.ObjectTypeService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newObjectTypeView(objectType)
	response.Response(c, view)
}
