package user_admin

import (
	"context"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *userAdminService) Update(ctx context.Context, data *service.UpdateUserByAdminData) (*entity.UserEx, error) {
	var (
		userUpdated        bool
		updateUserCoreData = &service.UpdateUserData{
			ID: data.ID,
		}
	)

	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := svc.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, core_errors.UserNotFound
	}

	if data.Name != nil && user.User.Name != *data.Name {
		userUpdated = true
		updateUserCoreData.Name = data.Name
	}

	if data.LastName != nil && user.User.LastName != *data.LastName {
		userUpdated = true
		updateUserCoreData.LastName = data.LastName
	}

	if data.Email != nil && user.User.Email != *data.Email {
		userUpdated = true
		updateUserCoreData.Email = data.Email
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		if data.Permissions != nil {
			setPermissionsData := &service.SetUserPermissionsData{
				UserID:      data.ID,
				Permissions: *data.Permissions,
			}

			_, err = svc.UserPermissionService.SetPermissionsToUser(tx, setPermissionsData)
			if err != nil {
				return err
			}
		}

		if data.IsAdmin != nil {
			err = svc.UserService.ChangeUserIsAdmin(tx, data.ID, *data.IsAdmin)
			if err != nil {
				return err
			}
		}

		if data.IsBlocked != nil {
			err = svc.UserService.BlockUserByID(tx, data.ID, *data.IsBlocked)
		}

		if userUpdated {
			_, err = svc.UserService.UpdateUser(tx, updateUserCoreData)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return svc.GetByID(ctx, data.ID)
}
