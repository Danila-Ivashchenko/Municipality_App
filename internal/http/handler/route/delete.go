package route

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/domain/service"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Delete(c *gin.Context) {
	ctx := parser.Context(c)

	route := context_paylod_parser.GetRouteFromContext(ctx)
	if route == nil {
		response.Error(c, errors.New("route not found"))
		return
	}

	err := h.Params.RouteService.DeleteToPartition(ctx, &service.DeleteRoutesToPartitionData{
		PartitionID: route.PartitionID,
		RoutesID:    route.ID,
	})

	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
