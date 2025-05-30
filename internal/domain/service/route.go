package service

import (
	"context"
	"municipality_app/internal/common/validator/field"
	"municipality_app/internal/common/validator/validator"
	"municipality_app/internal/domain/entity"
)

type RouteService interface {
	CreateToPartition(ctx context.Context, data *CreateRouteToPartitionData) (*entity.RouteEx, error)
	UpdateToPartition(ctx context.Context, data *UpdateRouteToPartitionData) (*entity.RouteEx, error)

	GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.RouteEx, error)
	GetByIDAndPartitionID(ctx context.Context, id, partitionID int64) (*entity.Route, error)

	DeleteToPartition(ctx context.Context, data *DeleteRoutesToPartitionData) error
}

type CreateRouteToPartitionData struct {
	PartitionID int64
	Route       CreateRouteData
}

type CreateRouteData struct {
	Name              string
	Length            int64
	Duration          int64
	Level             uint
	MovementWay       string
	Seasonality       string
	PersonalEquipment string
	Dangerous         string
	Rules             string
	RouteEquipment    string
	Geometry          *entity.Geometry
	Objects           *[]SetObjectToRoute
}

func (d *CreateRouteData) Validate() error {
	v := validator.Validator{}

	v.AddField(
		field.NewStringField("Название", d.Name).Required().Bigger(4),
	)

	return v.Validate()
}

type UpdateRouteData struct {
	ID                int64
	Name              *string
	Length            *int64
	Duration          *int64
	Level             *uint
	MovementWay       *string
	Seasonality       *string
	PersonalEquipment *string
	Dangerous         *string
	Rules             *string
	RouteEquipment    *string
	Geometry          *string
	Objects           *[]SetObjectToRoute
}

func (d *UpdateRouteData) Validate() error {
	v := validator.Validator{}

	if d.Name != nil {
		v.AddField(
			field.NewStringField("Название", *d.Name).Required().Bigger(4),
		)
	}

	return v.Validate()
}

type DeleteRoutesToPartitionData struct {
	PartitionID int64
	RoutesID    int64
}

type UpdateRouteToPartitionData struct {
	PartitionID int64
	Route       UpdateRouteData
}
