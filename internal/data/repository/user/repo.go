package user

import (
	"context"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type userRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.UserRepository {
	repo := &userRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r userRepository) CreateUser(ctx context.Context, data *repository.CreateUserData) (*entity.User, error) {
	var (
		id uint
	)

	m := newCreateUserModel(data)

	row := r.handler.QueryRowContext(ctx, createUserQuery, m.Email, m.Name, m.LastName, m.Password)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&id,
	)
	if err != nil {
		return nil, err
	}

	return r.fetchUserWithСondition(ctx, " id = $1", id)
}

func (r userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	m := newUserModel(user)

	err := r.exexUserQuery(ctx, updateUserQuery, m.Name, m.LastName, m.Email, m.ID)
	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}

func (r userRepository) ChangeUserPassword(ctx context.Context, userID int64, password string) error {
	return r.exexUserQuery(ctx, changeUserPasswordQuery, password, userID)
}

func (r userRepository) ChangeUserAdmin(ctx context.Context, userID int64, isAdmin bool) error {
	return r.exexUserQuery(ctx, changeUserAdminQuery, isAdmin, userID)
}

func (r userRepository) ChangeUserBlocked(ctx context.Context, userID int64, isBlocked bool) error {
	return r.exexUserQuery(ctx, changeUserBlockedQuery, isBlocked, userID)
}

func (r userRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return r.fetchUserWithСondition(ctx, "id = $1", id)
}

func (r userRepository) GetUserFullByID(ctx context.Context, id int64) (*entity.UserFull, error) {
	return r.fetchUserFullWithСondition(ctx, "id = $1", id)
}

func (r userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.fetchUserWithСondition(ctx, "email = $1", email)
}

func (r userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return r.fetchUsers(ctx, selectUserQuery)
}

func (r userRepository) GetUserFullByEmail(ctx context.Context, email string) (*entity.UserFull, error) {
	return r.fetchUserFullWithСondition(ctx, "email = $1", email)

}
