package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type RouteRepository interface {
	Create(ctx context.Context, data *entity.Route) (*entity.Route, error)
	Update(ctx context.Context, data *entity.Route) (*entity.Route, error)

	GetByID(ctx context.Context, id int64) (*entity.Route, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Route, error)
	GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.Route, error)

	GetByNamePartitionID(ctx context.Context, name string, partitionID int64) (*entity.Route, error)
	GetByIDPartitionID(ctx context.Context, id, partitionID int64) (*entity.Route, error)

	Delete(ctx context.Context, id int64) error
}
