package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportService interface {
	Create(ctx context.Context, data *CreatePassportData) (*entity.Passport, error)
	Update(ctx context.Context, data *UpdatePassportData) (*entity.Passport, error)
	UpdatedAt(ctx context.Context, passportID int64) error
	Delete(ctx context.Context, id, municipalityID int64) error

	MakeMainPassportToMunicipality(ctx context.Context, id, municipalityID int64) error

	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.Passport, error)
	GetByIDsAndMunicipalityID(ctx context.Context, ids []int64, municipalityID int64) ([]entity.Passport, error)
	GetMainByMunicipalityID(ctx context.Context, municipalityID int64) (*entity.Passport, error)
	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.Passport, error)

	GetByRevisionCode(ctx context.Context, revisionCode string) (*entity.Passport, error)
}

type CreatePassportData struct {
	Name           string
	MunicipalityID int64

	Description string
	Year        string

	IsMain bool
}

type UpdatePassportData struct {
	ID             int64
	MunicipalityID int64

	Name *string

	Description *string
	Year        *string
	IsMain      *bool

	IsHidden *bool
}
