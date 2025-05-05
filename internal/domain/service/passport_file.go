package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportFileService interface {
	Create(ctx context.Context, municipality *entity.Municipality, passport *entity.PassportEx) (*entity.PassportFile, error)

	GetByPassportID(ctx context.Context, passportID int64) ([]entity.PassportFile, error)
	GetLastByPassportID(ctx context.Context, passportID int64) (*entity.PassportFile, error)
}
