package mun

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (m *Middleware) WithMunicipality() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		municipalityID, err := parser.ParseMunicipalityID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		municipality, err := m.Params.MunicipalityService.GetById(ctx, municipalityID)
		if err != nil {
			response.Error(c, err)
		}

		ctx = context_paylod_parser.SetMunicipalityToContext(ctx, municipality)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
