package entity_to_partition

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createQuery = `INSERT INTO municipality_entity_to_passport_partition ( municipality_entity_id, municipality_passport_partitition_id) VALUES ($1, $2)`
	selectQuery = `SELECT municipality_entity_id, municipality_passport_partitition_id FROM municipality_entity_to_passport_partition `
	deleteQuery = "DELETE FROM municipality_entity_to_passport_partition WHERE  municipality_entity_id = $1 AND municipality_passport_partitition_id = $2"
)

func (r *entityToPartitionRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	res, err := r.handler.ExecContext(ctx, sqlQuery, args...)
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

func (r *entityToPartitionRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.EntityToPartition, error) {
	return r.fetchRow(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *entityToPartitionRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.EntityToPartition, error) {
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

func (r *entityToPartitionRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.EntityToPartition, error) {
	return r.fetchRows(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *entityToPartitionRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.EntityToPartition, error) {
	var (
		result []entity.EntityToPartition
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
		&m.EntityID,
		&m.PartitionID,
	)
}
