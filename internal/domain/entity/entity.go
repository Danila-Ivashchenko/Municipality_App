package entity

type Entity struct {
	ID               int64
	Name             string
	EntityTemplateID int64
	Description      string
}

type EntityEx struct {
	ID               int64
	Name             string
	EntityTemplateID int64
	Description      string
	AttributeValues  []EntityAttributeValueEx
}

func NewEntityExPtr(entity *Entity, attributeValues []EntityAttributeValueEx) *EntityEx {
	return &EntityEx{
		ID:               entity.ID,
		Name:             entity.Name,
		EntityTemplateID: entity.EntityTemplateID,
		Description:      entity.Description,
		AttributeValues:  attributeValues,
	}
}

func NewEntityEx(entity Entity, attributeValues []EntityAttributeValueEx) *EntityEx {
	return &EntityEx{
		ID:               entity.ID,
		Name:             entity.Name,
		EntityTemplateID: entity.EntityTemplateID,
		Description:      entity.Description,
		AttributeValues:  attributeValues,
	}
}
