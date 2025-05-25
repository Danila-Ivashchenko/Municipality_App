package route

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type routeRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.RouteRepository {
	repo := &routeRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *routeRepository) Create(ctx context.Context, data *entity.Route) (*entity.Route, error) {
	var (
		id int64
	)

	m := newModel(data)

	row := r.handler.QueryRowContext(ctx, createRouteQuery,
		m.PartitionID,
		m.Name,
		m.Length,
		m.Duration,
		m.Level,
		m.MovementWay,
		m.Seasonality,
		m.PersonalEquipment,
		m.Dangerous,
		m.Rules,
		m.RouteEquipment,
		m.Geometry,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&id,
	)
	if err != nil {
		return nil, err
	}

	m.ID = sql_common.NewNullInt64(id)

	return m.convert(), err
}

func (r *routeRepository) Update(ctx context.Context, data *entity.Route) (*entity.Route, error) {
	m := newModel(data)

	err := r.execQuery(ctx, updateRouteQuery,
		m.Name,
		m.Length,
		m.Duration,
		m.Level,
		m.MovementWay,
		m.Seasonality,
		m.PersonalEquipment,
		m.Dangerous,
		m.Rules,
		m.RouteEquipment,
		m.Geometry,
		m.ID,
	)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *routeRepository) GetByID(ctx context.Context, id int64) (*entity.Route, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *routeRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Route, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *routeRepository) GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.Route, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_passport_partitition_id = $1", sql_common.NewNullInt64(partitionID))
}

func (r *routeRepository) GetByNamePartitionID(ctx context.Context, name string, partitionID int64) (*entity.Route, error) {
	return r.fetchRowWithCondition(ctx, "municipality_passport_partitition_id = $1 and name = $2", sql_common.NewNullInt64(partitionID), name)
}

func (r *routeRepository) GetByIDPartitionID(ctx context.Context, id, partitionID int64) (*entity.Route, error) {
	return r.fetchRowWithCondition(ctx, "municipality_passport_partitition_id = $1 and id = $2", sql_common.NewNullInt64(partitionID), id)
}

func (r *routeRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteRouteQuery, id)
}
