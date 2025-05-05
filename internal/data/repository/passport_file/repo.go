package passport_file

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type passportFileRepository struct {
	db *sql.DB
}

func (r *passportFileRepository) Create(ctx context.Context, passportFile *entity.PassportFile) (*entity.PassportFile, error) {
	var (
		id int64
	)

	m := newPassportFileModel(passportFile)

	row := r.db.QueryRowContext(ctx, createPassportFileQuery, m.Path, m.PassportID, m.FileName, m.CreateAt)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	m.ID = sql_common.NewNullInt64(id)

	return m.convert(), nil
}

func (r *passportFileRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deletePassportFileQuery, id)
}

func (r *passportFileRepository) GetLastByPassportID(ctx context.Context, passportID int64) (*entity.PassportFile, error) {
	return r.fetchRowWithCondition(ctx, "passport_id = $1 ORDER BY created_at DESC LIMIT 1", passportID)
}

func (r *passportFileRepository) GetByPassportID(ctx context.Context, passportID int64) ([]entity.PassportFile, error) {
	return r.fetchRowsWithCondition(ctx, "passport_id = $1 ORDER BY created_at DESC", passportID)
}

func New(m db.DataBaseManager) repository.PassportFileRepository {
	repo := &passportFileRepository{
		db: m.GetDB(),
	}
	return repo
}
