package object_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *objectTemplateService) Create(ctx context.Context, data *service.CreateObjectTemplateData) (*entity.ObjectTemplateEx, error) {
	var (
		attributes []entity.ObjectAttribute
	)

	templateExist, err := svc.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if templateExist != nil {
		return nil, errors.New("template already exist")
	}

	repoData := &repository.CreateObjectTemplateData{
		Name:           data.Name,
		MunicipalityID: data.MunicipalityID,
		ObjectType:     data.ObjectType,
	}

	template, err := svc.ObjectTemplateRepository.Create(ctx, repoData)
	if err != nil {
		return nil, err
	}

	for _, attributeData := range data.Attributes {
		createAttributeData := service.CreateObjectAttributeData{
			ObjectTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.ObjectAttributeService.CreateAttribute(ctx, createAttributeData)
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, *attribute)
	}

	objectType, err := svc.ObjectTypeService.GetByID(ctx, data.ObjectType)
	if err != nil {
		return nil, err
	}

	return entity.NewObjectTemplateEx(*template, nil, attributes, objectType), nil
}
