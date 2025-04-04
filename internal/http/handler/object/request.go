package object

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type createObjectTemplateReq struct {
	Name       *string `json:"name"`
	ObjectType int64   `json:"object_type"`
	Attributes []attributeCreateReq
}

type updateObjectTemplateReq struct {
	Name             *string              `json:"name"`
	ObjectType       *int64               `json:"object_type"`
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

func (r attributeCreateReq) convert() service.CreateObjectAttributeToTemplateData {
	return service.CreateObjectAttributeToTemplateData{
		Name:         r.Name,
		DefaultValue: r.DefaultValue,
		ToShow:       r.ToShow,
	}
}

func (r attributeUpdateReq) convert() service.UpdateObjectAttributeToTemplateData {
	return service.UpdateObjectAttributeToTemplateData{
		ID:           r.ID,
		Name:         r.Name,
		DefaultValue: r.DefaultValue,
		ToShow:       r.ToShow,
	}
}

func (r *createObjectTemplateReq) Validate() error {
	if r.Name == nil {
		return errors.New("name is required")
	}

	if r.ObjectType == 0 {
		return errors.New("object_type is required")
	}

	return nil
}

func (r *updateObjectTemplateReq) Validate() error {
	if r.Name != nil && *r.Name == "" {
		return errors.New("name is required")
	}

	if r.ObjectType != nil && *r.ObjectType == 0 {
		return errors.New("object_type is required")
	}

	return nil
}

func (r *createObjectTemplateReq) Convert(municipalityID int64) *service.CreateObjectTemplateData {
	result := &service.CreateObjectTemplateData{
		Name:           *r.Name,
		MunicipalityID: municipalityID,
		ObjectType:     r.ObjectType,
	}

	for _, attr := range r.Attributes {
		result.Attributes = append(result.Attributes, attr.convert())
	}

	return result
}

func (r *updateObjectTemplateReq) Convert(id, municipalityID int64) *service.UpdateObjectTemplateData {
	result := &service.UpdateObjectTemplateData{
		ID:             id,
		Name:           r.Name,
		MunicipalityID: municipalityID,
		ObjectType:     r.ObjectType,
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

type createObjectReq struct {
	Name            string                    `json:"name"`
	LocationData    *createObjectLocationReq  `json:"location"`
	Description     string                    `json:"description"`
	AttributeValues []attributeValueCreateReq `json:"attribute_values"`
}

type updateObjectReq struct {
	ID              int64                     `json:"ID"`
	Name            *string                   `json:"name"`
	LocationData    *updateObjectLocationReq  `json:"location"`
	Description     *string                   `json:"description"`
	AttributeValues []attributeValueCreateReq `json:"attribute_values"`
}

type attributeValueCreateReq struct {
	AttributeID int64   `json:"attribute_id"`
	Value       *string `json:"value"`
}

func (r attributeValueCreateReq) convert() service.CreateObjectAttributeValueData {
	return service.CreateObjectAttributeValueData{
		AttributeID: r.AttributeID,
		Value:       r.Value,
	}
}

func (r *createObjectReq) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.LocationData != nil {
		if err := r.LocationData.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *updateObjectReq) Validate() error {
	if r.Name != nil && *r.Name == "" {
		return errors.New("name is required")
	}

	if r.LocationData != nil {
		if err := r.LocationData.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *createObjectReq) Convert() service.CreateObjectData {
	result := service.CreateObjectData{
		Name:         r.Name,
		Description:  r.Description,
		LocationData: r.LocationData.Convert(),
	}

	for _, attr := range r.AttributeValues {
		result.AttributeValues = append(result.AttributeValues, attr.convert())
	}

	return result
}

func (r *updateObjectReq) Convert() service.UpdateObjectData {
	result := service.UpdateObjectData{
		ID:           r.ID,
		Name:         r.Name,
		Description:  r.Description,
		LocationData: r.LocationData.Convert(),
	}

	for _, attr := range r.AttributeValues {
		result.AttributeValues = append(result.AttributeValues, attr.convert())
	}

	return result
}

type createObjectLocationReq struct {
	Address   *string  `json:"address"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Geometry  *string  `json:"geometry"`
}

func (r *createObjectLocationReq) Convert() *service.CreateObjectLocationData {
	if r == nil {
		return nil
	}

	return &service.CreateObjectLocationData{
		Address:   r.Address,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Geometry:  r.Geometry,
	}
}

type updateObjectLocationReq struct {
	Address   *string  `json:"address"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Geometry  *string  `json:"geometry"`
}

func (r *updateObjectLocationReq) Convert() *service.UpdateObjectLocationData {
	if r == nil {
		return nil
	}

	return &service.UpdateObjectLocationData{
		Address:   r.Address,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Geometry:  r.Geometry,
	}
}

func (r *createObjectLocationReq) Validate() error {
	var (
		params int
	)

	if r.Address != nil {
		params++
	}
	if r.Latitude != nil {
		params++
	}
	if r.Longitude != nil {
		params++
	}

	if params == 0 {
		return errors.New("location params not provided")
	}

	return nil
}

func (r *updateObjectLocationReq) Validate() error {
	var (
		params int
	)

	if r.Address != nil {
		params++
	}
	if r.Latitude != nil {
		params++
	}
	if r.Longitude != nil {
		params++
	}

	if params == 0 {
		return errors.New("location params not provided")
	}

	return nil
}

type createMultiplyObjetsData struct {
	Data []createObjectReq `json:"data"`
}

type updateMultiplyObjetsData struct {
	Data []updateObjectReq `json:"data"`
}

func (r *createMultiplyObjetsData) Validate() error {
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

func (r *updateMultiplyObjetsData) Validate() error {
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

func (r *createMultiplyObjetsData) Convert(templateID int64) *service.CreateMultiplyObjetsData {
	var (
		objectsData []service.CreateObjectData
	)

	for _, req := range r.Data {
		objectsData = append(objectsData, req.Convert())
	}

	return &service.CreateMultiplyObjetsData{
		ObjectTemplateID: templateID,
		Objects:          objectsData,
	}
}

func (r *updateMultiplyObjetsData) Convert(templateID int64) *service.UpdateMultiplyObjetsData {
	var (
		objectsData []service.UpdateObjectData
	)

	for _, req := range r.Data {
		objectsData = append(objectsData, req.Convert())
	}

	return &service.UpdateMultiplyObjetsData{
		ObjectTemplateID: templateID,
		Objects:          objectsData,
	}
}

type deleteObjectsRequest struct {
	IDs []int64 `json:"ids"`
}
