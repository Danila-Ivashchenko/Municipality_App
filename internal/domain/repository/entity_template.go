package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityTemplateRepository interface {
	Create(ctx context.Context, data *CreateEntityTemplateData) (*entity.EntityTemplate, error)
	Update(ctx context.Context, data *entity.EntityTemplate) (*entity.EntityTemplate, error)

	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.EntityTemplate, error)
	GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.EntityTemplate, error)
	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.EntityTemplate, error)

	GetByID(ctx context.Context, id int64) (*entity.EntityTemplate, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityTemplate, error)

	Delete(ctx context.Context, id int64) error
}

type CreateEntityTemplateData struct {
	Name           string
	MunicipalityID int64
	EntityType     int64
}
