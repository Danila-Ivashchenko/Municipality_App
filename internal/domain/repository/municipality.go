package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type MunicipalityRepository interface {
	Create(ctx context.Context, data *CreateMunicipalityData) (*entity.Municipality, error)
	Update(ctx context.Context, municipality *entity.Municipality) error
	Delete(ctx context.Context, id int64) error

	GetById(ctx context.Context, id int64) (*entity.Municipality, error)
	GetByName(ctx context.Context, name string) (*entity.Municipality, error)
	GetAll(ctx context.Context) ([]entity.Municipality, error)

	GetByParams(ctx context.Context, params *MunicipalityParams) ([]entity.Municipality, error)
}

type CreateMunicipalityData struct {
	Name     string
	RegionID int64
}

type Pagination struct {
	Limit  int
	Offset int
}

type MunicipalityParams struct {
	ID       *int64
	Name     *string
	RegionID *int64
	IsHidden *bool
}
