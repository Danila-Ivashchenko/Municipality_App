package view

import (
	"municipality_app/internal/domain/entity"
	"time"
)

type PassportFileView struct {
	ID         int64     `json:"id"`
	Path       string    `json:"path"`
	CreatedAt  time.Time `json:"created_at"`
	PassportID int64     `json:"passport_id"`
}

func NewPassportFileView(i *entity.PassportFile) *PassportFileView {
	if i == nil {
		return nil
	}

	return &PassportFileView{
		ID:         i.ID,
		Path:       i.Path,
		CreatedAt:  i.CreateAt,
		PassportID: i.PassportID,
	}
}
