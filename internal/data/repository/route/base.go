package route

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createRouteQuery = `INSERT INTO route (
                   municipality_passport_partitition_id, 
                   name, 
                   length,
                   duration,
                   level,
                   movement_way,
                   seasonality,
                   personal_equipment,
                   dangers,
                   rules,
                   route_equipment,
                   geometry
                   ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)  returning id;`

	updateRouteQuery = `UPDATE route SET
                   name = $1, 
                   length = $2,
                   duration = $3,
                   level = $4,
                   movement_way = $5,
                   seasonality = $6,
                   personal_equipment = $7,
                   dangers = $8,
                   rules = $9,
                   route_equipment = $10,
                   geometry = $11
                   WHERE id = $12`

	selectRouteQuery = `SELECT 
    						id,
    						municipality_passport_partitition_id, 
                   			name, 
                   			length,
                   			duration,
                   			level,
                   			movement_way,
                   			seasonality,
                   			personal_equipment,
                   			dangers,
                   			rules,
                   			route_equipment,
                   			geometry
    					FROM route `

	deleteRouteQuery = "DELETE FROM route WHERE id = $1"
)

func (r *routeRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	res, err := r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("now rows affected")
	}

	return nil
}

func (r *routeRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Route, error) {
	return r.fetchRow(ctx, selectRouteQuery+" WHERE "+condition, args...)
}

func (r *routeRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Route, error) {
	m := &model{}

	row := r.db.QueryRowContext(ctx, sqlQuery, args...)
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

func (r *routeRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Route, error) {
	return r.fetchRows(ctx, selectRouteQuery+" WHERE "+condition, args...)
}

func (r *routeRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Route, error) {
	var (
		result []entity.Route
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
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
		&m.PartitionID,
		&m.Name,
		&m.Length,
		&m.Duration,
		&m.Level,
		&m.MovementWay,
		&m.Seasonality,
		&m.PersonalEquipment,
		&m.Dangerous,
		&m.Rules,
		&m.RouteEquipment,
		&m.Geometry,
	)
}
