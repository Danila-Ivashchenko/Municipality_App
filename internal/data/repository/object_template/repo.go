package object_template

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type objectTemplateRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.ObjectTemplateRepository {
	repo := &objectTemplateRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *objectTemplateRepository) Create(ctx context.Context, data *repository.CreateObjectTemplateData) (*entity.ObjectTemplate, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createObjectQuery, m.Name, m.ObjectTypeID, m.MunicipalityID)
	if err != nil {
		return nil, err
	}

	return r.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
}

func (r *objectTemplateRepository) Update(ctx context.Context, data *entity.ObjectTemplate) (*entity.ObjectTemplate, error) {
	m := newModelObjectTemplate(data)

	err := r.execQuery(ctx, updateObjectQuery, m.Name, m.ObjectTypeID, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *objectTemplateRepository) GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.ObjectTemplate, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_id = $1", municipalityID)
}

func (r *objectTemplateRepository) GetByTypeID(ctx context.Context, typeID int64) ([]entity.ObjectTemplate, error) {
	return r.fetchRowsWithCondition(ctx, "object_type_id = $1", typeID)
}

func (r *objectTemplateRepository) GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.ObjectTemplate, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND municipality_id = $2", name, municipalityID)
}

func (r *objectTemplateRepository) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.ObjectTemplate, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND municipality_id = $2", id, municipalityID)
}

func (r *objectTemplateRepository) GetByID(ctx context.Context, id int64) (*entity.ObjectTemplate, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *objectTemplateRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectTemplate, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1) ORDER BY id ASC", sql_common.NewNullInt64Array(ids))
}

func (r *objectTemplateRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteObjectQuery, id)
}
