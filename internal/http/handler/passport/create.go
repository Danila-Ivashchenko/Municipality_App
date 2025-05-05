package passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) Create(c *gin.Context) {
	req := &reqCreatePassport{}

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

	data := req.Convert(municipality.ID)

	passport, err := h.Params.PassportService.Create(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newPassportView(passport)
	response.Response(c, view)
}

func (h *Handler) CreateFile(c *gin.Context) {
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

	passportEx, err := h.Params.PassportExService.GetByIDAndMunicipalityID(ctx, passport.ID, passport.MunicipalityID)
	if err != nil {
		response.Error(c, err)
		return
	}

	passportFile, err := h.Params.PassportFileService.Create(ctx, municipality, passportEx)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewPassportFileView(passportFile)
	response.Response(c, v)
}
