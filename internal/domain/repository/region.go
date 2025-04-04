package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type RegionRepository interface {
	Create(ctx context.Context, data *CreateRegionData) error
	Delete(ctx context.Context, id int64) error

	GetAll(ctx context.Context) ([]entity.Region, error)
	GetByParams(ctx context.Context, params *RegionParams) ([]entity.Region, error)

	GetById(ctx context.Context, id int64) (*entity.Region, error)
	GetByIds(ctx context.Context, ids []int64) ([]entity.Region, error)

	GetByName(ctx context.Context, name string) (*entity.Region, error)
	GetByCode(ctx context.Context, code string) (*entity.Region, error)
}

type CreateRegionData struct {
	Name string
	Code string
}

type RegionParams struct {
	ID   *int64
	Name *string
	Code *string
}
