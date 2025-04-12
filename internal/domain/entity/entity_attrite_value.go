package entity

type EntityAttributeValue struct {
	ID                int64
	EntityAttributeID int64
	EntityID          int64
	Value             string
}

type EntityAttributeValueEx struct {
	Attribute EntityAttribute
	Value     EntityAttributeValue
}

func NewEntityAttributeValueEx(attribute EntityAttribute, value EntityAttributeValue) EntityAttributeValueEx {
	return EntityAttributeValueEx{Attribute: attribute, Value: value}
}
