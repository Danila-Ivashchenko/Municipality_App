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
	rout.bindUserAdmin()
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

func (rout *Router) bindUserAdmin() {
	userRouter := rout.r.Group("/admin/users").Use(rout.AuthMiddleware.WithAuth())

	userRouter.GET("", rout.UserAdminHandler.GetAll)
	userRouter.PUT("", rout.UserAdminHandler.Update)
}

func (rout *Router) bindRegion() {
	regionRouter := rout.r.Group("/region")

	regionRouter.POST("/params", rout.RegionHandler.GetByParams)

	regionRouterWithAuth := regionRouter.
		Use(rout.AuthMiddleware.WithAuth()).
		Use(rout.AuthMiddleware.WithAdmin())

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
	passportRouterWithMunicipalityWithAuth.Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	passportRouterWithMunicipalityWithAuth.POST("", rout.PassportHandler.Create)

	passportRouterWithMunicipality.GET("", rout.PassportHandler.GetMunicipality)
	passportRouterWithMunicipality.GET("/main", rout.PassportHandler.GetMainByMunicipality)

	passportRouterWithID := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	passportRouterWithWithAuthIDWithDelete := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanDelete())

	passportRouterWithWithAuthIDWithWrite := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	passportRouterWithID.GET("", rout.PassportHandler.GetMunicipalityAndID)
	passportRouterWithWithAuthIDWithWrite.POST("/file", rout.PassportHandler.CreateFile)
	passportRouterWithWithAuthIDWithWrite.PUT("", rout.PassportHandler.Update)
	passportRouterWithWithAuthIDWithDelete.DELETE("", rout.PassportHandler.Delete)
	passportRouterWithWithAuthIDWithWrite.POST("", rout.PassportHandler.Copy)
}

func (rout *Router) bindObjectTypes() {
	objectTypeRouter := rout.r.Group("/object_type")
	objectTypeRouterWithWrite := rout.r.Group("/object_type").Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanDelete())
	objectTypeRouterWithDelete := rout.r.Group("/object_type").Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	objectTypeRouter.GET("", rout.ObjectTypeHandler.GetAll)
	objectTypeRouterWithWrite.POST("", rout.ObjectTypeHandler.Create)
	objectTypeRouterWithWrite.PUT("", rout.ObjectTypeHandler.Update)
	objectTypeRouterWithDelete.DELETE("", rout.ObjectTypeHandler.Delete)
}

func (rout *Router) bindChapter() {

	chapterRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	chapterRouterWithUpdated := chapterRouter.
		Use(updatedPassport()).Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	chapterRouterWithUpdated.POST("", rout.ChapterHandler.CreateChapter)

	chapterWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	chapterWithIDRouterWithUpdated := chapterWithIDRouter.
		Use(rout.AuthMiddleware.WithAuth()).
		Use(updatedPassport())

	chapterWithIDRouterWithUpdatedWithWrite := chapterWithIDRouterWithUpdated.Use(rout.AuthMiddleware.WithCanEdit())
	chapterWithIDRouterWithUpdatedWithDelete := chapterWithIDRouterWithUpdated.Use(rout.AuthMiddleware.WithCanDelete())

	chapterWithIDRouterWithUpdatedWithWrite.PUT("", rout.ChapterHandler.UpdateChapter).Use(updatedPassport())
	chapterWithIDRouterWithUpdatedWithDelete.DELETE("", rout.ChapterHandler.Delete).Use(updatedPassport())
	chapterWithIDRouter.GET("", rout.ChapterHandler.GetChapter)
}

func (rout *Router) bindPartition() {
	partitionRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	partitionRouterWithUpdated := partitionRouter.Use(updatedPassport()).Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	partitionRouterWithUpdated.POST("", rout.PartitionHandler.CreatePartition)

	partitionWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition())

	partitionWithIDRouterWithUpdated := partitionWithIDRouter.
		Use(rout.AuthMiddleware.WithAuth()).
		Use(updatedPassport())

	partitionWithIDRouterWithUpdatedWithWrite := partitionWithIDRouterWithUpdated.Use(rout.AuthMiddleware.WithCanEdit())
	partitionWithIDRouterWithUpdatedWithDelete := partitionWithIDRouterWithUpdated.Use(rout.AuthMiddleware.WithCanDelete())

	partitionWithIDRouterWithUpdatedWithWrite.PUT("", rout.PartitionHandler.UpdatePartition).Use(updatedPassport())
	partitionWithIDRouterWithUpdatedWithDelete.DELETE("", rout.PartitionHandler.DeletePartition).Use(updatedPassport())

	partitionWithIDRouter.GET("", rout.PartitionHandler.GetPartition)

}

func (rout *Router) bindObject() {
	objectRouter := rout.r.Group("/municipality/:municipality_id/object_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	objectRouterWithAuth := rout.r.Group("/municipality/:municipality_id/object_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.AuthMiddleware.WithAuth())

	objectRouterWithAuthWithWrite := objectRouterWithAuth.Use(rout.AuthMiddleware.WithCanEdit())

	objectRouterWithAuthWithWrite.POST("", rout.ObjectHandler.CreateTemplate)
	objectRouter.GET("", rout.ObjectHandler.GetTemplatesByMunicipality)

	objectRouterWithTemplateID := rout.r.Group("/municipality/:municipality_id/object_template/:object_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithObjectTemplate())

	objectRouterWithTemplateIDWithAuth := rout.r.Group("/municipality/:municipality_id/object_template/:object_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithObjectTemplate()).
		Use(rout.AuthMiddleware.WithAuth())

	objectRouterWithTemplateIDWithAuthWithWrite := objectRouterWithTemplateIDWithAuth.Use(rout.AuthMiddleware.WithCanEdit())
	objectRouterWithTemplateIDWithAuthWithDelete := objectRouterWithTemplateIDWithAuth.Use(rout.AuthMiddleware.WithCanDelete())

	objectRouterWithTemplateID.GET("", rout.ObjectHandler.GetTemplateByID)
	objectRouterWithTemplateIDWithAuthWithWrite.PUT("", rout.ObjectHandler.UpdateTemplate)
	objectRouterWithTemplateIDWithAuthWithDelete.DELETE("", rout.ObjectHandler.DeleteTemplate)

	objectRouterWithTemplateIDWithAuthWithWrite.POST("/objects", rout.ObjectHandler.CreateObjects)
	objectRouterWithTemplateID.GET("/objects", rout.ObjectHandler.GetObjects)
	objectRouterWithTemplateIDWithAuthWithWrite.PUT("/objects", rout.ObjectHandler.UpdateObjects)
	objectRouterWithTemplateIDWithAuthWithDelete.DELETE("/objects", rout.ObjectHandler.DeleteObjects)
}

func (rout *Router) bindEntityTypes() {
	objectTypeRouter := rout.r.Group("/entity_type")
	objectTypeRouterWithWrite := rout.r.Group("/entity_type").Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	objectTypeRouter.GET("", rout.EntityTypeHandler.GetAll)
	objectTypeRouterWithWrite.POST("", rout.EntityTypeHandler.Create)
	objectTypeRouterWithWrite.PUT("", rout.EntityTypeHandler.Update)
}

func (rout *Router) bindEntities() {
	entityRouter := rout.r.Group("/municipality/:municipality_id/entity_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	entityRouter.POST("", rout.EntityHandler.CreateTemplate)
	entityRouter.GET("", rout.EntityHandler.GetTemplatesByMunicipality)

	entityRouterWithTemplateID := rout.r.Group("/municipality/:municipality_id/entity_template/:entity_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithEntityTemplate())

	entityRouterWithTemplateIDWithAuth := entityRouterWithTemplateID.Use(rout.AuthMiddleware.WithAuth())

	entityRouterWithTemplateIDWithAuthWithWrite := entityRouterWithTemplateIDWithAuth.Use(rout.AuthMiddleware.WithCanEdit())
	entityRouterWithTemplateIDWithAuthWithDelete := entityRouterWithTemplateIDWithAuth.Use(rout.AuthMiddleware.WithCanDelete())

	entityRouterWithTemplateID.GET("", rout.EntityHandler.GetTemplateByID)
	entityRouterWithTemplateIDWithAuthWithWrite.PUT("", rout.EntityHandler.UpdateTemplate)
	entityRouterWithTemplateIDWithAuthWithDelete.DELETE("", rout.EntityHandler.DeleteTemplate)

	entityRouterWithTemplateIDWithAuthWithWrite.POST("/entities", rout.EntityHandler.CreateEntities)
	entityRouterWithTemplateID.GET("/entities", rout.EntityHandler.GetEntitys)
	entityRouterWithTemplateIDWithAuthWithWrite.PUT("/entities", rout.EntityHandler.UpdateEntities)
	entityRouterWithTemplateIDWithAuthWithDelete.DELETE("/entities", rout.EntityHandler.DeleteEntities)
}

func (rout *Router) bindRoutes() {
	routeRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id/route").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition())

	updatedPassport := rout.PassportMiddleware.UpdatePassportUpdatedAt

	routeRouterWithUpdated := routeRouter.Use(updatedPassport()).Use(rout.AuthMiddleware.WithAuth()).Use(rout.AuthMiddleware.WithCanEdit())

	routeRouterWithUpdated.POST("", rout.RouteHandler.Create).Use(updatedPassport())

	routeRouterWithID := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id/route/:route_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition()).
		Use(rout.PassportMiddleware.WithRoute())

	routeRouterWithIDWithUpdated := routeRouterWithID.Use(rout.AuthMiddleware.WithAuth()).Use(updatedPassport())
	routeRouterWithIDWithUpdatedWithWrite := routeRouterWithIDWithUpdated.Use(rout.AuthMiddleware.WithCanEdit())
	routeRouterWithIDWithUpdatedWithDelete := routeRouterWithIDWithUpdated.Use(rout.AuthMiddleware.WithCanDelete())

	routeRouterWithIDWithUpdatedWithWrite.PUT("", rout.RouteHandler.Update).Use(updatedPassport())
	routeRouterWithIDWithUpdatedWithDelete.DELETE("", rout.RouteHandler.Delete).Use(updatedPassport())
}
