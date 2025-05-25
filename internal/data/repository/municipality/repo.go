package municipality

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type municipalityRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.MunicipalityRepository {
	repo := &municipalityRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *municipalityRepository) Create(ctx context.Context, data *repository.CreateMunicipalityData) (*entity.Municipality, error) {
	m := newMunicipalityModelFromCreateData(data)

	err := r.execQuery(ctx, createMunicipality, m.Name, m.RegionID)
	if err != nil {
		return nil, err
	}

	return r.GetByName(ctx, data.Name)
}

func (r *municipalityRepository) Update(ctx context.Context, municipality *entity.Municipality) error {
	m := newMunicipalityModel(municipality)
	return r.execQuery(ctx, updateMunicipalityByID, m.Name, m.RegionID, m.IsHidden, m.ID)
}

func (r *municipalityRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteMunicipalityByID, id)
}

func (r *municipalityRepository) GetById(ctx context.Context, id int64) (*entity.Municipality, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *municipalityRepository) GetByName(ctx context.Context, name string) (*entity.Municipality, error) {
	return r.fetchRowWithCondition(ctx, "name = $1", name)
}

func (r *municipalityRepository) GetAll(ctx context.Context) ([]entity.Municipality, error) {
	return r.fetchRows(ctx, selectMunicipality)
}

func (r *municipalityRepository) GetByParams(ctx context.Context, params *repository.MunicipalityParams) ([]entity.Municipality, error) {
	condition := sql_common.NewCondition()

	if params.ID != nil {
		condition.AddCondition("id = $%d", "AND", *params.ID)
	}

	if params.Name != nil {
		condition.AddCondition("name = $%d", "AND", *params.Name)
	}

	if params.RegionID != nil {
		condition.AddCondition("region_id = $%d", "AND", *params.RegionID)
	}

	if params.IsHidden != nil {
		condition.AddCondition("is_hidden = $%d", "AND", *params.IsHidden)
	}

	if condition.ArgsCount() == 0 {
		return r.fetchRows(ctx, selectMunicipality)
	}

	return r.fetchRowsWithCondition(ctx, condition.Condition(), condition.Args()...)
}
