package passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetMainByMunicipality(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality not found"))
		return
	}

	passport, err := h.Params.PassportService.GetMainByMunicipalityID(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportView(passport)
	response.Response(c, view)
}

func (h *Handler) GetMunicipality(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality not found"))
		return
	}

	passports, err := h.Params.PassportService.GetByMunicipalityID(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportViews(passports)
	response.Response(c, view)
}

func (h *Handler) GetByRevisionCode(c *gin.Context) {
	ctx := parser.Context(c)

	code := c.Query(revisionCodeKey)

	if code == "" {
		response.Error(c, errors.New("code not provided"))
	}

	passport, err := h.Params.PassportService.GetByRevisionCode(ctx, code)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportView(passport)
	response.Response(c, view)
}
