package chapter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Delete(c *gin.Context) {
	ctx := parser.Context(c)

	passport := context_paylod_parser.GetPassportFromContext(ctx)
	if passport == nil {
		response.Error(c, errors.New("passport not found"))
		return
	}

	chapter := context_paylod_parser.GetChapterFromContext(ctx)
	if chapter == nil {
		response.Error(c, errors.New("chapter not found"))
		return
	}

	err := h.Params.PassportExService.DeleteChapterEx(ctx, passport.ID, chapter.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
