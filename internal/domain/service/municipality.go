package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type MunicipalityService interface {
	Create(ctx context.Context, data *CreateMunicipalityData) (*entity.MunicipalityEx, error)
	Update(ctx context.Context, data *UpdateMunicipalityData) (*entity.MunicipalityEx, error)
	Delete(ctx context.Context, id int64) error

	GetById(ctx context.Context, id int64) (*entity.Municipality, error)
	GetExById(ctx context.Context, id int64) (*entity.MunicipalityEx, error)

	GetByName(ctx context.Context, name string) (*entity.Municipality, error)
	GetAll(ctx context.Context) ([]entity.Municipality, error)

	GetByParams(ctx context.Context, params *GetMunicipalityParams) ([]entity.Municipality, error)
	GetExByParams(ctx context.Context, params *GetMunicipalityParams) ([]entity.MunicipalityEx, error)
}

type CreateMunicipalityData struct {
	Name     string
	RegionID int64
}

type UpdateMunicipalityData struct {
	ID       int64
	Name     *string
	RegionID *int64
	IsHidden *bool
}

type GetMunicipalityParams struct {
	ID       *int64
	Name     *string
	RegionID *int64
	IsHidden *bool
}
