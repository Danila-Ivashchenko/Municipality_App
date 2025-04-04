package object_type

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createObjectTypesReq struct {
	Data []createObjectTypeReq `json:"data"`
}

type createObjectTypeReq struct {
	Name *string `json:"name"`
}

func (r *createObjectTypesReq) Validate() error {
	for _, d := range r.Data {
		err := d.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *createObjectTypesReq) Convert() *service.CreateObjectTypeMultiplyData {
	var result []service.CreateObjectTypeData

	for _, d := range r.Data {
		result = append(result, *d.Convert())
	}

	return &service.CreateObjectTypeMultiplyData{
		Data: result,
	}
}

func (r *createObjectTypeReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	return nil
}

func (r *createObjectTypeReq) Convert() *service.CreateObjectTypeData {
	return &service.CreateObjectTypeData{
		Name: *r.Name,
	}
}

type updateObjectTypeReq struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
}

func (r *updateObjectTypeReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	if r.ID == 0 {
		return errors.New("id is required")
	}

	return nil
}

func (r *updateObjectTypeReq) Convert() *service.UpdateObjectTypeData {
	return &service.UpdateObjectTypeData{
		ID:   r.ID,
		Name: *r.Name,
	}
}
