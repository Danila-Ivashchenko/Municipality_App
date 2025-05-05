package passport

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
	"time"
)

type passportRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.PassportRepository {
	repo := &passportRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *passportRepository) Create(ctx context.Context, data *repository.CreatePassportData) (*entity.Passport, error) {
	m := newPassportDataFromCreateData(data)

	err := r.execQuery(ctx, createPassportQuery, m.Name, m.MunicipalityID, m.Description, m.Year, m.RevisionCode, m.IsMain)
	if err != nil {
		return nil, err
	}

	return r.GetByRevisionCode(ctx, data.RevisionCode)
}

func (r *passportRepository) Update(ctx context.Context, data *entity.Passport) error {
	m := newPassportModel(data)

	updatedAt := time.Now()

	return r.execQuery(ctx, updatePassportQuery, m.Name, m.Description, m.Year, m.IsHidden, updatedAt, m.ID)
}

func (r *passportRepository) UpdateUpdatedAt(ctx context.Context, passportID int64) error {
	updatedAt := time.Now().UTC()

	return r.execQuery(ctx, updateUpdatedAtPassportQuery, updatedAt, passportID)
}

func (r *passportRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deletePassportQuery, id)
}

func (r *passportRepository) ChangeIsMainByID(ctx context.Context, id int64, isMain bool) error {
	updatedAt := time.Now()

	return r.execQuery(ctx, updateIsMainPassportQuery+" WHERE id = $3", isMain, updatedAt, id)
}

func (r *passportRepository) ChangeIsMainByIDs(ctx context.Context, ids []int64, isMain bool) error {
	updatedAt := time.Now()

	return r.execQuery(ctx, updateIsMainPassportQuery+" WHERE id = any($3)", isMain, updatedAt, sql_common.NewNullInt64Array(ids))
}

func (r *passportRepository) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.Passport, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND municipality_id = $2", id, municipalityID)
}

func (r *passportRepository) GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.Passport, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND municipality_id = $2", name, municipalityID)
}

func (r *passportRepository) GetByIDsAndMunicipalityID(ctx context.Context, ids []int64, municipalityID int64) ([]entity.Passport, error) {
	return r.fetchRowsWithCondition(ctx, "id = any($1) AND municipality_id = $2", sql_common.NewNullInt64Array(ids), municipalityID)
}

func (r *passportRepository) GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.Passport, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_id = $1", municipalityID)
}

func (r *passportRepository) GetMainByMunicipalityID(ctx context.Context, municipalityID int64) (*entity.Passport, error) {
	return r.fetchRowWithCondition(ctx, "is_main = true AND municipality_id = $1", municipalityID)
}

func (r *passportRepository) GetByRevisionCode(ctx context.Context, revisionCode string) (*entity.Passport, error) {
	return r.fetchRowWithCondition(ctx, "revision_code = $1", revisionCode)
}
