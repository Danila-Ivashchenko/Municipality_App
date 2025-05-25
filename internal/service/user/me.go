package user

import (
	"context"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
)

func (svc *userService) Me(ctx context.Context, id int64) (*entity.UserEx, error) {
	userData, err := svc.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if userData == nil {
		return nil, core_errors.UserNotFound
	}

	permissions, err := svc.UserPermissionService.GetUserPermissions(ctx, id)
	if err != nil {
		return nil, err
	}

	result := entity.NewUserEx(*userData, permissions)
	return &result, nil
}
