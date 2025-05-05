package route

import (
	"context"
	"errors"
	"municipality_app/internal/domain/service"
)

func (svc *routeService) DeleteToPartition(ctx context.Context, data *service.DeleteRoutesToPartitionData) error {
	routeExists, err := svc.RouteRepository.GetByIDPartitionID(ctx, data.RoutesID, data.PartitionID)
	if err != nil {
		return err
	}

	if routeExists == nil {
		return errors.New("route not found")
	}

	return svc.RouteRepository.Delete(ctx, routeExists.ID)
}
