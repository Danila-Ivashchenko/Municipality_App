package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) GetEntitys(c *gin.Context) {

	ctx := parser.Context(c)

	objectTemplate := context_paylod_parser.GetEntityTemplateFromContext(ctx)
	if objectTemplate == nil {
		response.Error(c, errors.New("object_template is nil"))
		return
	}

	objects, err := h.Params.EntityService.GetExByTemplateID(ctx, objectTemplate.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewEntityViews(objects)
	response.Response(c, v)
}
