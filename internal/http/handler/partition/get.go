package partition

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) GetPartition(c *gin.Context) {
	ctx := parser.Context(c)

	partition := context_paylod_parser.GetPartitionFromContext(ctx)
	if partition == nil {
		response.Error(c, errors.New("partition not found"))
		return
	}

	partitionEx, err := h.Params.PassportExService.GetPartitionEx(ctx, partition.ID, partition.ChapterID)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewPartitionExView(*partitionEx)
	response.Response(c, v)
}
