package route

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) Create(c *gin.Context) {
	req := &createRouteRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	partition := context_paylod_parser.GetPartitionFromContext(ctx)
	if partition == nil {
		response.Error(c, errors.New("partition not found"))
		return
	}

	data := req.Convert(partition.ID)

	route, err := h.Params.RouteService.CreateToPartition(ctx, &data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewRouteView(route)
	response.Response(c, v)
}
