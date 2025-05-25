package object_type

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type objectRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.ObjectTypeRepository {
	repo := &objectRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *objectRepository) Create(ctx context.Context, data *repository.CreateObjectType) (*entity.ObjectType, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createObjectQuery, m.Name)
	if err != nil {
		return nil, err
	}

	return r.GetByName(ctx, data.Name)
}

func (r *objectRepository) Update(ctx context.Context, data *entity.ObjectType) error {
	m := newModelObjectType(data)

	return r.execQuery(ctx, updateObjectQuery, m.Name, m.ID)
}

func (r *objectRepository) GetByName(ctx context.Context, name string) (*entity.ObjectType, error) {
	return r.fetchRowWithCondition(ctx, "name = $1", name)
}

func (r *objectRepository) GetByNames(ctx context.Context, names []string) ([]entity.ObjectType, error) {
	return r.fetchRowsWithCondition(ctx, "name = ANY($1)", sql_common.NewNullStringArray(names))
}

func (r *objectRepository) GetByID(ctx context.Context, id int64) (*entity.ObjectType, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *objectRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectType, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *objectRepository) GetAll(ctx context.Context) ([]entity.ObjectType, error) {
	return r.fetchRows(ctx, selectObjectQuery)
}

func (r *objectRepository) DeleteByID(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteObjectQuery, id)
}
