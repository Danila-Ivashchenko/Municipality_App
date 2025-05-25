package passport

import (
	"errors"
	"municipality_app/internal/domain/service"
	"strconv"
	"time"
)

const (
	revisionCodeKey = "revision_code"
)

type UpdatePassportData struct {
	ID             int64
	MunicipalityID int64

	Name string

	Description string
	Year        string

	IsHidden bool
}

type reqUpdatePassport struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Year        *string `json:"year"`
	IsHidden    *bool   `json:"is_hidden"`
	IsMain      *bool   `json:"is_main"`
}

func (req *reqUpdatePassport) Validate() error {
	if req.Year != nil {
		if len(*req.Year) != 4 {
			return errors.New("invalid year value")
		}
	}

	return nil
}

func (req *reqUpdatePassport) Convert(id, municipalityID int64) *service.UpdatePassportData {
	return &service.UpdatePassportData{
		ID:             id,
		MunicipalityID: municipalityID,
		Name:           req.Name,
		Description:    req.Description,
		Year:           req.Year,
		IsHidden:       req.IsHidden,
		IsMain:         req.IsMain,
	}
}

type reqCreatePassport struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Year        *string `json:"year"`
	IsMain      *bool   `json:"is_main"`
}

func (req *reqCreatePassport) Validate() error {
	if req.Name == nil || *req.Name == "" {
		return errors.New("name is required")
	}

	if req.Year != nil {
		if len(*req.Year) != 4 {
			return errors.New("invalid year value")
		}
	}

	return nil
}

func (req *reqCreatePassport) Convert(municipalityID int64) *service.CreatePassportData {
	var (
		yearVal        string
		descriptionVal string
		isMainVal      bool
	)

	if req.Year != nil {
		yearVal = *req.Year
	} else {
		year := time.Now().Year()
		yearVal = strconv.Itoa(year)
	}

	if req.Description != nil {
		descriptionVal = *req.Description
	}

	if req.IsMain != nil {
		isMainVal = *req.IsMain
	}

	return &service.CreatePassportData{
		Name:        *req.Name,
		Description: descriptionVal,
		Year:        yearVal,
		IsMain:      isMainVal,

		MunicipalityID: municipalityID,
	}
}

type reqCreateChapter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Text        string `json:"text"`
	OrderNumber uint   `json:"order_number"`
}

func (req *reqCreateChapter) Validate() error {
	if req.Name == "" {
		return errors.New("name is required")
	}

	return nil
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
	ID             int64                 `json:"id"`
	ChapterData    *reqUpdateChapterData `json:"chapter"`
	PartitionsData *reqPartitionsDelta   `json:"partitions"`
}

type reqUpdateChapterData struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Text        *string `json:"text"`
	OrderNumber *uint   `json:"order_number"`
}

type reqPartitionsDelta struct {
	Create []reqPartitionsData       `json:"create"`
	Update []reqPartitionsDataUpdate `json:"update"`
	Delete []int64                   `json:"delete"`
}

type reqPartitionsData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Text        string `json:"text"`
	OrderNumber uint   `json:"order_number"`
}

type reqPartitionsDataUpdate struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Text        *string `json:"text"`
	OrderNumber *uint   `json:"order_number"`
}

func (req *reqUpdateChapter) Convert(passportID, municipalityID int64) *service.PassportUpdateChaptersData {
	result := &service.PassportUpdateChaptersData{
		PassportID:     passportID,
		MunicipalityID: municipalityID,
	}

	result.ChapterData = service.UpdateChapterExData{
		ChapterID: req.ID,
	}

	if req.ChapterData != nil {
		result.ChapterData.ChapterData = &service.UpdateChapterData{
			ID:          req.ID,
			Name:        req.ChapterData.Name,
			Description: req.ChapterData.Description,
			Text:        req.ChapterData.Text,
			OrderNumber: req.ChapterData.OrderNumber,
		}
	}

	if req.PartitionsData != nil {
		result.ChapterData.PartitionsData = &service.PartitionsData{
			Delete: req.PartitionsData.Delete,
		}

		for _, create := range req.PartitionsData.Create {
			result.ChapterData.PartitionsData.Create = append(result.ChapterData.PartitionsData.Create,
				service.CreateOnePartitionData{
					Name:        create.Name,
					ChapterID:   req.ID,
					Description: create.Description,
					Text:        create.Text,
					OrderNumber: create.OrderNumber,
				},
			)
		}

		for _, update := range req.PartitionsData.Update {
			result.ChapterData.PartitionsData.Update = append(result.ChapterData.PartitionsData.Update,
				service.UpdatePartitionData{
					ID:          update.ID,
					Name:        update.Name,
					Description: update.Description,
					Text:        update.Text,
					OrderNumber: update.OrderNumber,
				},
			)
		}
	}

	return result
}

type reqCopyPassport struct {
	Name   *string `json:"name"`
	Year   *string `json:"year"`
	IsMain *bool   `json:"is_main"`
}

func (req *reqCopyPassport) Validate() error {
	if req.Name == nil || *req.Name == "" {
		return errors.New("name is required")
	}

	if req.Year != nil {
		if len(*req.Year) != 4 {
			return errors.New("invalid year value")
		}
	}

	return nil
}

func (req *reqCopyPassport) Convert(passportID, municipalityID int64) *service.CopyData {
	var (
		yearVal   string
		isMainVal bool
	)

	if req.Year != nil {
		yearVal = *req.Year
	} else {
		year := time.Now().Year()
		yearVal = strconv.Itoa(year)
	}

	if req.IsMain != nil {
		isMainVal = *req.IsMain
	}

	return &service.CopyData{
		NewName: *req.Name,
		NewYear: yearVal,
		IsMain:  isMainVal,

		MunicipalityID: municipalityID,
		SrcID:          passportID,
	}
}
