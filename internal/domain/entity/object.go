package entity

type Object struct {
	ID               int64
	Name             string
	ObjectTemplateID int64
	LocationID       *int64
	Description      string
}

type ObjectEx struct {
	ID               int64
	Name             string
	ObjectTemplateID int64
	LocationID       *Location
	Description      string
	AttributeValues  []ObjectAttributeValueEx
}

func NewObjectExPtr(object *Object, location *Location, attributeValues []ObjectAttributeValueEx) *ObjectEx {
	return &ObjectEx{
		ID:               object.ID,
		Name:             object.Name,
		ObjectTemplateID: object.ObjectTemplateID,
		LocationID:       location,
		Description:      object.Description,
		AttributeValues:  attributeValues,
	}
}

func NewObjectEx(object Object, location Location, attributeValues []ObjectAttributeValueEx) *ObjectEx {
	return &ObjectEx{
		ID:               object.ID,
		Name:             object.Name,
		ObjectTemplateID: object.ObjectTemplateID,
		LocationID:       &location,
		Description:      object.Description,
		AttributeValues:  attributeValues,
	}
}
