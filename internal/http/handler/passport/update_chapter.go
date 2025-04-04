package passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) UpdateChapter(c *gin.Context) {
	req := &reqUpdateChapter{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	//if validateErr := req.Validate(); validateErr != nil {
	//	response.Error(c, validateErr)
	//	return
	//}

	passport := context_paylod_parser.GetPassportFromContext(ctx)
	if passport == nil {
		response.Error(c, errors.New("passport not found"))
		return
	}

	data := req.Convert(passport.ID, passport.MunicipalityID)

	chapterEx, err := h.Params.PassportExService.UpdateChapterEx(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewChapterExView(*chapterEx)
	response.Response(c, v)
}
