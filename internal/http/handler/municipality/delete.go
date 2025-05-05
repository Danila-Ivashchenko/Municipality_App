package municipality

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Delete(c *gin.Context) {

	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality not found"))
		return
	}

	err := h.Params.MunicipalityService.Delete(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
