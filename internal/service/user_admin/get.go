package user_admin

import (
	"context"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
)

func (svc *userAdminService) GetAll(ctx context.Context) ([]entity.UserEx, error) {
	var (
		result []entity.UserEx
	)

	allUsers, err := svc.UserService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range allUsers {
		userEx, err := svc.GetByID(ctx, user.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *userEx)
	}

	return result, nil
}

func (svc *userAdminService) GetByID(ctx context.Context, userID int64) (*entity.UserEx, error) {
	user, err := svc.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, core_errors.UserNotFound
	}

	userPermissions, err := svc.UserPermissionService.GetUserPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return entity.NewUserExPtr(*user, userPermissions), nil

}
