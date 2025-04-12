package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityAttributeService interface {
	CreateAttribute(ctx context.Context, data CreateEntityAttributeData) (*entity.EntityAttribute, error)
	UpdateAttribute(ctx context.Context, data UpdateEntityAttributeData) (*entity.EntityAttribute, error)
	DeleteAttribute(ctx context.Context, id int64) error

	CreateValues(ctx context.Context, data CreateEntityAttributesData) ([]entity.EntityAttributeValue, error)
	UpdateValues(ctx context.Context, data CreateEntityAttributesData) ([]entity.EntityAttributeValue, error)

	GetAttributesByEntityTemplateID(ctx context.Context, templateID int64) ([]entity.EntityAttribute, error)
	GetAttributeByIDAndTemplateID(ctx context.Context, attributeID, templateID int64) (*entity.EntityAttribute, error)
	GetValuesByEntityID(ctx context.Context, objectID int64) ([]entity.EntityAttributeValue, error)
	GetAttributesExByEntityID(ctx context.Context, objectID int64) ([]entity.EntityAttributeValueEx, error)
}

type CreateEntityAttributeData struct {
	EntityTemplateID int64
	Name             string
	DefaultValue     string
	ToShow           bool
}

type UpdateEntityAttributeData struct {
	ID               int64
	EntityTemplateID int64
	Name             *string
	DefaultValue     *string
	ToShow           *bool
}

type CreateEntityAttributesData struct {
	EntityID         int64
	EntityTemplateID int64
	ValuesData       []CreateEntityAttributeValueData
}

type CreateEntityAttributeValueData struct {
	AttributeID int64
	Value       *string
}

type UpdateEntityAttributesData struct {
	EntityID         int64
	EntityTemplateID int64
	ValuesData       []UpdateEntityAttributeValueData
}

type UpdateEntityAttributeValueData struct {
	AttributeID int64
	Value       *string
}
