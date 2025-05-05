package entity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) CreateEntities(c *gin.Context) {
	req := &createMultiplyEntitiesData{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	objectTemplate := context_paylod_parser.GetEntityTemplateFromContext(ctx)
	if objectTemplate == nil {
		response.Error(c, errors.New("object_template is nil"))
		return
	}

	data := req.Convert(objectTemplate.ID)

	objects, err := h.Params.EntityService.CreateMultiply(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewEntityViews(objects)
	response.Response(c, v)
}
