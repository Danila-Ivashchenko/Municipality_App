package route

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *routeService) UpdateToPartition(ctx context.Context, data *service.UpdateRouteToPartitionData) (*entity.RouteEx, error) {
	var (
		routeObjects []entity.RouteObjectEx
	)

	route, err := svc.RouteRepository.GetByID(ctx, data.Route.ID)
	if err != nil {
		return nil, err
	}

	if route == nil {
		return nil, errors.New("route does not exist")
	}
	if data.Route.Name != nil && *data.Route.Name != route.Name {
		routeByName, err := svc.RouteRepository.GetByNamePartitionID(ctx, *data.Route.Name, data.PartitionID)
		if err != nil {
			return nil, err
		}

		if routeByName != nil {
			return nil, errors.New("route already exists")
		}

	}

	if data.Route.Name != nil {
		route.Name = *data.Route.Name
	}

	if data.Route.Length != nil {
		route.Length = *data.Route.Length
	}

	if data.Route.Duration != nil {
		route.Duration = *data.Route.Duration
	}

	if data.Route.Level != nil {
		route.Level = *data.Route.Level
	}

	if data.Route.MovementWay != nil {
		route.MovementWay = *data.Route.MovementWay
	}

	if data.Route.Seasonality != nil {
		route.Seasonality = *data.Route.Seasonality
	}

	if data.Route.PersonalEquipment != nil {
		route.PersonalEquipment = *data.Route.PersonalEquipment
	}

	if data.Route.Dangerous != nil {
		route.Dangerous = *data.Route.Dangerous
	}

	if data.Route.Rules != nil {
		route.Rules = *data.Route.Rules
	}

	if data.Route.RouteEquipment != nil {
		route.RouteEquipment = *data.Route.RouteEquipment
	}

	//if data.Route.Geometry != nil {
	//	route.Geometry = *data.Route.Geometry
	//}

	if data.Route.RouteEquipment != nil {
		route.RouteEquipment = *data.Route.RouteEquipment
	}

	routeUpdated, err := svc.RouteRepository.Update(ctx, route)
	if err != nil {
		return nil, err
	}

	if data.Route.Objects != nil {
		setObjectsData := &service.SetObjectsToRoute{
			RouteID: routeUpdated.ID,
			Objects: *data.Route.Objects,
		}

		routeObjects, err = svc.RouteObjectService.SetToRoute(ctx, setObjectsData)
		if err != nil {
			return nil, err
		}
	}

	result := entity.NewRouteEx(*routeUpdated, routeObjects)

	return &result, nil
}
