package municipality_passport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (m *Middleware) WithPassport() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
		if municipality == nil {
			response.Error(c, errors.New("municipality not found"))
			c.Abort()
			return
		}

		passportID, err := parser.ParsePassportID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		passport, err := m.Params.PassportService.GetByIDAndMunicipalityID(ctx, passportID, municipality.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if passport == nil {
			response.Error(c, errors.New("passport not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetPassportToContext(ctx, passport)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *Middleware) UpdatePassportUpdatedAt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		ctx := c.Request.Context()

		passport := context_paylod_parser.GetPassportFromContext(ctx)
		if passport == nil {
			response.Error(c, errors.New("passport not found"))
			c.Abort()
			return
		}

		err := m.Params.PassportService.UpdatedAt(ctx, passport.ID)
		if err != nil {
			response.Error(c, err)
		}
	}
}

func (m *Middleware) WithChapter() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		passport := context_paylod_parser.GetPassportFromContext(ctx)
		if passport == nil {
			response.Error(c, errors.New("passport not found"))
			c.Abort()
			return
		}

		chapterID, err := parser.ParseChapterID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		chapter, err := m.Params.ChapterService.GetByIDAndPassportID(ctx, chapterID, passport.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if chapter == nil {
			response.Error(c, errors.New("chapter not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetChapterToContext(ctx, chapter)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *Middleware) WithPartition() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		chapter := context_paylod_parser.GetChapterFromContext(ctx)
		if chapter == nil {
			response.Error(c, errors.New("chapter not found"))
			c.Abort()
			return
		}

		partitionID, err := parser.ParsePartitionID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		partition, err := m.Params.PartitionService.GetByIDAndChapterID(ctx, partitionID, chapter.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if partition == nil {
			response.Error(c, errors.New("partition not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetPartitionToContext(ctx, partition)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *Middleware) WithObjectTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
		if municipality == nil {
			response.Error(c, errors.New("municipality not found"))
			c.Abort()
			return
		}

		templateID, err := parser.ParseObjectTemplateID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		template, err := m.Params.ObjectTemplateService.GetByID(ctx, templateID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if template == nil {
			response.Error(c, errors.New("template not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetObjectTemplateToContext(ctx, template)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *Middleware) WithEntityTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
		if municipality == nil {
			response.Error(c, errors.New("municipality not found"))
			c.Abort()
			return
		}

		templateID, err := parser.ParseEntityTemplateID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		template, err := m.Params.EntityTemplateService.GetByID(ctx, templateID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if template == nil {
			response.Error(c, errors.New("template not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetEntityTemplateToContext(ctx, template)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *Middleware) WithRoute() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		partition := context_paylod_parser.GetPartitionFromContext(ctx)
		if partition == nil {
			response.Error(c, errors.New("partition not found"))
			c.Abort()
			return
		}

		routeID, err := parser.ParseRouteID(c)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		route, err := m.Params.RouteService.GetByIDAndPartitionID(ctx, routeID, partition.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		if route == nil {
			response.Error(c, errors.New("route not found"))
			c.Abort()
			return
		}

		ctx = context_paylod_parser.SetRouteToContext(ctx, route)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
