package entity_attribute

import (
	"context"
	"database/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type repo struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.EntityAttributeRepository {
	r := &repo{
		db: m.GetDB(),
	}
	return r
}

func (r *repo) Create(ctx context.Context, obj *entity.EntityAttribute) (*entity.EntityAttribute, error) {
	m := newModel(obj)

	err := r.execQuery(ctx, createQuery, m.EntityTemplateID, m.Name, m.DefaultValue, m.ToShow)
	if err != nil {
		return nil, err
	}

	return r.GetByEntityTemplateIDAndName(ctx, obj.Name, obj.EntityTemplateID)

}

func (r *repo) Update(ctx context.Context, obj *entity.EntityAttribute) (*entity.EntityAttribute, error) {
	m := newModel(obj)

	err := r.execQuery(ctx, updateQuery, m.Name, m.DefaultValue, m.ToShow, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteQuery, id)
}

func (r *repo) GetByEntityTemplateID(ctx context.Context, templateID int64) ([]entity.EntityAttribute, error) {
	return r.fetchRowsWithCondition(ctx, "entity_template_id = $1", templateID)
}

func (r *repo) GetByEntityTemplateIDAndID(ctx context.Context, id, templateID int64) (*entity.EntityAttribute, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND entity_template_id = $2", id, templateID)
}

func (r *repo) GetByID(ctx context.Context, id int64) (*entity.EntityAttribute, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *repo) GetByEntityTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.EntityAttribute, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND entity_template_id = $2", name, templateID)
}
