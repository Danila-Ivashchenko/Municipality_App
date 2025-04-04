package chapter

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type chapterRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.ChapterRepository {
	repo := &chapterRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *chapterRepository) Create(ctx context.Context, data *repository.CreateChapterData) (*entity.Chapter, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createChapterQuery, m.Name, m.PassportID, m.Description, m.Text, m.OrderNumber)
	if err != nil {
		return nil, err
	}

	return r.GetByNameAndPassportID(ctx, data.Name, data.PassportID)
}

func (r *chapterRepository) CreateMultiply(ctx context.Context, data []repository.CreateChapterData) ([]entity.Chapter, error) {
	result := make([]entity.Chapter, 0, len(data))

	for _, m := range data {
		chapter, err := r.Create(ctx, &m)
		if err != nil {
			return nil, err
		}

		result = append(result, *chapter)
	}

	return result, nil
}

func (r *chapterRepository) Update(ctx context.Context, data *entity.Chapter) error {
	m := newModelChapter(data)

	return r.execQuery(ctx, updateChapterQuery+" WHERE id = $4", m.Name, m.Description, m.Text, m.ID)
}

func (r *chapterRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteChapterQuery+" WHERE id = $1", id)
}

func (r *chapterRepository) ChangeOrder(ctx context.Context, id int64, newOrder uint) error {
	return r.execQuery(ctx, changeChapterOrderQuery+" WHERE id = $2", newOrder, id)
}

func (r *chapterRepository) GetByPassportID(ctx context.Context, passportID int64) ([]entity.Chapter, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_passport_id = $1", passportID)
}

func (r *chapterRepository) GetByIDAndPassportID(ctx context.Context, id, passportID int64) (*entity.Chapter, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND municipality_passport_id = $2", id, passportID)
}

func (r *chapterRepository) GetByID(ctx context.Context, id int64) (*entity.Chapter, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}

func (r *chapterRepository) GetByIDsAndPassportID(ctx context.Context, ids []int64, passportID int64) ([]entity.Chapter, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1) AND municipality_passport_id = $2", sql_common.NewNullInt64Array(ids), passportID)
}

func (r *chapterRepository) GetByNameAndPassportID(ctx context.Context, name string, passportID int64) (*entity.Chapter, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND municipality_passport_id = $2", name, passportID)
}

func (r *chapterRepository) GetByNamesAndPassportID(ctx context.Context, names []string, passportID int64) ([]entity.Chapter, error) {
	return r.fetchRowsWithCondition(ctx, "name = ANY($1) AND municipality_passport_id = $2", sql_common.NewNullStringArray(names), passportID)
}
