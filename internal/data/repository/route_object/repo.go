package route_object

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type routeObjectRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.RouteObjectRepository {
	repo := &routeObjectRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *routeObjectRepository) Create(ctx context.Context, data *entity.RouteObject) (*entity.RouteObject, error) {
	var (
		id int64
	)

	m := newModel(data)

	row := r.db.QueryRowContext(ctx, createRouteObjectQuery,
		m.RouteID,
		m.Name,
		m.OrderNumber,
		m.SourceObjectID,
		m.LocationID,
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

func (r *routeObjectRepository) Update(ctx context.Context, data *entity.RouteObject) (*entity.RouteObject, error) {
	m := newModel(data)

	err := r.execQuery(ctx, createRouteObjectQuery,
		m.RouteID,
		m.Name,
		m.OrderNumber,
		m.SourceObjectID,
		m.LocationID,
	)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *routeObjectRepository) GetByRouteID(ctx context.Context, routeID int64) ([]entity.RouteObject, error) {
	return r.fetchRowsWithCondition(ctx, "route_id = $1", routeID)
}

func (r *routeObjectRepository) GetID(ctx context.Context, id int64) (*entity.RouteObject, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *routeObjectRepository) GetByNameAndRouteID(ctx context.Context, name string, routeID int64) (*entity.RouteObject, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 and route_id = $2", name, routeID)
}

func (r *routeObjectRepository) GetByNamesAndRouteID(ctx context.Context, names []string, routeID int64) ([]entity.RouteObject, error) {
	return r.fetchRowsWithCondition(ctx, "name = ANY($1) and route_id = $2", sql_common.NewNullStringArray(names), routeID)
}

func (r *routeObjectRepository) GetIDs(ctx context.Context, ids []int64) ([]entity.RouteObject, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *routeObjectRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteRouteObjectQuery, id)
}

func (r *routeObjectRepository) DeleteToRoute(ctx context.Context, routeId int64) error {
	return r.execQuery(ctx, deleteRouteObjectToRouteQuery, routeId)
}
