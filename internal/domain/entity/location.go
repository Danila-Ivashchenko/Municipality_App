package entity

type Location struct {
	ID        int64
	Address   string
	Latitude  float64
	Longitude float64
	Geometry  *string
}
