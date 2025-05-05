package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityExService interface {
	GetByID(ctx context.Context, id int64) (*entity.EntityTemplateEx, error)
	GetByMunicipalityID(ctx context.Context, id int64) ([]entity.EntityTemplateEx, error)
}
