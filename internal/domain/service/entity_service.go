package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityService interface {
	Create(ctx context.Context, data *CreateEntityData) (*entity.Entity, error)
	CreateMultiply(ctx context.Context, data *CreateMultiplyEntitiesData) ([]entity.EntityEx, error)

	UpdateMultiply(ctx context.Context, data *UpdateMultiplyEntitiesData) ([]entity.EntityEx, error)

	GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Entity, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Entity, error)

	GetByID(ctx context.Context, ids int64) (*entity.Entity, error)
	GetByNamesAndTemplateID(ctx context.Context, names []string, templateID int64) ([]entity.Entity, error)

	GetExByIDs(ctx context.Context, ids []int64) ([]entity.EntityEx, error)
	GetExByTemplateID(ctx context.Context, templateID int64) ([]entity.EntityEx, error)

	DeleteMultiple(ctx context.Context, ids []int64, templateID int64) error
}

type EntityRepository interface {
	GetByTemplateID(ctx context.Context, partitionID int64) ([]entity.Entity, error)
	GetByTemplateIDAndNames(ctx context.Context, name []string, partitionID int64) ([]entity.Entity, error)

	GetByTemplateIDAndName(ctx context.Context, name string, partitionID int64) (*entity.Entity, error)
	GetByID(ctx context.Context, id int64) (*entity.Entity, error)
	Delete(ctx context.Context, id int64) error
}

type CreateEntityData struct {
	Name            string
	Description     string
	AttributeValues []CreateEntityAttributeValueData
}

type CreateMultiplyEntitiesData struct {
	EntityTemplateID int64
	Entities         []CreateEntityData
}

type UpdateMultiplyEntitiesData struct {
	EntityTemplateID int64
	Entities         []UpdateEntityData
}

type UpdateEntityData struct {
	ID              int64
	Name            *string
	Description     *string
	AttributeValues []CreateEntityAttributeValueData
}
