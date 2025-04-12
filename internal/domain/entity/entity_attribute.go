package entity

type EntityAttribute struct {
	ID               int64
	EntityTemplateID int64
	Name             string
	DefaultValue     string
	ToShow           bool
}
