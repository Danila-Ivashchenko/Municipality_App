package municipality

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Create(c *gin.Context) {
	req := &createMunicipalityRequest{}

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

	municipalityEx, err := h.Params.MunicipalityService.Create(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newMunicipalityExView(municipalityEx)
	response.Response(c, view)
}
