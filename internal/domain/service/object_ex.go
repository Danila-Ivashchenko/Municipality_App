package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectExService interface {
	GetByID(ctx context.Context, id int64) (*entity.ObjectTemplateEx, error)
	GetByMunicipalityID(ctx context.Context, id int64) ([]entity.ObjectTemplateEx, error)
}
