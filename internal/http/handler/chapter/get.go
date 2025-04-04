package chapter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) GetChapter(c *gin.Context) {
	ctx := parser.Context(c)

	passport := context_paylod_parser.GetPassportFromContext(ctx)
	if passport == nil {
		response.Error(c, errors.New("passport not found"))
		return
	}

	chapter := context_paylod_parser.GetChapterFromContext(ctx)
	if chapter == nil {
		response.Error(c, errors.New("passport not found"))
		return
	}

	chapterEx, err := h.Params.PassportExService.GetChapterEx(ctx, chapter.ID, passport.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewChapterExView(*chapterEx)
	response.Response(c, v)
}
