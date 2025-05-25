package entity

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

func New(m db.DataBaseManager) repository.EntityRepository {
	repo := &entityTemplateRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *entityTemplateRepository) Create(ctx context.Context, data *repository.CreateEntityData) (*entity.Entity, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createEntityQuery, m.Name, m.EntityTemplateID, m.Description)
	if err != nil {
		return nil, err
	}

	return r.GetByTemplateIDAndName(ctx, data.Name, data.EntityTemplateID)
}

func (r *entityTemplateRepository) Update(ctx context.Context, data *entity.Entity) (*entity.Entity, error) {
	m := newModelEntity(data)

	err := r.execQuery(ctx, updateEntityQuery, m.Name, m.Description, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *entityTemplateRepository) GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Entity, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_entity_template_id = $1", templateID)
}

func (r *entityTemplateRepository) GetByTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.Entity, error) {
	return r.fetchRowWithCondition(ctx, "municipality_entity_template_id = $1 and name = $2", templateID, name)
}

func (r *entityTemplateRepository) GetByTemplateIDAndNames(ctx context.Context, names []string, templateID int64) ([]entity.Entity, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_entity_template_id = $1 and name = ANY($2)", templateID, sql_common.NewNullStringArray(names))
}

func (r *entityTemplateRepository) GetByIDsAndTemplateID(ctx context.Context, ids []int64, templateID int64) ([]entity.Entity, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_entity_template_id = $1 and id = ANY($2)", templateID, sql_common.NewNullInt64Array(ids))
}

func (r *entityTemplateRepository) GetByID(ctx context.Context, id int64) (*entity.Entity, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *entityTemplateRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Entity, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *entityTemplateRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteEntityQuery, id)
}
