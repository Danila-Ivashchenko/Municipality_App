package view

import "municipality_app/internal/domain/entity"

type RouteView struct {
	ID                int64             `json:"id"`
	PartitionID       int64             `json:"partition_id"`
	Name              string            `json:"name"`
	Length            int64             `json:"length"`
	Duration          int64             `json:"duration"`
	Level             uint              `json:"level"`
	MovementWay       string            `json:"movement_way"`
	Seasonality       string            `json:"seasonality"`
	PersonalEquipment string            `json:"personal_equipment"`
	Dangerous         string            `json:"dangerous"`
	Rules             string            `json:"rules"`
	RouteEquipment    string            `json:"route_equipment"`
	Geometry          *GeometryView     `json:"geometry"`
	Objects           []RouteObjectView `json:"objects"`
}

func NewRouteView(route *entity.RouteEx) *RouteView {
	var (
		geometryView *GeometryView
	)

	if route.Geometry != nil {
		geometryVal := GeometryView{}

		for _, p := range route.Geometry.Points {
			geometryVal.Points = append(geometryVal.Points, PointView{
				Latitude:  p.Latitude,
				Longitude: p.Longitude,
			})
		}

		geometryView = &geometryVal
	}

	objects := make([]RouteObjectView, 0)

	for _, o := range route.RouteObjects {
		objects = append(objects, *NewRouteObjectView(o))
	}

	return &RouteView{
		ID:                route.ID,
		PartitionID:       route.PartitionID,
		Name:              route.Name,
		Length:            route.Length,
		Duration:          route.Duration,
		Level:             route.Level,
		MovementWay:       route.MovementWay,
		Seasonality:       route.Seasonality,
		PersonalEquipment: route.PersonalEquipment,
		Dangerous:         route.Dangerous,
		Rules:             route.Rules,
		RouteEquipment:    route.RouteEquipment,
		Geometry:          geometryView,
		Objects:           objects,
	}
}

type PointView struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeometryView struct {
	Points []PointView `json:"points"`
}

type RouteObjectView struct {
	ID           int64               `json:"id"`
	Name         string              `json:"name"`
	RouteID      int64               `json:"route_id"`
	OrderNumber  int                 `json:"order_number"`
	SourceObject *ObjectView         `json:"source_object"`
	Location     *ObjectLocationView `json:"location"`
}

func NewRouteObjectView(routeObject entity.RouteObjectEx) *RouteObjectView {
	var (
		objectView   *ObjectView
		locationView *ObjectLocationView
	)

	if routeObject.SourceObject != nil {
		objectViewValue := NewObjectView(routeObject.SourceObject)
		objectView = &objectViewValue
	}

	if routeObject.LocationID != nil {
		locationView = NewObjectLocationView(routeObject.LocationID)
	}

	return &RouteObjectView{
		ID:           routeObject.ID,
		Name:         routeObject.Name,
		RouteID:      routeObject.RouteID,
		OrderNumber:  routeObject.OrderNumber,
		SourceObject: objectView,
		Location:     locationView,
	}
}
