package region

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Create(c *gin.Context) {
	req := &createRegionReq{}

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

	region, err := h.Params.RegionService.Create(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newRegionView(region)
	response.Response(c, view)
}
