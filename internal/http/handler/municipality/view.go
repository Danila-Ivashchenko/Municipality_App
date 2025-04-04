package municipality

import (
	"municipality_app/internal/domain/entity"
	"time"
)

type municipalityView struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	RegionID  int64     `json:"region_id"`
	IsHidden  bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at"`
}

func newMunicipalityView(i *entity.Municipality) *municipalityView {
	return &municipalityView{
		ID:        i.ID,
		Name:      i.Name,
		RegionID:  i.RegionID,
		IsHidden:  i.IsHidden,
		CreatedAt: i.CreatedAt,
	}
}

type municipalityExView struct {
	ID        int64                   `json:"id"`
	Name      string                  `json:"name"`
	Region    *municipalityRegionView `json:"region"`
	IsHidden  bool                    `json:"is_hidden"`
	CreatedAt time.Time               `json:"created_at"`
}

type municipalityRegionView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func newMunicipalityRegionView(i *entity.Region) *municipalityRegionView {
	return &municipalityRegionView{
		ID:   i.ID,
		Name: i.Name,
		Code: i.Code,
	}
}

func newMunicipalityExView(i *entity.MunicipalityEx) *municipalityExView {
	view := &municipalityExView{
		ID:        i.ID,
		Name:      i.Name,
		IsHidden:  i.IsHidden,
		CreatedAt: i.CreatedAt,
	}

	if i.Region != nil {
		view.Region = newMunicipalityRegionView(i.Region)
	}

	return view
}

func newMunicipalityExViews(i []entity.MunicipalityEx) []municipalityExView {
	result := make([]municipalityExView, 0, len(i))

	for _, v := range i {
		result = append(result, *newMunicipalityExView(&v))
	}

	return result
}
