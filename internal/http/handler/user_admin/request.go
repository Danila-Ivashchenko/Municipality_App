package user_admin

import (
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

type reqUpdateUser struct {
	ID        int64   `json:"id"`
	Name      *string `json:"name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	IsAdmin   *bool   `json:"is_admin"`
	IsBlocked *bool   `json:"is_blocked"`

	Permissions *[]uint8 `json:"permissions"`
}

func (r *reqUpdateUser) Convert() *service.UpdateUserByAdminData {
	var (
		permissionArr *[]entity.Permission
	)

	if r.Permissions != nil {
		permissionVal := make([]entity.Permission, 0)

		for _, v := range *r.Permissions {
			permissionVal = append(permissionVal, entity.PermissionFromUint8(v))
		}

		permissionArr = &permissionVal
	}

	return &service.UpdateUserByAdminData{
		ID:          r.ID,
		Name:        r.Name,
		LastName:    r.LastName,
		Email:       r.Email,
		IsAdmin:     r.IsAdmin,
		IsBlocked:   r.IsBlocked,
		Permissions: permissionArr,
	}
}
