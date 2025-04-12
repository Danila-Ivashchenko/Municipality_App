package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityTemplateService interface {
	Create(ctx context.Context, data *CreateEntityTemplateData) (*entity.EntityTemplateEx, error)
	Update(ctx context.Context, data *UpdateEntityTemplateData) (*entity.EntityTemplateEx, error)

	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.EntityTemplate, error)
	GetExByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.EntityTemplateEx, error)
	GetExByID(ctx context.Context, id int64) (*entity.EntityTemplateEx, error)
	GetByID(ctx context.Context, id int64) (*entity.EntityTemplate, error)
	GetExByIDs(ctx context.Context, ids []int64) ([]entity.EntityTemplateEx, error)
	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.EntityTemplate, error)

	GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.EntityTemplate, error)

	DeleteByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) error
	Delete(ctx context.Context, id int64) error
}

type CreateEntityTemplateData struct {
	Name           string
	MunicipalityID int64
	EntityType     int64
	Attributes     []CreateEntityAttributeToTemplateData
}

type CreateEntityAttributeToTemplateData struct {
	Name         string
	DefaultValue string
	ToShow       bool
}

type UpdateEntityTemplateData struct {
	ID                 int64
	Name               *string
	MunicipalityID     int64
	EntityType         *int64
	AttributesToCreate []CreateEntityAttributeToTemplateData
	AttributesToUpdate []UpdateEntityAttributeToTemplateData
	AttributesToDelete []int64
}

type UpdateEntityAttributeToTemplateData struct {
	ID           int64
	Name         *string
	DefaultValue *string
	ToShow       *bool
}
