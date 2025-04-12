package entity_type

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createEntityTypesReq struct {
	Data []createEntityTypeReq `json:"data"`
}

type createEntityTypeReq struct {
	Name *string `json:"name"`
}

func (r *createEntityTypesReq) Validate() error {
	for _, d := range r.Data {
		err := d.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *createEntityTypesReq) Convert() *service.CreateEntityTypeMultiplyData {
	var result []service.CreateEntityTypeData

	for _, d := range r.Data {
		result = append(result, *d.Convert())
	}

	return &service.CreateEntityTypeMultiplyData{
		Data: result,
	}
}

func (r *createEntityTypeReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	return nil
}

func (r *createEntityTypeReq) Convert() *service.CreateEntityTypeData {
	return &service.CreateEntityTypeData{
		Name: *r.Name,
	}
}

type updateEntityTypeReq struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
}

func (r *updateEntityTypeReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	if r.ID == 0 {
		return errors.New("id is required")
	}

	return nil
}

func (r *updateEntityTypeReq) Convert() *service.UpdateEntityTypeData {
	return &service.UpdateEntityTypeData{
		ID:   r.ID,
		Name: *r.Name,
	}
}
