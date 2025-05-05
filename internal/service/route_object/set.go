package route_object

import (
	"context"
	"errors"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
	"sort"
)

func (svc *routeObjectService) SetToRoute(ctx context.Context, data *service.SetObjectsToRoute) ([]entity.RouteObjectEx, error) {
	var (
		result          []entity.RouteObjectEx
		currentOrder    = 1
		objectsToCreate []entity.RouteObject
	)

	err := svc.RouteObjectRepository.DeleteToRoute(ctx, data.RouteID)
	if err != nil {
		return nil, err
	}

	sort.Slice(data.Objects, func(i, j int) bool {
		return data.Objects[i].OrderNumber < data.Objects[j].OrderNumber
	})

	for _, object := range data.Objects {
		var (
			locationCreateData *repository.CreateLocationData
			location           *entity.Location

			locationID *int64
		)

		if validationErr := object.Validate(); validationErr != nil {
			return nil, validationErr
		}

		if object.OrderNumber != currentOrder {
			return nil, errors.New(fmt.Sprintf("invalid order number: %d for object with name: %s", object.OrderNumber, object.Name))
		}

		currentOrder++

		if object.SourceObjectID != nil {
			objectEx, err := svc.ObjectService.GetExByID(ctx, *object.SourceObjectID)
			if err != nil {
				return nil, err
			}

			if objectEx == nil {
				return nil, errors.New("object does not exist")
			}
		}

		if object.LocationData != nil {
			locationCreateData = &repository.CreateLocationData{
				Address:   object.LocationData.Address,
				Latitude:  object.LocationData.Latitude,
				Longitude: object.LocationData.Longitude,
				Geometry:  object.LocationData.Geometry,
			}

			location, err = svc.LocationRepository.Create(ctx, locationCreateData)
			if err != nil {
				return nil, err
			}

			locationID = &location.ID
		}

		routeObject := entity.RouteObject{
			RouteID:        data.RouteID,
			Name:           object.Name,
			OrderNumber:    object.OrderNumber,
			SourceObjectID: object.SourceObjectID,
			LocationID:     locationID,
		}

		objectsToCreate = append(objectsToCreate, routeObject)
	}

	for _, object := range objectsToCreate {
		resultObject, err := svc.RouteObjectRepository.Create(ctx, &object)
		if err != nil {
			return nil, err
		}

		resultObjectEx, err := svc.GetExByID(ctx, resultObject.ID)
		if err != nil {
			return nil, err
		}

		if resultObjectEx == nil {
			continue
		}

		result = append(result, *resultObjectEx)
	}

	return result, nil
}
