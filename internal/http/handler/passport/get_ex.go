package passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetMunicipalityAndID(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality not found"))
		return
	}

	passport := context_paylod_parser.GetPassportFromContext(ctx)
	if passport == nil {
		response.Error(c, errors.New("passport not found"))
		return
	}

	passportEx, err := h.Params.PassportExService.GetByIDAndMunicipalityID(ctx, passport.ID, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportExView(passportEx)
	response.Response(c, view)
}
