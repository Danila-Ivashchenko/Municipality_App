package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type RouteObjectRepository interface {
	Create(ctx context.Context, data *entity.RouteObject) (*entity.RouteObject, error)
	Update(ctx context.Context, data *entity.RouteObject) (*entity.RouteObject, error)

	GetByRouteID(ctx context.Context, routeID int64) ([]entity.RouteObject, error)

	GetID(ctx context.Context, id int64) (*entity.RouteObject, error)
	GetIDs(ctx context.Context, ids []int64) ([]entity.RouteObject, error)

	GetByNameAndRouteID(ctx context.Context, name string, routeID int64) (*entity.RouteObject, error)
	GetByNamesAndRouteID(ctx context.Context, names []string, routeID int64) ([]entity.RouteObject, error)

	Delete(ctx context.Context, id int64) error
	DeleteToRoute(ctx context.Context, routeId int64) error
}
