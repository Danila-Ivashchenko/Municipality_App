package region

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetByParams(c *gin.Context) {
	req := &getByParamsReq{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	data := req.Convert()

	regions, err := h.Params.RegionService.GetByParams(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newRegionsView(regions)
	response.Response(c, view)
}

func (h *Handler) Get(c *gin.Context) {}
