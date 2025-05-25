package entity_attribute_value

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type repo struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.EntityAttributeValueRepository {
	r := &repo{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return r
}

func (r *repo) Create(ctx context.Context, obj *entity.EntityAttributeValue) (*entity.EntityAttributeValue, error) {
	var (
		id int64
	)

	m := newModel(obj)

	row := r.handler.QueryRowContext(ctx, createQuery, m.EntityAttributeID, m.EntityID, m.Value)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&id,
	)
	if err != nil {
		return nil, err
	}

	m.ID = sql_common.NewNullInt64(id)

	return m.convert(), err

}

func (r *repo) Update(ctx context.Context, obj *entity.EntityAttributeValue) (*entity.EntityAttributeValue, error) {
	m := newModel(obj)

	err := r.execQuery(ctx, updateQuery, m.Value, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *repo) GetByAttributeID(ctx context.Context, attributeID int64) ([]entity.EntityAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "entity_attribute_id = $1", attributeID)
}

func (r *repo) GetByEntityID(ctx context.Context, entity int64) ([]entity.EntityAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "entity_id = $1", entity)
}

func (r *repo) GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *repo) GetByAttributeIDAndEntityID(ctx context.Context, attributeID, entityID int64) (*entity.EntityAttributeValue, error) {
	return r.fetchRowWithCondition(ctx, " entity_attribute_id = $1 AND entity_id = $2", attributeID, entityID)
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteQuery, id)
}
