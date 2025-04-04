package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectAttributeValueRepository interface {
	Create(ctx context.Context, obj *entity.ObjectAttributeValue) (*entity.ObjectAttributeValue, error)
	Update(ctx context.Context, obj *entity.ObjectAttributeValue) (*entity.ObjectAttributeValue, error)

	Delete(ctx context.Context, id int64) error
	GetByAttributeID(ctx context.Context, attributeID int64) ([]entity.ObjectAttributeValue, error)
	GetByAttributeIDAndObjectID(ctx context.Context, objectID, attributeID int64) (*entity.ObjectAttributeValue, error)
	GetByObjectID(ctx context.Context, object int64) ([]entity.ObjectAttributeValue, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectAttributeValue, error)
}
