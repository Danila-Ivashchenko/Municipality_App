package route

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) Update(c *gin.Context) {
	req := &updateRouteRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	route := context_paylod_parser.GetRouteFromContext(ctx)
	if route == nil {
		response.Error(c, errors.New("route not found"))
		return
	}

	data := req.Convert(route.ID, route.PartitionID)

	routeUpdated, err := h.Params.RouteService.UpdateToPartition(ctx, &data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewRouteView(routeUpdated)
	response.Response(c, v)
}
