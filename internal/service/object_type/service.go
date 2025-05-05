package object_type

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *objectTypeService) Create(ctx context.Context, data *service.CreateObjectTypeData) (*entity.ObjectType, error) {
	objectTypeExist, err := svc.GetByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}

	if objectTypeExist != nil {
		return nil, errors.New("object type already exist")
	}

	repoData := &repository.CreateObjectType{Name: data.Name}

	return svc.ObjectTypeRepository.Create(ctx, repoData)
}

func (svc *objectTypeService) Update(ctx context.Context, data *service.UpdateObjectTypeData) (*entity.ObjectType, error) {
	objectType, err := svc.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if objectType == nil {
		return nil, errors.New("object type does not exist")
	}

	if objectType.Name == data.Name {
		return objectType, nil
	}

	objectTypeExist, err := svc.GetByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}

	if objectTypeExist != nil {
		return nil, errors.New("object type already exist")
	}

	objectType.Name = data.Name

	err = svc.ObjectTypeRepository.Update(ctx, objectType)

	return objectType, err
}

func (svc *objectTypeService) CreateMultiply(ctx context.Context, data *service.CreateObjectTypeMultiplyData) ([]entity.ObjectType, error) {
	var (
		namesMap   = make(map[string]struct{})
		names      []string
		result     []entity.ObjectType
		objectType *entity.ObjectType
	)

	for _, d := range data.Data {
		if _, ok := namesMap[d.Name]; ok {
			return nil, errors.New("duplicate name")
		}

		names = append(names, d.Name)
		namesMap[d.Name] = struct{}{}
	}

	objectsExists, err := svc.GetByNames(ctx, names)
	if err != nil {
		return nil, err
	}

	if len(objectsExists) > 0 {
		return nil, errors.New("duplicate name")
	}

	for _, d := range data.Data {
		repoData := &repository.CreateObjectType{Name: d.Name}

		objectType, err = svc.ObjectTypeRepository.Create(ctx, repoData)
		if err != nil {
			return nil, err
		}

		result = append(result, *objectType)
	}

	return result, nil
}

func (svc *objectTypeService) Delete(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		objectTeplates, err := svc.ObjectTemplateRepository.GetByTypeID(ctx, id)
		if err != nil {
			return err
		}

		if len(objectTeplates) > 0 {
			return errors.New("object type is used")
		}

		err = svc.ObjectTypeRepository.DeleteByID(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *objectTypeService) GetAll(ctx context.Context) ([]entity.ObjectType, error) {
	return svc.ObjectTypeRepository.GetAll(ctx)
}

func (svc *objectTypeService) GetByName(ctx context.Context, name string) (*entity.ObjectType, error) {
	return svc.ObjectTypeRepository.GetByName(ctx, name)
}

func (svc *objectTypeService) GetByNames(ctx context.Context, names []string) ([]entity.ObjectType, error) {
	if len(names) == 0 {
		return nil, nil
	}

	return svc.ObjectTypeRepository.GetByNames(ctx, names)
}

func (svc *objectTypeService) GetByID(ctx context.Context, id int64) (*entity.ObjectType, error) {
	return svc.ObjectTypeRepository.GetByID(ctx, id)
}

func (svc *objectTypeService) GetByIDs(ctx context.Context, ids []int64) ([]entity.ObjectType, error) {
	return svc.ObjectTypeRepository.GetByIDs(ctx, ids)
}
