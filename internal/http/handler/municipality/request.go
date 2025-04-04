package municipality

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createMunicipalityRequest struct {
	Name     *string `json:"name"`
	RegionID *int64  `json:"region_id"`
}

func (req *createMunicipalityRequest) Validate() error {
	if req.Name == nil {
		return errors.New("name is required")
	}
	if req.RegionID == nil {
		return errors.New("region_id is required")
	}

	return nil
}

type updateMunicipalityRequest struct {
	Name     *string `json:"name"`
	RegionID *int64  `json:"region_id"`
	IsHidden *bool   `json:"is_hidden"`
}

func (req *updateMunicipalityRequest) Validate() error {
	paramsCount := 0

	if req.Name != nil {
		paramsCount++
	}

	if req.RegionID != nil {
		paramsCount++
	}

	if req.IsHidden != nil {
		paramsCount++
	}

	if paramsCount == 0 {
		return errors.New("must be at least 1 param")
	}

	return nil
}

func (req *updateMunicipalityRequest) Convert(id int64) *service.UpdateMunicipalityData {
	return &service.UpdateMunicipalityData{
		ID:       id,
		Name:     req.Name,
		RegionID: req.RegionID,
		IsHidden: req.IsHidden,
	}
}

func (req *createMunicipalityRequest) Convert() *service.CreateMunicipalityData {
	return &service.CreateMunicipalityData{
		Name:     *req.Name,
		RegionID: *req.RegionID,
	}
}

type getByParamsRequest struct {
	ID       *int64  `json:"id"`
	Name     *string `json:"name"`
	RegionID *int64  `json:"region_id"`
	IsHidden *bool   `json:"is_hidden"`
}

func (req *getByParamsRequest) Convert() *service.GetMunicipalityParams {
	return &service.GetMunicipalityParams{
		ID:       req.ID,
		Name:     req.Name,
		RegionID: req.RegionID,
		IsHidden: req.IsHidden,
	}
}
