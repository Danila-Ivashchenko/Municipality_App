package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectAttributeRepository interface {
	Create(ctx context.Context, obj *entity.ObjectAttribute) (*entity.ObjectAttribute, error)
	Update(ctx context.Context, obj *entity.ObjectAttribute) (*entity.ObjectAttribute, error)

	Delete(ctx context.Context, id int64) error
	GetByObjectTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectAttribute, error)
	GetByObjectTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.ObjectAttribute, error)
	GetByObjectTemplateIDAndID(ctx context.Context, id, templateID int64) (*entity.ObjectAttribute, error)
	GetByID(ctx context.Context, id int64) (*entity.ObjectAttribute, error)
}
