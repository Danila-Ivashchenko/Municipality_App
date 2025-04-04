package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectRepository interface {
	Create(ctx context.Context, data *CreateObjectData) (*entity.Object, error)
	Update(ctx context.Context, data *entity.Object) (*entity.Object, error)

	GetByTemplateID(ctx context.Context, partitionID int64) ([]entity.Object, error)
	GetByTemplateIDAndNames(ctx context.Context, name []string, partitionID int64) ([]entity.Object, error)

	GetByTemplateIDAndName(ctx context.Context, name string, partitionID int64) (*entity.Object, error)
	GetByIDsAndTemplateID(ctx context.Context, ids []int64, templateID int64) ([]entity.Object, error)

	GetByID(ctx context.Context, id int64) (*entity.Object, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Object, error)

	Delete(ctx context.Context, id int64) error
}

type CreateObjectData struct {
	Name             string
	ObjectTemplateID int64
	LocationID       *int64
	Description      string
}
