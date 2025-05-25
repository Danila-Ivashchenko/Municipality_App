package partition

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) CreatePartition(c *gin.Context) {
	req := &reqPartitionCreateData{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	chapter := context_paylod_parser.GetChapterFromContext(ctx)
	if chapter == nil {
		response.Error(c, errors.New("chapter not found"))
		return
	}

	data := req.Convert(chapter.ID)

	partition, err := h.Params.PartitionService.Create(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewPartitionExView(*partition)
	response.Response(c, v)
}
