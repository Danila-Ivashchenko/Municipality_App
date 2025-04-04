package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectTemplateService interface {
	Create(ctx context.Context, data *CreateObjectTemplateData) (*entity.ObjectTemplateEx, error)
	Update(ctx context.Context, data *UpdateObjectTemplateData) (*entity.ObjectTemplateEx, error)

	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.ObjectTemplate, error)
	GetExByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.ObjectTemplateEx, error)
	GetExByID(ctx context.Context, id int64) (*entity.ObjectTemplateEx, error)
	GetByID(ctx context.Context, id int64) (*entity.ObjectTemplate, error)
	GetExByIDs(ctx context.Context, ids []int64) ([]entity.ObjectTemplateEx, error)
	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.ObjectTemplate, error)

	GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.ObjectTemplate, error)

	DeleteByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) error
	Delete(ctx context.Context, id int64) error
}

type CreateObjectTemplateData struct {
	Name           string
	MunicipalityID int64
	ObjectType     int64
	Attributes     []CreateObjectAttributeToTemplateData
}

type CreateObjectAttributeToTemplateData struct {
	Name         string
	DefaultValue string
	ToShow       bool
}

type UpdateObjectTemplateData struct {
	ID                 int64
	Name               *string
	MunicipalityID     int64
	ObjectType         *int64
	AttributesToCreate []CreateObjectAttributeToTemplateData
	AttributesToUpdate []UpdateObjectAttributeToTemplateData
	AttributesToDelete []int64
}

type UpdateObjectAttributeToTemplateData struct {
	ID           int64
	Name         *string
	DefaultValue *string
	ToShow       *bool
}
