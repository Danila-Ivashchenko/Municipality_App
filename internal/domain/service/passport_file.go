package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportFileService interface {
	Create(ctx context.Context, passport *entity.PassportEx) error
}
