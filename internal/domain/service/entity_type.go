package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityTypeService interface {
	Create(ctx context.Context, data *CreateEntityTypeData) (*entity.EntityType, error)
	Update(ctx context.Context, data *UpdateEntityTypeData) (*entity.EntityType, error)

	CreateMultiply(ctx context.Context, data *CreateEntityTypeMultiplyData) ([]entity.EntityType, error)

	Delete(ctx context.Context, ids []int64) error

	GetAll(ctx context.Context) ([]entity.EntityType, error)

	GetByName(ctx context.Context, name string) (*entity.EntityType, error)
	GetByNames(ctx context.Context, names []string) ([]entity.EntityType, error)

	GetByID(ctx context.Context, id int64) (*entity.EntityType, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityType, error)
}

type CreateEntityTypeData struct {
	Name string
}

type CreateEntityTypeMultiplyData struct {
	Data []CreateEntityTypeData
}

type UpdateEntityTypeData struct {
	ID   int64
	Name string
}

type UpdateEntityTypeMultiplyData struct {
	Data []UpdateEntityTypeData
}
