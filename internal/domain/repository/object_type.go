package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectTypeRepository interface {
	Create(ctx context.Context, data *CreateObjectType) (*entity.ObjectType, error)
	Update(ctx context.Context, data *entity.ObjectType) error

	GetByName(ctx context.Context, name string) (*entity.ObjectType, error)
	GetByNames(ctx context.Context, names []string) ([]entity.ObjectType, error)

	GetByID(ctx context.Context, id int64) (*entity.ObjectType, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectType, error)

	GetAll(ctx context.Context) ([]entity.ObjectType, error)

	DeleteByID(ctx context.Context, id int64) error
}

type CreateObjectType struct {
	Name string
}

type UpdateObjectType struct {
	ID   int64
	Name *string
}
