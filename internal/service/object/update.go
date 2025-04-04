package object

import (
	"context"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *objectService) UpdateMultiply(ctx context.Context, data *service.UpdateMultiplyObjetsData) ([]entity.ObjectEx, error) {
	var (
		uniqueNames = make(map[string]struct{})
		names       []string

		result []entity.ObjectEx
	)

	for _, objectData := range data.Objects {
		if objectData.Name != nil {
			_, ok := uniqueNames[*objectData.Name]
			if ok {
				return nil, fmt.Errorf("duplicate object name: %s", objectData.Name)
			}

			uniqueNames[*objectData.Name] = struct{}{}
			names = append(names, *objectData.Name)
		}

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
			location *entity.Location
		)

		object, err := svc.ObjectRepository.GetByID(ctx, objectData.ID)
		if err != nil {
			return nil, err
		}

		if object == nil {
			continue
		}

		locationData := objectData.LocationData

		if locationData != nil && object.LocationID != nil {
			location, err = svc.LocationRepository.GetByID(ctx, *object.LocationID)

			if locationData.Address != nil {
				location.Address = *locationData.Address
			}

			if locationData.Latitude != nil {
				location.Latitude = *locationData.Latitude
			}

			if locationData.Longitude != nil {
				location.Longitude = *locationData.Longitude
			}

			if locationData.Geometry != nil {
				location.Geometry = locationData.Geometry
			}

			location, err = svc.LocationRepository.Update(ctx, location)
			if err != nil {
				return nil, err
			}
		}

		if objectData.Name != nil {
			object.Name = *objectData.Name
		}

		if objectData.Description != nil {
			object.Description = *objectData.Description
		}

		object, err = svc.ObjectRepository.Update(ctx, object)
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
