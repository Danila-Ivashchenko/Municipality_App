package chapter

import (
	"municipality_app/internal/domain/service"
)

type reqCreateChapter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Text        string `json:"text"`
	OrderNumber uint   `json:"order_number"`
}

func (req *reqCreateChapter) Convert(passportID, municipalityID int64) *service.PassportCreateChaptersData {
	result := &service.PassportCreateChaptersData{
		PassportID:     passportID,
		MunicipalityID: municipalityID,
	}

	result.ChapterData = service.CreateChapterExData{
		ChapterData: service.CreateOneChapterData{
			Name:        req.Name,
			PassportID:  passportID,
			Description: req.Description,
			Text:        req.Text,
			OrderNumber: req.OrderNumber,
		},
	}

	return result
}

type reqUpdateChapter struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Text        *string `json:"text"`
	OrderNumber *uint   `json:"order_number"`
}

func (req *reqUpdateChapter) Convert(chapterID, passportID, municipalityID int64) *service.PassportUpdateChaptersData {
	result := &service.PassportUpdateChaptersData{
		PassportID:     passportID,
		MunicipalityID: municipalityID,
	}

	result.ChapterData = service.UpdateChapterExData{
		ChapterID: chapterID,
	}

	result.ChapterData.ChapterData = &service.UpdateChapterData{
		ID:          chapterID,
		Name:        req.Name,
		Description: req.Description,
		Text:        req.Text,
		OrderNumber: req.OrderNumber,
	}

	return result
}
