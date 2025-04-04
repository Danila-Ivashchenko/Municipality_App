package entity

type ObjectAttributeValue struct {
	ID                int64
	ObjectAttributeID int64
	ObjectID          int64
	Value             string
}

type ObjectAttributeValueEx struct {
	Attribute ObjectAttribute
	Value     ObjectAttributeValue
}

func NewObjectAttributeValueEx(attribute ObjectAttribute, value ObjectAttributeValue) ObjectAttributeValueEx {
	return ObjectAttributeValueEx{Attribute: attribute, Value: value}
}
