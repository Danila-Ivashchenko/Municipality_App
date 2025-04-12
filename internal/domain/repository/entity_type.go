package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityTypeRepository interface {
	Create(ctx context.Context, data *CreateEntityType) (*entity.EntityType, error)
	Update(ctx context.Context, data *entity.EntityType) error

	GetByName(ctx context.Context, name string) (*entity.EntityType, error)
	GetByNames(ctx context.Context, names []string) ([]entity.EntityType, error)

	GetByID(ctx context.Context, id int64) (*entity.EntityType, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityType, error)

	GetAll(ctx context.Context) ([]entity.EntityType, error)

	DeleteByID(ctx context.Context, id int64) error
}

type CreateEntityType struct {
	Name string
}

type UpdateEntityType struct {
	ID   int64
	Name *string
}
