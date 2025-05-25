package user

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return svc.UserRepository.GetAllUsers(ctx)
}
