package entity

import "time"

type Municipality struct {
	ID        int64
	Name      string
	RegionID  int64
	CreatedAt time.Time
	IsHidden  bool
}

type MunicipalityEx struct {
	ID        int64
	Name      string
	Region    *Region
	CreatedAt time.Time
	IsHidden  bool
}

func NewMunicipalityEx(m *Municipality, r *Region) *MunicipalityEx {
	mEx := &MunicipalityEx{
		ID:        m.ID,
		Name:      m.Name,
		Region:    r,
		CreatedAt: m.CreatedAt,
		IsHidden:  m.IsHidden,
	}

	if r != nil {
		regionValue := *r
		mEx.Region = &regionValue
	}

	return mEx
}
