package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportRepository interface {
	Create(ctx context.Context, data *CreatePassportData) (*entity.Passport, error)
	Update(ctx context.Context, data *entity.Passport) error
	UpdateUpdatedAt(ctx context.Context, passportID int64) error
	Delete(ctx context.Context, id int64) error

	ChangeIsMainByID(ctx context.Context, id int64, isMain bool) error
	ChangeIsMainByIDs(ctx context.Context, ids []int64, isMain bool) error

	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.Passport, error)
	GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.Passport, error)

	GetByIDsAndMunicipalityID(ctx context.Context, ids []int64, municipalityID int64) ([]entity.Passport, error)
	GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.Passport, error)
	GetMainByMunicipalityID(ctx context.Context, municipalityID int64) (*entity.Passport, error)

	GetByRevisionCode(ctx context.Context, revisionCode string) (*entity.Passport, error)
	GetByID(ctx context.Context, id int64) (*entity.Passport, error)
}

type CreatePassportData struct {
	Name           string
	MunicipalityID int64

	Description  string
	Year         string
	RevisionCode string

	IsMain bool
}

type UpdatePassportData struct {
	ID int64

	Name string

	Description string
	Year        string

	IsHidden bool
}
