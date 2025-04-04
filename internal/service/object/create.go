package object

import (
	"context"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *objectService) Create(ctx context.Context, data *service.CreateObjectData) (*entity.Object, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *objectService) CreateMultiply(ctx context.Context, data *service.CreateMultiplyObjetsData) ([]entity.ObjectEx, error) {
	var (
		uniqueNames = make(map[string]struct{})
		names       []string

		result []entity.ObjectEx
	)

	for _, objectData := range data.Objects {
		_, ok := uniqueNames[objectData.Name]
		if ok {
			return nil, fmt.Errorf("duplicate object name: %s", objectData.Name)
		}

		uniqueNames[objectData.Name] = struct{}{}
		names = append(names, objectData.Name)
	}

	objectExists, err := svc.GetByNamesAndTemplateID(ctx, names, data.ObjectTemplateID)
	if err != nil {
		return nil, err
	}

	if len(objectExists) > 0 {
		return nil, fmt.Errorf("duplicate object name: %s", objectExists[0].Name)
	}

	for _, objectData := range data.Objects {
		var (
			locationCreateData *repository.CreateLocationData
			location           *entity.Location

			locationID *int64

			objectRepoData *repository.CreateObjectData
		)

		if objectData.LocationData != nil {
			locationCreateData = &repository.CreateLocationData{
				Address:   objectData.LocationData.Address,
				Latitude:  objectData.LocationData.Latitude,
				Longitude: objectData.LocationData.Longitude,
				Geometry:  objectData.LocationData.Geometry,
			}

			location, err = svc.LocationRepository.Create(ctx, locationCreateData)
			if err != nil {
				return nil, err
			}

			locationID = &location.ID
		}

		objectRepoData = &repository.CreateObjectData{
			Name:             objectData.Name,
			LocationID:       locationID,
			ObjectTemplateID: data.ObjectTemplateID,
			Description:      objectData.Description,
		}

		object, err := svc.ObjectRepository.Create(ctx, objectRepoData)
		if err != nil {
			return nil, err
		}

		createAttributeValuesData := service.CreateObjectAttributesData{
			ObjectID:         object.ID,
			ObjectTemplateID: object.ObjectTemplateID,
			ValuesData:       objectData.AttributeValues,
		}

		_, err = svc.ObjectAttributeService.UpdateValues(ctx, createAttributeValuesData)
		if err != nil {
			return nil, err
		}

		attributeValues, err := svc.ObjectAttributeService.GetAttributesExByObjectID(ctx, object.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *entity.NewObjectExPtr(object, location, attributeValues))
	}

	return result, nil
}
