package route

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *routeService) GetExByID(ctx context.Context, id int64) (*entity.RouteEx, error) {
	route, err := svc.RouteRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if route == nil {
		return nil, nil
	}

	routeObjects, err := svc.RouteObjectService.GetByRouteID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := entity.NewRouteEx(*route, routeObjects)

	return &result, nil
}

func (svc *routeService) GetByIDAndPartitionID(ctx context.Context, id, partitionID int64) (*entity.Route, error) {
	return svc.RouteRepository.GetByIDPartitionID(ctx, id, partitionID)
}

func (svc *routeService) GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.RouteEx, error) {
	var (
		result []entity.RouteEx
	)

	routes, err := svc.RouteRepository.GetByPartitionID(ctx, partitionID)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		routeEx, err := svc.GetExByID(ctx, route.ID)
		if err != nil {
			return nil, err
		}

		if routeEx == nil {
			continue
		}

		result = append(result, *routeEx)
	}

	return result, nil
}
