package object

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createEntityTemplateReq struct {
	Name       *string `json:"name"`
	EntityType int64   `json:"object_type"`
	Attributes []attributeCreateReq
}

type updateEntityTemplateReq struct {
	Name             *string              `json:"name"`
	EntityType       *int64               `json:"object_type"`
	AttributesUpdate []attributeUpdateReq `json:"attributes_update"`
	AttributesCreate []attributeCreateReq `json:"attributes_create"`
	AttributesDelete []int64              `json:"attributes_delete"`
}

type attributeCreateReq struct {
	Name         string `json:"name"`
	DefaultValue string `json:"default_value"`
	ToShow       bool   `json:"to_show"`
}

type attributeUpdateReq struct {
	ID           int64   `json:"id"`
	Name         *string `json:"name"`
	DefaultValue *string `json:"default_value"`
	ToShow       *bool   `json:"to_show"`
}

func (r attributeCreateReq) convert() service.CreateEntityAttributeToTemplateData {
	return service.CreateEntityAttributeToTemplateData{
		Name:         r.Name,
		DefaultValue: r.DefaultValue,
		ToShow:       r.ToShow,
	}
}

func (r attributeUpdateReq) convert() service.UpdateEntityAttributeToTemplateData {
	return service.UpdateEntityAttributeToTemplateData{
		ID:           r.ID,
		Name:         r.Name,
		DefaultValue: r.DefaultValue,
		ToShow:       r.ToShow,
	}
}

func (r *createEntityTemplateReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	if r.EntityType == 0 {
		return errors.New("object_type is required")
	}

	return nil
}

func (r *updateEntityTemplateReq) Validate() error {
	if r.Name != nil && *r.Name == "" {
		return errors.New("name is required")
	}

	if r.EntityType != nil && *r.EntityType == 0 {
		return errors.New("object_type is required")
	}

	return nil
}

func (r *createEntityTemplateReq) Convert(municipalityID int64) *service.CreateEntityTemplateData {
	result := &service.CreateEntityTemplateData{
		Name:           *r.Name,
		MunicipalityID: municipalityID,
		EntityType:     r.EntityType,
	}

	for _, attr := range r.Attributes {
		result.Attributes = append(result.Attributes, attr.convert())
	}

	return result
}

func (r *updateEntityTemplateReq) Convert(id, municipalityID int64) *service.UpdateEntityTemplateData {
	result := &service.UpdateEntityTemplateData{
		ID:             id,
		Name:           r.Name,
		MunicipalityID: municipalityID,
		EntityType:     r.EntityType,
	}

	for _, attr := range r.AttributesUpdate {
		result.AttributesToUpdate = append(result.AttributesToUpdate, attr.convert())
	}

	for _, attr := range r.AttributesCreate {
		result.AttributesToCreate = append(result.AttributesToCreate, attr.convert())
	}

	result.AttributesToDelete = append(result.AttributesToDelete, r.AttributesDelete...)

	return result
}

type createEntityReq struct {
	Name            string                    `json:"name"`
	Description     string                    `json:"description"`
	AttributeValues []attributeValueCreateReq `json:"attribute_values"`
}

type updateEntityReq struct {
	ID              int64                     `json:"ID"`
	Name            *string                   `json:"name"`
	Description     *string                   `json:"description"`
	AttributeValues []attributeValueCreateReq `json:"attribute_values"`
}

type attributeValueCreateReq struct {
	AttributeID int64   `json:"attribute_id"`
	Value       *string `json:"value"`
}

func (r attributeValueCreateReq) convert() service.CreateEntityAttributeValueData {
	return service.CreateEntityAttributeValueData{
		AttributeID: r.AttributeID,
		Value:       r.Value,
	}
}

func (r *createEntityReq) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func (r *updateEntityReq) Validate() error {
	if r.Name != nil && *r.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func (r *createEntityReq) Convert() service.CreateEntityData {
	result := service.CreateEntityData{
		Name:        r.Name,
		Description: r.Description,
	}

	for _, attr := range r.AttributeValues {
		result.AttributeValues = append(result.AttributeValues, attr.convert())
	}

	return result
}

func (r *updateEntityReq) Convert() service.UpdateEntityData {
	result := service.UpdateEntityData{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}

	for _, attr := range r.AttributeValues {
		result.AttributeValues = append(result.AttributeValues, attr.convert())
	}

	return result
}

type createMultiplyEntitiesData struct {
	Data []createEntityReq `json:"data"`
}

type updateMultiplyEntitiesData struct {
	Data []updateEntityReq `json:"data"`
}

func (r *createMultiplyEntitiesData) Validate() error {
	if len(r.Data) == 0 {
		return errors.New("data not provided")
	}

	for _, req := range r.Data {
		if err := req.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *updateMultiplyEntitiesData) Validate() error {
	if len(r.Data) == 0 {
		return errors.New("data not provided")
	}

	for _, req := range r.Data {
		if err := req.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *createMultiplyEntitiesData) Convert(templateID int64) *service.CreateMultiplyEntitiesData {
	var (
		objectsData []service.CreateEntityData
	)

	for _, req := range r.Data {
		objectsData = append(objectsData, req.Convert())
	}

	return &service.CreateMultiplyEntitiesData{
		EntityTemplateID: templateID,
		Entities:         objectsData,
	}
}

func (r *updateMultiplyEntitiesData) Convert(templateID int64) *service.UpdateMultiplyEntitiesData {
	var (
		objectsData []service.UpdateEntityData
	)

	for _, req := range r.Data {
		objectsData = append(objectsData, req.Convert())
	}

	return &service.UpdateMultiplyEntitiesData{
		EntityTemplateID: templateID,
		Entities:         objectsData,
	}
}

type deleteEntitysRequest struct {
	IDs []int64 `json:"ids"`
}
