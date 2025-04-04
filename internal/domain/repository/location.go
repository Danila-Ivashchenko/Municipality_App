package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type LocationRepository interface {
	Create(ctx context.Context, data *CreateLocationData) (*entity.Location, error)
	Update(ctx context.Context, data *entity.Location) (*entity.Location, error)

	Delete(ctx context.Context, id int64) error

	GetByID(ctx context.Context, id int64) (*entity.Location, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Location, error)
}

type CreateLocationData struct {
	Address   *string
	Latitude  *float64
	Longitude *float64
	Geometry  *string
}
