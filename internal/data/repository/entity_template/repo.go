package entity_template

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type entityTemplateRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.EntityTemplateRepository {
	repo := &entityTemplateRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *entityTemplateRepository) Create(ctx context.Context, data *repository.CreateEntityTemplateData) (*entity.EntityTemplate, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createEntityQuery, m.Name, m.EntityTypeID, m.MunicipalityID)
	if err != nil {
		return nil, err
	}

	return r.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
}

func (r *entityTemplateRepository) Update(ctx context.Context, data *entity.EntityTemplate) (*entity.EntityTemplate, error) {
	m := newModelEntityTemplate(data)

	err := r.execQuery(ctx, updateEntityQuery, m.Name, m.EntityTypeID, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *entityTemplateRepository) GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.EntityTemplate, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_id = $1", municipalityID)
}

func (r *entityTemplateRepository) GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.EntityTemplate, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND municipality_id = $2", name, municipalityID)
}

func (r *entityTemplateRepository) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.EntityTemplate, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND municipality_id = $2", id, municipalityID)
}

func (r *entityTemplateRepository) GetByID(ctx context.Context, id int64) (*entity.EntityTemplate, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *entityTemplateRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityTemplate, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *entityTemplateRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteEntityQuery, id)
}
