package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectAttributeService interface {
	CreateAttribute(ctx context.Context, data CreateObjectAttributeData) (*entity.ObjectAttribute, error)
	UpdateAttribute(ctx context.Context, data UpdateObjectAttributeData) (*entity.ObjectAttribute, error)
	DeleteAttribute(ctx context.Context, id int64) error

	CreateValues(ctx context.Context, data CreateObjectAttributesData) ([]entity.ObjectAttributeValue, error)
	UpdateValues(ctx context.Context, data CreateObjectAttributesData) ([]entity.ObjectAttributeValue, error)

	GetAttributesByObjectTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectAttribute, error)
	GetAttributeByIDAndTemplateID(ctx context.Context, attributeID, templateID int64) (*entity.ObjectAttribute, error)
	GetValuesByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectAttributeValue, error)
	GetAttributesExByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectAttributeValueEx, error)
}

type CreateObjectAttributeData struct {
	ObjectTemplateID int64
	Name             string
	DefaultValue     string
	ToShow           bool
}

type UpdateObjectAttributeData struct {
	ID               int64
	ObjectTemplateID int64
	Name             *string
	DefaultValue     *string
	ToShow           *bool
}

type CreateObjectAttributesData struct {
	ObjectID         int64
	ObjectTemplateID int64
	ValuesData       []CreateObjectAttributeValueData
}

type CreateObjectAttributeValueData struct {
	AttributeID int64
	Value       *string
}

type UpdateObjectAttributesData struct {
	ObjectID         int64
	ObjectTemplateID int64
	ValuesData       []UpdateObjectAttributeValueData
}

type UpdateObjectAttributeValueData struct {
	AttributeID int64
	Value       *string
}
