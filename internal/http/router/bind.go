package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/keys"
	"time"
)

func (rout *Router) bind() *gin.Engine {
	r := rout.r

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	rout.bindUser()
	rout.bindRegion()
	rout.bindMunicipality()
	rout.bindPassport()
	rout.bindObjectTypes()
	rout.bindChapter()
	rout.bindPartition()
	rout.bindObject()
	rout.bindEntityTypes()
	rout.bindEntities()
	rout.bindRoutes()

	return r
}

func (rout *Router) bindUser() {
	userRouter := rout.r.Group("/user")

	userRouter.POST("/register", rout.UserHandler.Register)
	userRouter.POST("/login", rout.UserHandler.Login)

	userRouterWithAuth := userRouter.Use(rout.AuthMiddleware.WithAuth())
	userRouterWithAuth.GET("/me", rout.UserHandler.Me)
	userRouterWithAuth.POST("/change_password", rout.UserHandler.ChangePassword)
	userRouterWithAuth.PUT("", rout.UserHandler.Update)
	userRouterWithAuth.POST("/logout", rout.UserHandler.Logout)
}

func (rout *Router) bindRegion() {
	regionRouter := rout.r.Group("/region")

	regionRouter.POST("/params", rout.RegionHandler.GetByParams)

	regionRouterWithAuth := regionRouter.Use(rout.AuthMiddleware.WithAuth())

	regionRouterWithAuth.POST("", rout.RegionHandler.Create)
}

func (rout *Router) bindMunicipality() {
	municipalityRouter := rout.r.Group("/municipality")
	municipalityRouterWithAuth := rout.r.Group("/municipality")

	municipalityRouterWithAuth.
		Use(rout.AuthMiddleware.WithAuth()).
		Use(rout.AuthMiddleware.WithAdmin())

	municipalityRouter.POST("/params", rout.MunicipalityHandler.GetByParams)

	municipalityRouterWithAuth.POST("", rout.MunicipalityHandler.Create)

	municipalityRouterWithID := municipalityRouter.Group(keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey)).Use(rout.MunicipalityMiddleware.WithMunicipality())
	municipalityRouterWithAuthWithID := municipalityRouterWithAuth.Group(keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey)).Use(rout.MunicipalityMiddleware.WithMunicipality())

	municipalityRouterWithID.GET("", rout.MunicipalityHandler.GetByID)
	municipalityRouterWithAuthWithID.Use().PUT("", rout.MunicipalityHandler.Update)
	municipalityRouterWithAuthWithID.Use().DELETE("", rout.MunicipalityHandler.Delete)
}

func (rout *Router) bindPassport() {
	passportRouter := rout.r.Group("/passport")

	passportRouter.GET("/revision_code", rout.PassportHandler.GetByRevisionCode)

	passportRouterWithMunicipality := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	passportRouterWithMunicipalityWithAuth := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").
		Use(rout.MunicipalityMiddleware.WithMunicipality())
	passportRouterWithMunicipalityWithAuth.Use(rout.AuthMiddleware.WithAuth(), rout.AuthMiddleware.WithAdmin())

	passportRouterWithMunicipalityWithAuth.POST("", rout.PassportHandler.Create)

	passportRouterWithMunicipality.GET("", rout.PassportHandler.GetMunicipality)
	passportRouterWithMunicipality.GET("/main", rout.PassportHandler.GetMainByMunicipality)

	passportRouterWithID := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	passportRouterWithWithAuthID := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.AuthMiddleware.WithAuth()).
		Use(rout.AuthMiddleware.WithAdmin())

	passportRouterWithID.GET("", rout.PassportHandler.GetMunicipalityAndID)
	passportRouterWithWithAuthID.POST("/file", rout.PassportHandler.CreateFile)
	passportRouterWithWithAuthID.PUT("", rout.PassportHandler.Update)
	passportRouterWithWithAuthID.DELETE("", rout.PassportHandler.Delete)
}

func (rout *Router) bindObjectTypes() {
	objectTypeRouter := rout.r.Group("/object_type")

	objectTypeRouter.GET("", rout.ObjectTypeHandler.GetAll)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth(), rout.AuthMiddleware.WithAdmin()).POST("", rout.ObjectTypeHandler.Create)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth(), rout.AuthMiddleware.WithAdmin()).PUT("", rout.ObjectTypeHandler.Update)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth(), rout.AuthMiddleware.WithAdmin()).DELETE("", rout.ObjectTypeHandler.Delete)
}

func (rout *Router) bindChapter() {

	chapterRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	chapterRouterWithUpdated := chapterRouter.
		Use(updatedPassport())

	chapterRouterWithUpdated.POST("", rout.ChapterHandler.CreateChapter)

	chapterWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	chapterWithIDRouterWithUpdated := chapterWithIDRouter.
		Use(updatedPassport())

	chapterWithIDRouterWithUpdated.PUT("", rout.ChapterHandler.UpdateChapter).Use(updatedPassport())
	chapterWithIDRouterWithUpdated.DELETE("", rout.ChapterHandler.Delete).Use(updatedPassport())
	chapterWithIDRouter.GET("", rout.ChapterHandler.GetChapter)
}

func (rout *Router) bindPartition() {
	partitionRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	partitionRouterWithUpdated := partitionRouter.Use(updatedPassport())

	partitionRouterWithUpdated.POST("", rout.PartitionHandler.CreatePartition).Use(updatedPassport())

	partitionWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition())

	partitionWithIDRouterWithUpdated := partitionWithIDRouter.Use(updatedPassport())

	partitionWithIDRouterWithUpdated.PUT("", rout.PartitionHandler.UpdatePartition).Use(updatedPassport())
	partitionWithIDRouterWithUpdated.DELETE("", rout.PartitionHandler.DeletePartition).Use(updatedPassport())

	partitionWithIDRouter.GET("", rout.PartitionHandler.GetPartition)

}

func (rout *Router) bindObject() {
	objectRouter := rout.r.Group("/municipality/:municipality_id/object_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	objectRouterWithAuth := rout.r.Group("/municipality/:municipality_id/object_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.AuthMiddleware.WithAuth()).
		Use(rout.AuthMiddleware.WithAdmin())

	objectRouterWithAuth.POST("", rout.ObjectHandler.CreateTemplate)
	objectRouter.GET("", rout.ObjectHandler.GetTemplatesByMunicipality)

	objectRouterWithTemplateID := rout.r.Group("/municipality/:municipality_id/object_template/:object_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithObjectTemplate())

	objectRouterWithTemplateIDWithAuth := rout.r.Group("/municipality/:municipality_id/object_template/:object_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithObjectTemplate()).
		Use(rout.AuthMiddleware.WithAuth()).
		Use(rout.AuthMiddleware.WithAdmin())

	objectRouterWithTemplateID.GET("", rout.ObjectHandler.GetTemplateByID)
	objectRouterWithTemplateIDWithAuth.PUT("", rout.ObjectHandler.UpdateTemplate)
	objectRouterWithTemplateIDWithAuth.DELETE("", rout.ObjectHandler.DeleteTemplate)

	objectRouterWithTemplateIDWithAuth.POST("/objects", rout.ObjectHandler.CreateObjects)
	objectRouterWithTemplateID.GET("/objects", rout.ObjectHandler.GetObjects)
	objectRouterWithTemplateIDWithAuth.PUT("/objects", rout.ObjectHandler.UpdateObjects)
	objectRouterWithTemplateIDWithAuth.DELETE("/objects", rout.ObjectHandler.DeleteObjects)
}

func (rout *Router) bindEntityTypes() {
	objectTypeRouter := rout.r.Group("/entity_type")

	objectTypeRouter.GET("", rout.EntityTypeHandler.GetAll)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth()).POST("", rout.EntityTypeHandler.Create)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth()).PUT("", rout.EntityTypeHandler.Update)
}

func (rout *Router) bindEntities() {
	entityRouter := rout.r.Group("/municipality/:municipality_id/entity_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	entityRouter.POST("", rout.EntityHandler.CreateTemplate)
	entityRouter.GET("", rout.EntityHandler.GetTemplatesByMunicipality)

	entityRouterWithTemplateID := rout.r.Group("/municipality/:municipality_id/entity_template/:entity_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithEntityTemplate())

	entityRouterWithTemplateID.GET("", rout.EntityHandler.GetTemplateByID)
	entityRouterWithTemplateID.PUT("", rout.EntityHandler.UpdateTemplate)
	entityRouterWithTemplateID.DELETE("", rout.EntityHandler.DeleteTemplate)

	entityRouterWithTemplateID.POST("/entities", rout.EntityHandler.CreateEntities)
	entityRouterWithTemplateID.GET("/entities", rout.EntityHandler.GetEntitys)
	entityRouterWithTemplateID.PUT("/entities", rout.EntityHandler.UpdateEntities)
	entityRouterWithTemplateID.DELETE("/entities", rout.EntityHandler.DeleteEntities)
}

func (rout *Router) bindRoutes() {
	routeRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id/route").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	routeRouterWithUpdated := routeRouter.Use(updatedPassport())

	routeRouterWithUpdated.POST("", rout.RouteHandler.Create).Use(updatedPassport())

	routeRouterWithID := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id/route/:route_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition()).
		Use(rout.PassportMiddleware.WithRoute())

	routeRouterWithIDWithUpdated := routeRouterWithID.Use(updatedPassport())

	routeRouterWithIDWithUpdated.PUT("", rout.RouteHandler.Update).Use(updatedPassport())
	routeRouterWithIDWithUpdated.DELETE("", rout.RouteHandler.Delete).Use(updatedPassport())
}
