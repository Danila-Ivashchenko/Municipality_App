package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) CreateObjects(c *gin.Context) {
	req := &createMultiplyObjetsData{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	objectTemplate := context_paylod_parser.GetObjectTemplateFromContext(ctx)
	if objectTemplate == nil {
		response.Error(c, errors.New("object_template is nil"))
		return
	}

	data := req.Convert(objectTemplate.ID)

	objects, err := h.Params.ObjectService.CreateMultiply(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewObjectViews(objects)
	response.Response(c, v)
}
