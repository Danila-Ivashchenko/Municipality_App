package router

import (
	"github.com/gin-gonic/gin"
	"municipality_app/internal/http/keys"
)

func (rout *Router) bind() *gin.Engine {
	r := rout.r

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

	return r
}

func (rout *Router) bindUser() {
	userRouter := rout.r.Group("/user")

	userRouter.POST("/register", rout.UserHandler.Register)
	userRouter.POST("/login", rout.UserHandler.Login)

	userRouterWithAuth := userRouter.Use(rout.AuthMiddleware.WithAuth())
	userRouterWithAuth.GET("/me", rout.UserHandler.Me)
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

	municipalityRouter.Use(rout.AuthMiddleware.WithAuth()).POST("", rout.MunicipalityHandler.Create)
	municipalityRouter.POST("/params", rout.MunicipalityHandler.GetByParams)

	municipalityRouterWithID := municipalityRouter.Group(keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey)).Use(rout.MunicipalityMiddleware.WithMunicipality())

	municipalityRouterWithID.GET("", rout.MunicipalityHandler.GetByID)
	municipalityRouterWithID.Use(rout.AuthMiddleware.WithAuth()).PUT("", rout.MunicipalityHandler.Update)
}

func (rout *Router) bindPassport() {
	passportRouter := rout.r.Group("/passport")

	passportRouter.GET("/revision_code", rout.PassportHandler.GetByRevisionCode)
	passportRouterWithMunicipality := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.AuthMiddleware.WithAuth())

	passportRouterWithMunicipality.POST("", rout.PassportHandler.Create)

	passportRouterWithMunicipality.GET("", rout.PassportHandler.GetMunicipality).
		Use(rout.MunicipalityMiddleware.WithMunicipality())
	passportRouterWithMunicipality.GET("/main", rout.PassportHandler.GetMainByMunicipality).
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	passportRouterWithID := rout.r.Group("/municipality/" + keys.NewUriKeyPlaceHolder(keys.MunicipalityIdKey) + "/passport").Group(keys.NewUriKeyPlaceHolder(keys.PassportID)).
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	passportRouterWithID.GET("", rout.PassportHandler.GetMunicipalityAndID)
	passportRouterWithID.GET("/file", rout.PassportHandler.CreateFile)
	passportRouterWithID.PUT("", rout.PassportHandler.Update)
}

func (rout *Router) bindObjectTypes() {
	objectTypeRouter := rout.r.Group("/object_type")

	objectTypeRouter.GET("", rout.ObjectTypeHandler.GetAll)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth()).POST("", rout.ObjectTypeHandler.Create)
	objectTypeRouter.Use(rout.AuthMiddleware.WithAuth()).PUT("", rout.ObjectTypeHandler.Update)
}

func (rout *Router) bindChapter() {
	chapterRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport())

	chapterRouter.POST("", rout.ChapterHandler.CreateChapter)

	chapterWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	chapterWithIDRouter.PUT("", rout.ChapterHandler.UpdateChapter)
	chapterWithIDRouter.GET("", rout.ChapterHandler.GetChapter)
}

func (rout *Router) bindPartition() {
	partitionRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter())

	partitionRouter.POST("", rout.PartitionHandler.CreatePartition)

	partitionWithIDRouter := rout.r.Group("/municipality/:municipality_id/passport/:passport_id/chapter/:chapter_id/partition/:partition_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithPassport()).
		Use(rout.PassportMiddleware.WithChapter()).
		Use(rout.PassportMiddleware.WithPartition())

	partitionWithIDRouter.PUT("", rout.PartitionHandler.UpdatePartition)
	partitionWithIDRouter.GET("", rout.PartitionHandler.GetPartition)
}

func (rout *Router) bindObject() {
	objectRouter := rout.r.Group("/municipality/:municipality_id/object_template").
		Use(rout.MunicipalityMiddleware.WithMunicipality())

	objectRouter.POST("", rout.ObjectHandler.CreateTemplate)

	objectRouterWithTemplateID := rout.r.Group("/municipality/:municipality_id/object_template/:object_template_id").
		Use(rout.MunicipalityMiddleware.WithMunicipality()).
		Use(rout.PassportMiddleware.WithObjectTemplate())

	objectRouterWithTemplateID.GET("", rout.ObjectHandler.GetTemplateByID)
	objectRouterWithTemplateID.PUT("", rout.ObjectHandler.UpdateTemplate)
	objectRouterWithTemplateID.DELETE("", rout.ObjectHandler.DeleteTemplate)

	objectRouterWithTemplateID.POST("/objects", rout.ObjectHandler.CreateObjects)
	objectRouterWithTemplateID.GET("/objects", rout.ObjectHandler.GetObjects)
	objectRouterWithTemplateID.PUT("/objects", rout.ObjectHandler.UpdateObjects)
	objectRouterWithTemplateID.DELETE("/objects", rout.ObjectHandler.DeleteObjects)
}
