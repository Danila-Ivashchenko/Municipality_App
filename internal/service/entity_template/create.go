package entity_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *objectTemplateService) Create(ctx context.Context, data *service.CreateEntityTemplateData) (*entity.EntityTemplateEx, error) {
	var (
		attributes []entity.EntityAttribute
	)

	templateExist, err := svc.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if templateExist != nil {
		return nil, errors.New("template already exist")
	}

	repoData := &repository.CreateEntityTemplateData{
		Name:           data.Name,
		MunicipalityID: data.MunicipalityID,
		EntityType:     data.EntityType,
	}

	template, err := svc.EntityTemplateRepository.Create(ctx, repoData)
	if err != nil {
		return nil, err
	}

	for _, attributeData := range data.Attributes {
		createAttributeData := service.CreateEntityAttributeData{
			EntityTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.EntityAttributeService.CreateAttribute(ctx, createAttributeData)
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, *attribute)
	}

	entityTyper, err := svc.EntityTypeService.GetByID(ctx, data.EntityType)
	if err != nil {
		return nil, err
	}

	return entity.NewEntityTemplateEx(*template, nil, attributes, entityTyper), nil
}
