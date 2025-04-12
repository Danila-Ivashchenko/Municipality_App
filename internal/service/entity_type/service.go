package entity_type

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *entityTypeService) Create(ctx context.Context, data *service.CreateEntityTypeData) (*entity.EntityType, error) {
	entityTypeExist, err := svc.GetByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}

	if entityTypeExist != nil {
		return nil, errors.New("entity type already exist")
	}

	repoData := &repository.CreateEntityType{Name: data.Name}

	return svc.EntityTypeRepository.Create(ctx, repoData)
}

func (svc *entityTypeService) Update(ctx context.Context, data *service.UpdateEntityTypeData) (*entity.EntityType, error) {
	entityType, err := svc.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if entityType == nil {
		return nil, errors.New("entity type does not exist")
	}

	entityTypeExist, err := svc.GetByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}

	if entityTypeExist != nil {
		return nil, errors.New("entity type already exist")
	}

	entityType.Name = data.Name

	err = svc.EntityTypeRepository.Update(ctx, entityType)

	return entityType, err
}

func (svc *entityTypeService) CreateMultiply(ctx context.Context, data *service.CreateEntityTypeMultiplyData) ([]entity.EntityType, error) {
	var (
		namesMap   = make(map[string]struct{})
		names      []string
		result     []entity.EntityType
		entityType *entity.EntityType
	)

	for _, d := range data.Data {
		if _, ok := namesMap[d.Name]; ok {
			return nil, errors.New("duplicate name")
		}

		names = append(names, d.Name)
		namesMap[d.Name] = struct{}{}
	}

	entitysExists, err := svc.GetByNames(ctx, names)
	if err != nil {
		return nil, err
	}

	if len(entitysExists) > 0 {
		return nil, errors.New("duplicate name")
	}

	for _, d := range data.Data {
		repoData := &repository.CreateEntityType{Name: d.Name}

		entityType, err = svc.EntityTypeRepository.Create(ctx, repoData)
		if err != nil {
			return nil, err
		}

		result = append(result, *entityType)
	}

	return result, nil
}

func (svc *entityTypeService) Delete(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		err := svc.EntityTypeRepository.DeleteByID(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *entityTypeService) GetAll(ctx context.Context) ([]entity.EntityType, error) {
	return svc.EntityTypeRepository.GetAll(ctx)
}

func (svc *entityTypeService) GetByName(ctx context.Context, name string) (*entity.EntityType, error) {
	return svc.EntityTypeRepository.GetByName(ctx, name)
}

func (svc *entityTypeService) GetByNames(ctx context.Context, names []string) ([]entity.EntityType, error) {
	if len(names) == 0 {
		return nil, nil
	}

	return svc.EntityTypeRepository.GetByNames(ctx, names)
}

func (svc *entityTypeService) GetByID(ctx context.Context, id int64) (*entity.EntityType, error) {
	return svc.EntityTypeRepository.GetByID(ctx, id)
}

func (svc *entityTypeService) GetByIDs(ctx context.Context, ids []int64) ([]entity.EntityType, error) {
	return svc.EntityTypeRepository.GetByIDs(ctx, ids)
}
