package container

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/container"
	"municipality_app/internal/container/repository"
	"municipality_app/internal/container/service"
	"municipality_app/internal/http/container/handler"
	"municipality_app/internal/http/container/middleware"
	"municipality_app/internal/http/container/router"
	infrastructure "municipality_app/internal/infrastructure/container"
)

var (
	MainContainer = fx.Options(
		router.RouterContainer,
		infrastructure.InfrastructureContainer,
		repository.RepositoryContainer,
		service.ServiceContainer,
		handler.HandlerContainer,
		container.CommonContainer,
		middleware.MiddlewareContainer,
	)
)
