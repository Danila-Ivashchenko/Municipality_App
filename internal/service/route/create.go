package route

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *routeService) CreateToPartition(ctx context.Context, data *service.CreateRouteToPartitionData) (*entity.RouteEx, error) {
	var (
		routeObjects []entity.RouteObjectEx
		routeCreate  *entity.Route
	)

	routeExists, err := svc.RouteRepository.GetByNamePartitionID(ctx, data.Route.Name, data.PartitionID)
	if err != nil {
		return nil, err
	}

	if routeExists != nil {
		return nil, errors.New("route already exists")
	}

	routeToCreate := &entity.Route{
		Name:              data.Route.Name,
		PartitionID:       data.PartitionID,
		Length:            data.Route.Length,
		Duration:          data.Route.Duration,
		Level:             data.Route.Level,
		MovementWay:       data.Route.MovementWay,
		Seasonality:       data.Route.Seasonality,
		PersonalEquipment: data.Route.PersonalEquipment,
		Dangerous:         data.Route.Dangerous,
		Rules:             data.Route.Rules,
		RouteEquipment:    data.Route.RouteEquipment,
		Geometry:          data.Route.Geometry,
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		routeCreate, err = svc.RouteRepository.Create(ctx, routeToCreate)
		if err != nil {
			return err
		}

		if data.Route.Objects != nil {
			setObjectsData := &service.SetObjectsToRoute{
				RouteID: routeCreate.ID,
				Objects: *data.Route.Objects,
			}

			routeObjects, err = svc.RouteObjectService.SetToRoute(ctx, setObjectsData)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	result := entity.NewRouteEx(*routeCreate, routeObjects)

	return &result, nil
}
