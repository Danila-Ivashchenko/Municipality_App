package chapter

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createChapterQuery      = `INSERT INTO municipality_passport_chapter (name, municipality_passport_id, description, chapter_text, order_number) VALUES ($1, $2, $3, $4, $5)`
	selectChapterQuery      = `SELECT id, name, municipality_passport_id, description, chapter_text, order_number FROM municipality_passport_chapter `
	updateChapterQuery      = `UPDATE municipality_passport_chapter SET name = $1, description = $2, chapter_text = $3`
	changeChapterOrderQuery = `UPDATE municipality_passport_chapter SET order_number = $1`
	deleteChapterQuery      = "DELETE FROM municipality_passport_chapter"
)

func (r *chapterRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *chapterRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Chapter, error) {
	return r.fetchRow(ctx, selectChapterQuery+" WHERE "+condition, args...)
}

func (r *chapterRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Chapter, error) {
	m := &modelChapter{}

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

func (r *chapterRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Chapter, error) {
	return r.fetchRows(ctx, selectChapterQuery+" WHERE "+condition, args...)
}

func (r *chapterRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Chapter, error) {
	var (
		result []entity.Chapter
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelChapter{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelChapter, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.PassportID,
		&m.Description,
		&m.Text,
		&m.OrderNumber,
	)
}
