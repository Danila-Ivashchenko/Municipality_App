package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityAttributeValueRepository interface {
	Create(ctx context.Context, obj *entity.EntityAttributeValue) (*entity.EntityAttributeValue, error)
	Update(ctx context.Context, obj *entity.EntityAttributeValue) (*entity.EntityAttributeValue, error)

	Delete(ctx context.Context, id int64) error
	GetByAttributeID(ctx context.Context, attributeID int64) ([]entity.EntityAttributeValue, error)
	GetByAttributeIDAndEntityID(ctx context.Context, entityID, attributeID int64) (*entity.EntityAttributeValue, error)
	GetByEntityID(ctx context.Context, entityID int64) ([]entity.EntityAttributeValue, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityAttributeValue, error)
}
