package passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Update(c *gin.Context) {
	req := &reqUpdatePassport{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

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

	data := req.Convert(passport.ID, municipality.ID)

	passportUpdated, err := h.Params.PassportService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportView(passportUpdated)
	response.Response(c, view)
}
