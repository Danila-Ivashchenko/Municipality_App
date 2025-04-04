package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectTypeService interface {
	Create(ctx context.Context, data *CreateObjectTypeData) (*entity.ObjectType, error)
	Update(ctx context.Context, data *UpdateObjectTypeData) (*entity.ObjectType, error)

	CreateMultiply(ctx context.Context, data *CreateObjectTypeMultiplyData) ([]entity.ObjectType, error)

	Delete(ctx context.Context, ids []int64) error

	GetAll(ctx context.Context) ([]entity.ObjectType, error)

	GetByName(ctx context.Context, name string) (*entity.ObjectType, error)
	GetByNames(ctx context.Context, names []string) ([]entity.ObjectType, error)

	GetByID(ctx context.Context, id int64) (*entity.ObjectType, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectType, error)
}

type CreateObjectTypeData struct {
	Name string
}

type CreateObjectTypeMultiplyData struct {
	Data []CreateObjectTypeData
}

type UpdateObjectTypeData struct {
	ID   int64
	Name string
}

type UpdateObjectTypeMultiplyData struct {
	Data []UpdateObjectTypeData
}
