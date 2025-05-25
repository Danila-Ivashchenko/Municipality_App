package entity_type

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type entityRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.EntityTypeRepository {
	repo := &entityRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *entityRepository) Create(ctx context.Context, data *repository.CreateEntityType) (*entity.EntityType, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createEntityQuery, m.Name)
	if err != nil {
		return nil, err
	}

	return r.GetByName(ctx, data.Name)
}

func (r *entityRepository) Update(ctx context.Context, data *entity.EntityType) error {
	m := newModelEntityType(data)

	return r.execQuery(ctx, updateEntityQuery, m.Name, m.ID)
}

func (r *entityRepository) GetByName(ctx context.Context, name string) (*entity.EntityType, error) {
	return r.fetchRowWithCondition(ctx, "name = $1", name)
}

func (r *entityRepository) GetByNames(ctx context.Context, names []string) ([]entity.EntityType, error) {
	return r.fetchRowsWithCondition(ctx, "name = ANY($1)", sql_common.NewNullStringArray(names))
}

func (r *entityRepository) GetByID(ctx context.Context, id int64) (*entity.EntityType, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *entityRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityType, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *entityRepository) GetAll(ctx context.Context) ([]entity.EntityType, error) {
	return r.fetchRows(ctx, selectEntityQuery)
}

func (r *entityRepository) DeleteByID(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteEntityQuery, id)
}
