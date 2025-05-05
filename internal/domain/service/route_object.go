package service

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
)

type RouteObjectService interface {
	SetToRoute(ctx context.Context, data *SetObjectsToRoute) ([]entity.RouteObjectEx, error)

	GetByRouteID(ctx context.Context, id int64) ([]entity.RouteObjectEx, error)
}

type SetObjectsToRoute struct {
	RouteID int64
	Objects []SetObjectToRoute
}

type SetObjectToRoute struct {
	Name           string                    `json:"name"`
	OrderNumber    int                       `json:"order_number"`
	SourceObjectID *int64                    `json:"source_object_id"`
	LocationData   *CreateObjectLocationData `json:"location"`
}

func (d SetObjectToRoute) Validate() error {
	if len(d.Name) == 0 {
		return errors.New("name is required")
	}

	if d.OrderNumber < 0 {
		return errors.New("order umber is required")
	}

	return nil
}
