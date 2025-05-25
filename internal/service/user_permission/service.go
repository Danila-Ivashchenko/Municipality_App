package user_permission

import (
	"context"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *userPermissionService) SetPermissionsToUser(ctx context.Context, data *service.SetUserPermissionsData) ([]entity.Permission, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	var (
		permissionExistsMap = make(map[entity.Permission]struct{})
		permissionsToDelete []entity.Permission
		permissionsToCreate []entity.Permission
	)

	permissionExists, err := svc.UserPermissionRepository.GetUserPermissions(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	for _, p := range permissionExists {
		permissionExistsMap[p] = struct{}{}
	}

	for _, p := range data.Permissions {
		_, ok := permissionExistsMap[p]
		if !ok {
			permissionsToCreate = append(permissionsToCreate, p)
		}
		delete(permissionExistsMap, p)
	}

	for p := range permissionExistsMap {
		permissionsToDelete = append(permissionsToDelete, p)
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		deletePermissionsData := &repository.DeleteUserPermissionsData{
			UserID:      data.UserID,
			Permissions: permissionsToDelete,
		}

		err = svc.UserPermissionRepository.DeleteUserPermissions(tx, deletePermissionsData)
		if err != nil {
			return err
		}

		setPermissionsData := &repository.SetUserPermissionsData{
			UserID:      data.UserID,
			Permissions: permissionsToCreate,
		}

		_, err = svc.UserPermissionRepository.SetPermissionsToUser(tx, setPermissionsData)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return svc.UserPermissionRepository.GetUserPermissions(ctx, data.UserID)
}

func (svc *userPermissionService) GetUserPermissions(ctx context.Context, id int64) ([]entity.Permission, error) {
	return svc.UserPermissionRepository.GetUserPermissions(ctx, id)
}
