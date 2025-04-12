package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityAttributeRepository interface {
	Create(ctx context.Context, obj *entity.EntityAttribute) (*entity.EntityAttribute, error)
	Update(ctx context.Context, obj *entity.EntityAttribute) (*entity.EntityAttribute, error)

	Delete(ctx context.Context, id int64) error
	GetByEntityTemplateID(ctx context.Context, templateID int64) ([]entity.EntityAttribute, error)
	GetByEntityTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.EntityAttribute, error)
	GetByEntityTemplateIDAndID(ctx context.Context, id, templateID int64) (*entity.EntityAttribute, error)
	GetByID(ctx context.Context, id int64) (*entity.EntityAttribute, error)
}
