package entity

type RouteObject struct {
	ID             int64
	Name           string
	RouteID        int64
	OrderNumber    int
	SourceObjectID *int64
	LocationID     *int64
}

type RouteObjectEx struct {
	ID           int64
	Name         string
	RouteID      int64
	OrderNumber  int
	SourceObject *ObjectEx
	LocationID   *Location
}

func NewRouteObjectEx(routeObject *RouteObject, object *ObjectEx, location *Location) RouteObjectEx {
	return RouteObjectEx{
		ID:           routeObject.ID,
		Name:         routeObject.Name,
		RouteID:      routeObject.RouteID,
		OrderNumber:  routeObject.OrderNumber,
		SourceObject: object,
		LocationID:   location,
	}
}
