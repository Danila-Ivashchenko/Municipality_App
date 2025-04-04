package passport

import (
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/http/view"
	"sort"
	"time"
)

type passportView struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	MunicipalityID int64  `json:"municipality_id"`

	Description  string `json:"description"`
	Year         string `json:"year"`
	RevisionCode string `json:"revisionCode"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsMain    bool      `json:"is_main"`
	IsHidden  bool      `json:"is_hidden"`
}

func newPassportView(i *entity.Passport) *passportView {
	return &passportView{
		ID:             i.ID,
		Name:           i.Name,
		MunicipalityID: i.MunicipalityID,
		Description:    i.Description,
		Year:           i.Year,
		RevisionCode:   i.RevisionCode,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
		IsMain:         i.IsMain,
		IsHidden:       i.IsHidden,
	}
}

func newPassportViews(i []entity.Passport) []passportView {
	result := make([]passportView, 0, len(i))

	for _, v := range i {
		result = append(result, *newPassportView(&v))
	}

	return result
}

type passportExView struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	MunicipalityID int64  `json:"municipality_id"`

	Description  string `json:"description"`
	Year         string `json:"year"`
	RevisionCode string `json:"revisionCode"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsMain    bool      `json:"is_main"`
	IsHidden  bool      `json:"is_hidden"`

	Chapters []view.ChapterExView `json:"chapters"`
}

func newPassportExView(i *entity.PassportEx) *passportExView {
	passportEx := &passportExView{
		ID:             i.ID,
		Name:           i.Name,
		MunicipalityID: i.MunicipalityID,
		Description:    i.Description,
		Year:           i.Year,
		RevisionCode:   i.RevisionCode,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
		IsMain:         i.IsMain,
		IsHidden:       i.IsHidden,
	}

	for _, ch := range i.Chapters {
		passportEx.Chapters = append(passportEx.Chapters, view.NewChapterExView(ch))
	}

	sort.Slice(passportEx.Chapters, func(i, j int) bool {
		return passportEx.Chapters[i].OrderNumber < passportEx.Chapters[j].OrderNumber
	})

	return passportEx
}
