package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportFileRepository interface {
	Create(ctx context.Context, passportFile *entity.PassportFile) (*entity.PassportFile, error)
	Delete(ctx context.Context, id int64) error

	GetByPassportID(ctx context.Context, passportID int64) ([]entity.PassportFile, error)
	GetLastByPassportID(ctx context.Context, passportID int64) (*entity.PassportFile, error)
}
