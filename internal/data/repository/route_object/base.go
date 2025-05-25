package route_object

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createRouteObjectQuery = `INSERT INTO municipality_route_object (
                   route_id, 
                   object_name, 
                   order_number,
                   source_object_id,
                   location_id
                   ) VALUES ($1, $2, $3, $4, $5)  returning id;`

	updateRouteObjectQuery = `UPDATE municipality_route_object SET
                   route_id = $1, 
                   object_name = $2, 
                   order_number = $3,
                   source_object_id = $4,
                   location_id = $5
                   WHERE id = $6`

	selectRouteObjectQuery = `SELECT 
    						id,
    						route_id, 
                   			object_name, 
                   			order_number,
                   			source_object_id,
                   			location_id
    					FROM municipality_route_object `

	deleteRouteObjectQuery        = "DELETE FROM municipality_route_object WHERE id = $1"
	deleteRouteObjectToRouteQuery = "DELETE FROM municipality_route_object WHERE route_id = $1"
)

func (r *routeObjectRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	_, err := r.handler.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *routeObjectRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.RouteObject, error) {
	return r.fetchRow(ctx, selectRouteObjectQuery+" WHERE "+condition, args...)
}

func (r *routeObjectRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.RouteObject, error) {
	m := &model{}

	row := r.handler.QueryRowContext(ctx, sqlQuery, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := scan(m, row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return m.convert(), nil
}

func (r *routeObjectRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.RouteObject, error) {
	return r.fetchRows(ctx, selectRouteObjectQuery+" WHERE "+condition, args...)
}

func (r *routeObjectRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.RouteObject, error) {
	var (
		result []entity.RouteObject
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &model{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *model, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.RouteID,
		&m.Name,
		&m.OrderNumber,
		&m.SourceObjectID,
		&m.LocationID,
	)
}
