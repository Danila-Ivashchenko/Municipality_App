package entity_template

import (
	"context"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *objectTemplateService) Create(ctx context.Context, data *service.CreateEntityTemplateData) (*entity.EntityTemplateEx, error) {
	var (
		attributes  []entity.EntityAttribute
		entityTyper *entity.EntityType
		template    *entity.EntityTemplate
	)

	templateExist, err := svc.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if templateExist != nil {
		return nil, core_errors.EntityTemplateNameIsUsed
	}

	repoData := &repository.CreateEntityTemplateData{
		Name:           data.Name,
		MunicipalityID: data.MunicipalityID,
		EntityType:     data.EntityType,
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		template, err = svc.EntityTemplateRepository.Create(tx, repoData)
		if err != nil {
			return err
		}

		for _, attributeData := range data.Attributes {
			createAttributeData := service.CreateEntityAttributeData{
				EntityTemplateID: template.ID,
				Name:             attributeData.Name,
				DefaultValue:     attributeData.DefaultValue,
				ToShow:           attributeData.ToShow,
			}

			attribute, err := svc.EntityAttributeService.CreateAttribute(tx, createAttributeData)
			if err != nil {
				return err
			}

			attributes = append(attributes, *attribute)
		}

		entityTyper, err = svc.EntityTypeService.GetByID(tx, data.EntityType)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return entity.NewEntityTemplateEx(*template, nil, attributes, entityTyper), nil
}
