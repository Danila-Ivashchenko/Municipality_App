package route_object

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *routeObjectService) GetByRouteID(ctx context.Context, id int64) ([]entity.RouteObjectEx, error) {
	var (
		result []entity.RouteObjectEx
	)

	routeObjects, err := svc.RouteObjectRepository.GetByRouteID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, routeObject := range routeObjects {
		routeObjectEx, err := svc.GetExByID(ctx, routeObject.ID)
		if err != nil {
			return nil, err
		}

		if routeObjectEx == nil {
			continue
		}

		result = append(result, *routeObjectEx)
	}

	return result, nil
}

func (svc *routeObjectService) GetExByID(ctx context.Context, id int64) (*entity.RouteObjectEx, error) {
	var (
		result entity.RouteObjectEx

		objectEx *entity.ObjectEx
		location *entity.Location
		err      error
	)

	routeObject, err := svc.RouteObjectRepository.GetID(ctx, id)
	if err != nil {
		return nil, err
	}

	if routeObject == nil {
		return nil, nil
	}

	if routeObject.SourceObjectID != nil {
		objectEx, err = svc.ObjectService.GetExByID(ctx, *routeObject.SourceObjectID)
		if err != nil {
			return nil, err
		}
	}

	if routeObject.LocationID != nil {
		location, err = svc.LocationRepository.GetByID(ctx, *routeObject.LocationID)
		if err != nil {
			return nil, err
		}
	}

	result = entity.NewRouteObjectEx(routeObject, objectEx, location)

	return &result, nil
}
