package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) DeleteEntities(c *gin.Context) {
	req := &deleteEntitysRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	objectTemplate := context_paylod_parser.GetEntityTemplateFromContext(ctx)
	if objectTemplate == nil {
		response.Error(c, errors.New("object_template is nil"))
		return
	}

	err = h.Params.EntityService.DeleteMultiple(ctx, req.IDs, objectTemplate.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
