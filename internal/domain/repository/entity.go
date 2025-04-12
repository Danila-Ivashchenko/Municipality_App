package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityRepository interface {
	Create(ctx context.Context, data *CreateEntityData) (*entity.Entity, error)
	Update(ctx context.Context, data *entity.Entity) (*entity.Entity, error)

	GetByTemplateID(ctx context.Context, partitionID int64) ([]entity.Entity, error)
	GetByTemplateIDAndNames(ctx context.Context, name []string, partitionID int64) ([]entity.Entity, error)

	GetByTemplateIDAndName(ctx context.Context, name string, partitionID int64) (*entity.Entity, error)
	GetByIDsAndTemplateID(ctx context.Context, ids []int64, templateID int64) ([]entity.Entity, error)

	GetByID(ctx context.Context, id int64) (*entity.Entity, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Entity, error)

	Delete(ctx context.Context, id int64) error
}

type CreateEntityData struct {
	Name             string
	EntityTemplateID int64
	Description      string
}
