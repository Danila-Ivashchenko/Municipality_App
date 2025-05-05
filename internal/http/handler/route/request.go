package route

import (
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

type createRouteRequest struct {
	Name              string                   `json:"name"`
	Length            int64                    `json:"length"`
	Duration          int64                    `json:"duration"`
	Level             uint                     `json:"level"`
	MovementWay       string                   `json:"movement_way"`
	Seasonality       string                   `json:"seasonality"`
	PersonalEquipment string                   `json:"personal_equipment"`
	Dangerous         string                   `json:"dangerous"`
	Rules             string                   `json:"rules"`
	RouteEquipment    string                   `json:"route_equipment"`
	Geometry          *reqGeometry             `json:"geometry"`
	Objects           *[]setRouteObjectRequest `json:"objects,omitempty"`
}

type reqPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type reqGeometry struct {
	Points []reqPoint `json:"points"`
}

type setRouteObjectRequest struct {
	Name           string                              `json:"name"`
	OrderNumber    int                                 `json:"order_number"`
	SourceObjectID *int64                              `json:"source_object_id,omitempty"`
	Location       *createLocationToRouteObjectRequest `json:"location,omitempty"`
}

type createLocationToRouteObjectRequest struct {
	Address   *string  `json:"address,omitempty"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Geometry  *string  `json:"geometry"`
}

func (r *createLocationToRouteObjectRequest) Convert() *service.CreateObjectLocationData {
	if r == nil {
		return nil
	}

	return &service.CreateObjectLocationData{
		Address:   r.Address,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Geometry:  r.Geometry,
	}
}

func (r setRouteObjectRequest) Convert() service.SetObjectToRoute {
	return service.SetObjectToRoute{
		Name:           r.Name,
		OrderNumber:    r.OrderNumber,
		SourceObjectID: r.SourceObjectID,
		LocationData:   r.Location.Convert(),
	}
}

func (r createRouteRequest) Convert(partitionID int64) service.CreateRouteToPartitionData {
	result := service.CreateRouteToPartitionData{
		PartitionID: partitionID,
	}

	route := service.CreateRouteData{
		Name:              r.Name,
		Length:            r.Length,
		Duration:          r.Duration,
		Level:             r.Level,
		MovementWay:       r.MovementWay,
		Seasonality:       r.Seasonality,
		PersonalEquipment: r.PersonalEquipment,
		Dangerous:         r.Dangerous,
		Rules:             r.Rules,
		RouteEquipment:    r.RouteEquipment,
	}

	if r.Objects != nil {
		objects := make([]service.SetObjectToRoute, 0, len(*r.Objects))

		for _, object := range *r.Objects {
			objects = append(objects, object.Convert())
		}

		route.Objects = &objects
	}

	if r.Geometry != nil {
		geometry := entity.Geometry{}

		for _, p := range r.Geometry.Points {
			geometry.Points = append(geometry.Points, entity.Point{
				Latitude:  p.Latitude,
				Longitude: p.Longitude,
			})
		}

		route.Geometry = &geometry
	}

	result.Route = route

	return result
}

type updateRouteRequest struct {
	Name              *string                  `json:"name"`
	Length            *int64                   `json:"length"`
	Duration          *int64                   `json:"duration"`
	Level             *uint                    `json:"level"`
	MovementWay       *string                  `json:"movement_way"`
	Seasonality       *string                  `json:"seasonality"`
	PersonalEquipment *string                  `json:"personal_equipment"`
	Dangerous         *string                  `json:"dangerous"`
	Rules             *string                  `json:"rules"`
	RouteEquipment    *string                  `json:"route_equipment"`
	Geometry          *string                  `json:"geometry"`
	Objects           *[]setRouteObjectRequest `json:"objects,omitempty"`
}

func (r updateRouteRequest) Convert(routeID, partitionID int64) service.UpdateRouteToPartitionData {
	result := service.UpdateRouteToPartitionData{
		PartitionID: partitionID,
	}

	route := service.UpdateRouteData{
		ID:                routeID,
		Name:              r.Name,
		Length:            r.Length,
		Duration:          r.Duration,
		Level:             r.Level,
		MovementWay:       r.MovementWay,
		Seasonality:       r.Seasonality,
		PersonalEquipment: r.PersonalEquipment,
		Dangerous:         r.Dangerous,
		Rules:             r.Rules,
		RouteEquipment:    r.RouteEquipment,
		Geometry:          r.Geometry,
	}

	if r.Objects != nil {
		objects := make([]service.SetObjectToRoute, 0, len(*r.Objects))

		for _, object := range *r.Objects {
			objects = append(objects, object.Convert())
		}

		route.Objects = &objects
	}

	return result
}
