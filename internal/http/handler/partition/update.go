package partition

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) UpdatePartition(c *gin.Context) {
	req := &reqPartitionUpdateData{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	//if validateErr := req.Validate(); validateErr != nil {
	//	response.Error(c, validateErr)
	//	return
	//}

	partition := context_paylod_parser.GetPartitionFromContext(ctx)
	if partition == nil {
		response.Error(c, errors.New("partition not found"))
		return
	}

	data := req.Convert(partition.ID)

	partitionUpdated, err := h.Params.PartitionService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewPartitionExView(*partitionUpdated)
	response.Response(c, v)
}
