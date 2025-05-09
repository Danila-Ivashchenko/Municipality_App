package entity

type ObjectTemplate struct {
	ID             int64
	Name           string
	ObjectTypeID   int64
	MunicipalityID int64
}

type ObjectTemplateWithAttributes struct {
	ObjectTemplate
	Attributes []ObjectAttribute
}

type ObjectTemplateEx struct {
	Template   ObjectTemplate
	Attributes []ObjectAttribute
	Objects    []ObjectEx
	ObjectType *ObjectType
}

func NewObjectTemplateEx(template ObjectTemplate, objectsEx []ObjectEx, attributes []ObjectAttribute, objectType *ObjectType) *ObjectTemplateEx {
	result := &ObjectTemplateEx{}

	result.Template = template
	result.Objects = objectsEx
	result.Attributes = attributes
	result.ObjectType = objectType

	return result
}
