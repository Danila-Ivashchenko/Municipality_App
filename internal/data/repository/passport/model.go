package passport

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type passportModel struct {
	ID             sql.NullInt64
	Name           sql.NullString
	MunicipalityID sql.NullInt64
	Description    sql.NullString
	Year           sql.NullString
	RevisionCode   sql.NullString
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	IsMain         sql.NullBool
	IsHidden       sql.NullBool
}

func (m *passportModel) convert() *entity.Passport {
	return &entity.Passport{
		ID:             m.ID.Int64,
		Name:           m.Name.String,
		MunicipalityID: m.MunicipalityID.Int64,
		Description:    m.Description.String,
		Year:           m.Year.String,
		RevisionCode:   m.RevisionCode.String,
		CreatedAt:      m.CreatedAt.Time,
		UpdatedAt:      m.UpdatedAt.Time,
		IsMain:         m.IsMain.Bool,
		IsHidden:       m.IsHidden.Bool,
	}
}

func newPassportModel(i *entity.Passport) *passportModel {
	return &passportModel{
		ID:             sql_common.NewNullInt64(i.ID),
		Name:           sql_common.NewNullString(i.Name),
		MunicipalityID: sql_common.NewNullInt64(i.MunicipalityID),
		Description:    sql_common.NewNullString(i.Description),
		Year:           sql_common.NewNullString(i.Year),
		RevisionCode:   sql_common.NewNullString(i.RevisionCode),
		CreatedAt:      sql_common.NewNullTime(i.CreatedAt),
		UpdatedAt:      sql_common.NewNullTime(i.UpdatedAt),
		IsMain:         sql_common.NewNullBool(i.IsMain),
		IsHidden:       sql_common.NewNullBool(i.IsHidden),
	}
}

func newPassportDataFromCreateData(i *repository.CreatePassportData) *passportModel {
	return &passportModel{
		Name:           sql_common.NewNullString(i.Name),
		MunicipalityID: sql_common.NewNullInt64(i.MunicipalityID),
		Description:    sql_common.NewNullString(i.Description),
		Year:           sql_common.NewNullString(i.Year),
		RevisionCode:   sql_common.NewNullString(i.RevisionCode),
		IsMain:         sql_common.NewNullBool(i.IsMain),
	}
}

func newPassportDataFromUpdateData(i *repository.UpdatePassportData) *passportModel {
	return &passportModel{
		ID:          sql_common.NewNullInt64(i.ID),
		Name:        sql_common.NewNullString(i.Name),
		Description: sql_common.NewNullString(i.Description),
		Year:        sql_common.NewNullString(i.Year),
		IsHidden:    sql_common.NewNullBool(i.IsHidden),
	}
}
