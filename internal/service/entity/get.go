package entity

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *entityService) GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Entity, error) {
	return svc.EntityRepository.GetByTemplateID(ctx, templateID)
}

func (svc *entityService) GetByIDs(ctx context.Context, ids []int64) ([]entity.Entity, error) {
	return svc.EntityRepository.GetByIDs(ctx, ids)
}

func (svc *entityService) GetByID(ctx context.Context, id int64) (*entity.Entity, error) {
	return svc.EntityRepository.GetByID(ctx, id)
}

func (svc *entityService) GetByNamesAndTemplateID(ctx context.Context, names []string, templateID int64) ([]entity.Entity, error) {
	return svc.EntityRepository.GetByTemplateIDAndNames(ctx, names, templateID)
}

func (svc *entityService) GetExByIDs(ctx context.Context, ids []int64) ([]entity.EntityEx, error) {
	var (
		entityByID = make(map[int64]entity.Entity)
		result     []entity.EntityEx
	)

	entities, err := svc.EntityRepository.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, obj := range entities {
		entityByID[obj.ID] = obj
	}

	for _, obj := range entities {
		attributeValues, err := svc.EntityAttributeService.GetAttributesExByEntityID(ctx, obj.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *entity.NewEntityExPtr(&obj, attributeValues))
	}

	return result, nil
}

func (svc *entityService) GetExByTemplateID(ctx context.Context, templateID int64) ([]entity.EntityEx, error) {
	var (
		entityIDs []int64
	)

	entities, err := svc.EntityRepository.GetByTemplateID(ctx, templateID)
	if err != nil {
		return nil, err
	}

	for _, obj := range entities {
		entityIDs = append(entityIDs, obj.ID)
	}

	return svc.GetExByIDs(ctx, entityIDs)
}
