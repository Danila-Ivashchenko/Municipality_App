package object_to_partition

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createQuery = `INSERT INTO municipality_object_to_passport_partition ( municipality_object_id, municipality_passport_partitition_id) VALUES ($1, $2)`
	selectQuery = `SELECT municipality_object_id, municipality_passport_partitition_id FROM municipality_object_to_passport_partition `
	deleteQuery = "DELETE FROM municipality_object_to_passport_partition WHERE  municipality_object_id = $1 AND municipality_passport_partitition_id = $2"
)

func (r *objectToPartitionRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *objectToPartitionRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.ObjectToPartition, error) {
	return r.fetchRow(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *objectToPartitionRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.ObjectToPartition, error) {
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

func (r *objectToPartitionRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.ObjectToPartition, error) {
	return r.fetchRows(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *objectToPartitionRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.ObjectToPartition, error) {
	var (
		result []entity.ObjectToPartition
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
		&m.ObjectID,
		&m.PartitionID,
	)
}
