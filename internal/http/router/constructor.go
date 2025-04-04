package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"municipality_app/internal/http/handler/chapter"
	"municipality_app/internal/http/handler/municipality"
	"municipality_app/internal/http/handler/object"
	"municipality_app/internal/http/handler/object_type"
	"municipality_app/internal/http/handler/partition"
	"municipality_app/internal/http/handler/passport"
	"municipality_app/internal/http/handler/region"
	"municipality_app/internal/http/handler/user"
	"municipality_app/internal/http/middleware/auth"
	"municipality_app/internal/http/middleware/mun"
	"municipality_app/internal/http/middleware/municipality_passport"
)

type RouterParams struct {
	fx.In

	AuthMiddleware         auth.Middleware
	MunicipalityMiddleware mun.Middleware
	PassportMiddleware     municipality_passport.Middleware

	UserHandler         user.Handler
	RegionHandler       region.Handler
	MunicipalityHandler municipality.Handler
	PassportHandler     passport.Handler
	ObjectTypeHandler   object_type.Handler
	ObjectHandler       object.Handler
	ChapterHandler      chapter.Handler
	PartitionHandler    partition.Handler
}

type Router struct {
	r *gin.Engine
	RouterParams
}

type RouterResult struct {
	fx.Out

	Router *gin.Engine `group:"routes"`
}

func New(params RouterParams) *gin.Engine {
	r := gin.Default()
	router := &Router{
		r:            r,
		RouterParams: params,
	}

	return router.bind()
}
