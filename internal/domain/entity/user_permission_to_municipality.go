package entity

type UserPermissionToMunicipality struct {
	ID             string
	UserID         string
	MunicipalityID string
	UserPermission Permission
}
