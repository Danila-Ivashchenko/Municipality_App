package region

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type regionRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.RegionRepository {
	repo := &regionRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *regionRepository) Create(ctx context.Context, data *repository.CreateRegionData) error {
	model := newRegionModel(data)

	err := r.execQuery(ctx, createRegionQuery, model.Name, model.Code)
	if err != nil {
		return err
	}

	return nil
}

func (r *regionRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteRegionQuery+" WHERE id = $1", id)
}

func (r *regionRepository) GetAll(ctx context.Context) ([]entity.Region, error) {
	return r.fetchRows(ctx, selectRegionQuery)
}

func (r *regionRepository) GetByParams(ctx context.Context, params *repository.RegionParams) ([]entity.Region, error) {
	condition := sql_common.NewCondition()

	if params.ID != nil {
		condition.AddCondition("id = $%d", "AND", *params.ID)
	}

	if params.Name != nil {
		condition.AddCondition("name = $%d", "AND", *params.Name)
	}

	if params.Code != nil {
		condition.AddCondition("code = $%d", "AND", *params.Code)
	}

	if condition.ArgsCount() == 0 {
		return r.fetchRows(ctx, selectRegionQuery)
	}

	return r.fetchRowsWithCondition(ctx, condition.Condition(), condition.Args()...)
}

func (r *regionRepository) GetById(ctx context.Context, id int64) (*entity.Region, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *regionRepository) GetByIds(ctx context.Context, ids []int64) ([]entity.Region, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *regionRepository) GetByName(ctx context.Context, name string) (*entity.Region, error) {
	return r.fetchRowWithCondition(ctx, "name = $1", name)
}

func (r *regionRepository) GetByCode(ctx context.Context, code string) (*entity.Region, error) {
	return r.fetchRowWithCondition(ctx, "code = $1", code)
}
