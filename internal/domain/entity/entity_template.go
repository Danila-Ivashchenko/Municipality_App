package entity

type EntityTemplate struct {
	ID             int64
	Name           string
	EntityTypeID   int64
	MunicipalityID int64
}

type EntityTemplateWithAttributes struct {
	EntityTemplate
	Attributes []EntityAttribute
}

type EntityTemplateEx struct {
	Template   EntityTemplate
	Attributes []EntityAttribute
	Entities   []EntityEx
	EntityType *EntityType
}

func NewEntityTemplateEx(template EntityTemplate, entityEx []EntityEx, attributes []EntityAttribute, entityType *EntityType) *EntityTemplateEx {
	result := &EntityTemplateEx{}

	result.Template = template
	result.Entities = entityEx
	result.Attributes = attributes
	result.EntityType = entityType

	return result
}
