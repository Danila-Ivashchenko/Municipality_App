package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectService interface {
	Create(ctx context.Context, data *CreateObjectData) (*entity.Object, error)
	CreateMultiply(ctx context.Context, data *CreateMultiplyObjetsData) ([]entity.ObjectEx, error)

	UpdateMultiply(ctx context.Context, data *UpdateMultiplyObjetsData) ([]entity.ObjectEx, error)

	GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Object, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Object, error)

	GetByID(ctx context.Context, ids int64) (*entity.Object, error)
	GetByNamesAndTemplateID(ctx context.Context, names []string, templateID int64) ([]entity.Object, error)

	GetExByIDs(ctx context.Context, ids []int64) ([]entity.ObjectEx, error)
	GetExByTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectEx, error)

	DeleteMultiple(ctx context.Context, ids []int64, templateID int64) error
}

type ObjectRepository interface {
	GetByTemplateID(ctx context.Context, partitionID int64) ([]entity.Object, error)
	GetByTemplateIDAndNames(ctx context.Context, name []string, partitionID int64) ([]entity.Object, error)

	GetByTemplateIDAndName(ctx context.Context, name string, partitionID int64) (*entity.Object, error)
	GetByID(ctx context.Context, id int64) (*entity.Object, error)
	Delete(ctx context.Context, id int64) error
}

type CreateObjectData struct {
	Name            string
	LocationData    *CreateObjectLocationData
	Description     string
	AttributeValues []CreateObjectAttributeValueData
}

type CreateObjectLocationData struct {
	Address   *string
	Latitude  *float64
	Longitude *float64
	Geometry  *string
}

type CreateMultiplyObjetsData struct {
	ObjectTemplateID int64
	Objects          []CreateObjectData
}

type UpdateMultiplyObjetsData struct {
	ObjectTemplateID int64
	Objects          []UpdateObjectData
}

type UpdateObjectData struct {
	ID              int64
	Name            *string
	LocationData    *UpdateObjectLocationData
	Description     *string
	AttributeValues []CreateObjectAttributeValueData
}

type UpdateObjectLocationData struct {
	Address   *string
	Latitude  *float64
	Longitude *float64
	Geometry  *string
}
