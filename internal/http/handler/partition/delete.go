package partition

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) DeletePartition(c *gin.Context) {
	ctx := parser.Context(c)

	chapter := context_paylod_parser.GetChapterFromContext(ctx)
	if chapter == nil {
		response.Error(c, errors.New("chapter not found"))
		return
	}

	partition := context_paylod_parser.GetPartitionFromContext(ctx)
	if partition == nil {
		response.Error(c, errors.New("partition not found"))
		return
	}

	err := h.Params.PartitionService.DeleteToChapter(ctx, []int64{partition.ID}, chapter.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
