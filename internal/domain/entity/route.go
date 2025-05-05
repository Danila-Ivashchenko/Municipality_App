package entity

type Route struct {
	ID                int64
	PartitionID       int64
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
	Geometry          *Geometry
}

type RouteEx struct {
	Route
	RouteObjects []RouteObjectEx
}

type Point struct {
	Latitude  float64
	Longitude float64
}

type Geometry struct {
	Points []Point
}

func NewRouteEx(r Route, objects []RouteObjectEx) RouteEx {
	return RouteEx{
		Route:        r,
		RouteObjects: objects,
	}
}
