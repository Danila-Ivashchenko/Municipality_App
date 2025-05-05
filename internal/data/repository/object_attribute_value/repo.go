package object_attribute_value

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type repo struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.ObjectAttributeValueRepository {
	r := &repo{
		db: m.GetDB(),
	}
	return r
}

func (r *repo) Create(ctx context.Context, obj *entity.ObjectAttributeValue) (*entity.ObjectAttributeValue, error) {
	var (
		id int64
	)

	m := newModel(obj)

	row := r.db.QueryRowContext(ctx, createQuery, m.ObjectAttributeID, m.ObjectID, m.Value)
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

func (r *repo) Update(ctx context.Context, obj *entity.ObjectAttributeValue) (*entity.ObjectAttributeValue, error) {
	m := newModel(obj)

	err := r.execQuery(ctx, updateQuery, m.Value, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *repo) GetByAttributeID(ctx context.Context, attributeID int64) ([]entity.ObjectAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "object_attribute_id = $1 ORDER BY id ASC", attributeID)
}

func (r *repo) GetByObjectID(ctx context.Context, object int64) ([]entity.ObjectAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "object_id = $1 ORDER BY id ASC", object)
}

func (r *repo) GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectAttributeValue, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1) ORDER BY id ASC", sql_common.NewNullInt64Array(ids))
}

func (r *repo) GetByAttributeIDAndObjectID(ctx context.Context, attributeID, objectID int64) (*entity.ObjectAttributeValue, error) {
	return r.fetchRowWithCondition(ctx, " object_attribute_id = $1 AND object_id = $2", attributeID, objectID)
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteQuery, id)
}
