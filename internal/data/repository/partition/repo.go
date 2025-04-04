package partition

import (
	"context"
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type partitionRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.PartitionRepository {
	repo := &partitionRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *partitionRepository) Create(ctx context.Context, data *repository.CreatePartitionData) (*entity.Partition, error) {
	m := newModelFromCreateData(data)

	err := r.execQuery(ctx, createChapterQuery, m.Name, m.ChapterID, m.Description, m.Text, m.OrderNumber)
	if err != nil {
		return nil, err
	}

	return r.GetByNameAndChapterID(ctx, data.Name, data.ChapterID)
}

func (r *partitionRepository) Update(ctx context.Context, data *entity.Partition) error {
	m := newModelPartition(data)

	return r.execQuery(ctx, updateChapterQuery+" WHERE id = $4", m.Name, m.Description, m.Text, m.ID)
}

func (r *partitionRepository) Delete(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteChapterQuery+" WHERE id = $1", id)
}

func (r *partitionRepository) ChangeOrder(ctx context.Context, id int64, newOrder uint) error {
	return r.execQuery(ctx, changeChapterOrderQuery+" WHERE id = $2", newOrder, id)
}

func (r *partitionRepository) GetByChapterID(ctx context.Context, chapterID int64) ([]entity.Partition, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_passport_chapter_id = $1", chapterID)
}

func (r *partitionRepository) GetByIDAndChapterID(ctx context.Context, id, chapterID int64) (*entity.Partition, error) {
	return r.fetchRowWithCondition(ctx, "id = $1 AND municipality_passport_chapter_id = $2", id, chapterID)
}

func (r *partitionRepository) GetByID(ctx context.Context, id int64) (*entity.Partition, error) {
	return r.fetchRowWithCondition(ctx, "id = $1", id)
}
func (r *partitionRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Partition, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1)", sql_common.NewNullInt64Array(ids))
}

func (r *partitionRepository) GetByIDsAndChapterID(ctx context.Context, ids []int64, chapterID int64) ([]entity.Partition, error) {
	return r.fetchRowsWithCondition(ctx, "id = ANY($1) AND municipality_passport_chapter_id = $2", sql_common.NewNullInt64Array(ids), chapterID)
}

func (r *partitionRepository) GetByNameAndChapterID(ctx context.Context, name string, chapterID int64) (*entity.Partition, error) {
	return r.fetchRowWithCondition(ctx, "name = $1 AND municipality_passport_chapter_id = $2", name, chapterID)
}

func (r *partitionRepository) GetByNamesAndChapterID(ctx context.Context, names []string, chapterID int64) ([]entity.Partition, error) {
	return r.fetchRowsWithCondition(ctx, "name = ANY($1) AND municipality_passport_chapter_id = $2", sql_common.NewNullStringArray(names), chapterID)
}
