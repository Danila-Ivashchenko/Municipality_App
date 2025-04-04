package entity

type ObjectAttribute struct {
	ID               int64
	ObjectTemplateID int64
	Name             string
	DefaultValue     string
	ToShow           bool
}
