package location

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type locationRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.LocationRepository {
	repo := &locationRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *locationRepository) Create(ctx context.Context, data *repository.CreateLocationData) (*entity.Location, error) {
	m := newModelFromCreateData(data)

	row := r.handler.QueryRowContext(ctx, createLocationQuery, m.Address, m.Latitude, m.Longitude, m.Geometry)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&m.ID,
	)

	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *locationRepository) Update(ctx context.Context, data *entity.Location) (*entity.Location, error) {
	m := newModelObject(data)

	err := r.execQuery(ctx, updateLocationQuery, m.Address, m.Latitude, m.Longitude, m.Geometry, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *locationRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteLocationQuery, id)
}

func (r *locationRepository) GetByID(ctx context.Context, id int64) (*entity.Location, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *locationRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Location, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}
