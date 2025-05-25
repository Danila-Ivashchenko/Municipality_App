package object_attribute

import (
	"context"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type repo struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.ObjectAttributeRepository {
	r := &repo{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return r
}

func (r *repo) Create(ctx context.Context, obj *entity.ObjectAttribute) (*entity.ObjectAttribute, error) {
	m := newModel(obj)

	err := r.execQuery(ctx, createQuery, m.ObjectTemplateID, m.Name, m.DefaultValue, m.ToShow)
	if err != nil {
		return nil, err
	}

	return r.GetByObjectTemplateIDAndName(ctx, obj.Name, obj.ObjectTemplateID)

}

func (r *repo) Update(ctx context.Context, obj *entity.ObjectAttribute) (*entity.ObjectAttribute, error) {
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

func (r *repo) GetByObjectTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectAttribute, error) {
	return r.fetchRowsWithCondition(ctx, "object_template_id = $1 ORDER BY id ASC", templateID)
}

func (r *repo) GetByObjectTemplateIDAndID(ctx context.Context, id, templateID int64) (*entity.ObjectAttribute, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND object_template_id = $2", id, templateID)
}

func (r *repo) GetByID(ctx context.Context, id int64) (*entity.ObjectAttribute, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *repo) GetByObjectTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.ObjectAttribute, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND object_template_id = $2", name, templateID)
}
