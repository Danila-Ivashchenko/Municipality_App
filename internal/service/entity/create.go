package entity

import (
	"context"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *entityService) Create(ctx context.Context, data *service.CreateEntityData) (*entity.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *entityService) CreateMultiply(ctx context.Context, data *service.CreateMultiplyEntitiesData) ([]entity.EntityEx, error) {
	var (
		uniqueNames = make(map[string]struct{})
		names       []string

		result []entity.EntityEx
	)

	for _, entityData := range data.Entities {
		_, ok := uniqueNames[entityData.Name]
		if ok {
			return nil, fmt.Errorf("duplicate entity name: %s", entityData.Name)
		}

		uniqueNames[entityData.Name] = struct{}{}
		names = append(names, entityData.Name)
	}

	entityExists, err := svc.GetByNamesAndTemplateID(ctx, names, data.EntityTemplateID)
	if err != nil {
		return nil, err
	}

	if len(entityExists) > 0 {
		return nil, fmt.Errorf("duplicate entity name: %s", entityExists[0].Name)
	}

	for _, entityData := range data.Entities {
		var (
			entityRepoData *repository.CreateEntityData
		)

		entityRepoData = &repository.CreateEntityData{
			Name:             entityData.Name,
			EntityTemplateID: data.EntityTemplateID,
			Description:      entityData.Description,
		}

		e, err := svc.EntityRepository.Create(ctx, entityRepoData)
		if err != nil {
			return nil, err
		}

		createAttributeValuesData := service.CreateEntityAttributesData{
			EntityID:         e.ID,
			EntityTemplateID: e.EntityTemplateID,
			ValuesData:       entityData.AttributeValues,
		}

		_, err = svc.EntityAttributeService.UpdateValues(ctx, createAttributeValuesData)
		if err != nil {
			return nil, err
		}

		attributeValues, err := svc.EntityAttributeService.GetAttributesExByEntityID(ctx, e.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *entity.NewEntityExPtr(e, attributeValues))
	}

	return result, nil
}
