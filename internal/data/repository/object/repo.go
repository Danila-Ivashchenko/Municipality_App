package object

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type objectTemplateRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.ObjectRepository {
	repo := &objectTemplateRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *objectTemplateRepository) Create(ctx context.Context, data *repository.CreateObjectData) (*entity.Object, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createObjectQuery, m.Name, m.ObjectTemplateID, m.LocationID, m.Description)
	if err != nil {
		return nil, err
	}

	return r.GetByTemplateIDAndName(ctx, data.Name, data.ObjectTemplateID)
}

func (r *objectTemplateRepository) Update(ctx context.Context, data *entity.Object) (*entity.Object, error) {
	m := newModelObject(data)

	err := r.execQuery(ctx, updateObjectQuery, m.Name, m.Description, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r *objectTemplateRepository) GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Object, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_object_template_id = $1", templateID)
}

func (r *objectTemplateRepository) GetByTemplateIDAndName(ctx context.Context, name string, templateID int64) (*entity.Object, error) {
	return r.fetchRowWithCondition(ctx, "municipality_object_template_id = $1 and name = $2", templateID, name)
}

func (r *objectTemplateRepository) GetByTemplateIDAndNames(ctx context.Context, names []string, templateID int64) ([]entity.Object, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_object_template_id = $1 and name = ANY($2)", templateID, sql_common.NewNullStringArray(names))
}

func (r *objectTemplateRepository) GetByIDsAndTemplateID(ctx context.Context, ids []int64, templateID int64) ([]entity.Object, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_object_template_id = $1 and id = ANY($2)", templateID, sql_common.NewNullInt64Array(ids))
}

func (r *objectTemplateRepository) GetByID(ctx context.Context, id int64) (*entity.Object, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *objectTemplateRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Object, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *objectTemplateRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteObjectQuery, id)
}
