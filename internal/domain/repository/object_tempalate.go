package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectTemplateRepository interface {
	Create(ctx context.Context, data *CreateObjectTemplateData) (*entity.ObjectTemplate, error)
	Update(ctx context.Context, data *entity.ObjectTemplate) (*entity.ObjectTemplate, error)

	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.ObjectTemplate, error)
	GetByTypeID(ctx context.Context, typeID int64) ([]entity.ObjectTemplate, error)
	GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.ObjectTemplate, error)
	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.ObjectTemplate, error)

	GetByID(ctx context.Context, id int64) (*entity.ObjectTemplate, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectTemplate, error)

	Delete(ctx context.Context, id int64) error
}

type CreateObjectTemplateData struct {
	Name           string
	MunicipalityID int64
	ObjectType     int64
}
