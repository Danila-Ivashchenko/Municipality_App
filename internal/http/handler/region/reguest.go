package region

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createRegionReq struct {
	Name *string `json:"name"`
	Code *string `json:"code"`
}

func (r *createRegionReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	if r.Code == nil {
		return errors.New("code is required")
	}

	return nil
}

func (r *createRegionReq) Convert() *service.CreateRegionData {
	return &service.CreateRegionData{
		Code: *r.Code,
		Name: *r.Name,
	}
}

type getByParamsReq struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
	Code *string `json:"code"`
}

func (r *getByParamsReq) Convert() *service.GetRegionParams {
	return &service.GetRegionParams{
		ID:   r.ID,
		Name: r.Name,
		Code: r.Code,
	}
}
