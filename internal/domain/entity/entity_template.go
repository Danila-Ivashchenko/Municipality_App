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
}

func NewEntityTemplateEx(template EntityTemplate, entityEx []EntityEx, attributes []EntityAttribute) *EntityTemplateEx {
	result := &EntityTemplateEx{}

	result.Template = template
	result.Entities = entityEx
	result.Attributes = attributes

	return result
}
