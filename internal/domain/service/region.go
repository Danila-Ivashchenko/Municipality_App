package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type RegionService interface {
	Create(ctx context.Context, data *CreateRegionData) (*entity.Region, error)
	Delete(ctx context.Context, id int64) error

	GetAll(ctx context.Context) ([]entity.Region, error)
	GetByParams(ctx context.Context, params *GetRegionParams) ([]entity.Region, error)
	GetById(ctx context.Context, id int64) (*entity.Region, error)
	GetByName(ctx context.Context, name string) (*entity.Region, error)
	GetByCode(ctx context.Context, code string) (*entity.Region, error)
}

type CreateRegionData struct {
	Name string
	Code string
}

type GetRegionParams struct {
	ID   *int64
	Name *string
	Code *string
}
