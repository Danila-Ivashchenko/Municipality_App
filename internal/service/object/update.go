package object

import (
	"context"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
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
				return nil, core_errors.ObjectNameIsUsed
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
		return nil, core_errors.ObjectNameIsUsed
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		for _, objectData := range data.Objects {
			var (
				location *entity.Location
			)

			object, err := svc.ObjectRepository.GetByID(tx, objectData.ID)
			if err != nil {
				return err
			}

			if object == nil {
				continue
			}

			locationData := objectData.LocationData

			if locationData != nil {
				if object.LocationID != nil {
					location, err = svc.LocationRepository.GetByID(tx, *object.LocationID)

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

					location, err = svc.LocationRepository.Update(tx, location)
					if err != nil {
						return err
					}
				} else {
					locationCreateData := &repository.CreateLocationData{
						Address:   objectData.LocationData.Address,
						Latitude:  objectData.LocationData.Latitude,
						Longitude: objectData.LocationData.Longitude,
						Geometry:  objectData.LocationData.Geometry,
					}

					location, err = svc.LocationRepository.Create(tx, locationCreateData)
					if err != nil {
						return err
					}

					object.LocationID = &location.ID
				}
			}

			if objectData.Name != nil {
				object.Name = *objectData.Name
			}

			if objectData.Description != nil {
				object.Description = *objectData.Description
			}

			object, err = svc.ObjectRepository.Update(tx, object)
			if err != nil {
				return err
			}

			createAttributeValuesData := service.CreateObjectAttributesData{
				ObjectID:         object.ID,
				ObjectTemplateID: object.ObjectTemplateID,
				ValuesData:       objectData.AttributeValues,
			}

			_, err = svc.ObjectAttributeService.UpdateValues(tx, createAttributeValuesData)
			if err != nil {
				return err
			}

			attributeValues, err := svc.ObjectAttributeService.GetAttributesExByObjectID(tx, object.ID)
			if err != nil {
				return err
			}

			result = append(result, *entity.NewObjectExPtr(object, location, attributeValues))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
