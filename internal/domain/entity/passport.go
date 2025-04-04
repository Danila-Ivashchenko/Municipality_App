package entity

import "time"

type Passport struct {
	ID             int64
	Name           string
	MunicipalityID int64

	Description  string
	Year         string
	RevisionCode string

	CreatedAt time.Time
	UpdatedAt time.Time
	IsMain    bool
	IsHidden  bool
}
