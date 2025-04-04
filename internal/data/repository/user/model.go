package user

import (
	"database/sql"
	sql_convertor "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type createUserModel struct {
	ID       sql.NullInt64
	Email    sql.NullString
	Name     sql.NullString
	LastName sql.NullString
	Password sql.NullString
	IsAdmin  sql.NullBool
}

func newCreateUserModel(data *repository.CreateUserData) createUserModel {
	return createUserModel{
		Email:    sql_convertor.NewNullString(data.Email),
		Name:     sql_convertor.NewNullString(data.Name),
		LastName: sql_convertor.NewNullString(data.LastName),
		Password: sql_convertor.NewNullString(data.Password),
		IsAdmin:  sql_convertor.NewNullBool(data.IsAdmin),
	}
}

type userModel struct {
	ID        sql.NullInt64
	Email     sql.NullString
	Name      sql.NullString
	LastName  sql.NullString
	IsAdmin   sql.NullBool
	IsBlocked sql.NullBool
	CreatedAt sql.NullTime
}

func newUserModel(i *entity.User) userModel {
	return userModel{
		ID:        sql_convertor.NewNullInt64(i.ID),
		Email:     sql_convertor.NewNullString(i.Email),
		Name:      sql_convertor.NewNullString(i.Name),
		LastName:  sql_convertor.NewNullString(i.LastName),
		IsAdmin:   sql_convertor.NewNullBool(i.IsAdmin),
		IsBlocked: sql_convertor.NewNullBool(i.IsBlocked),
		CreatedAt: sql_convertor.NewNullTime(i.CreatedAt),
	}
}

func (m userModel) convert() *entity.User {
	return &entity.User{
		ID:        m.ID.Int64,
		Email:     m.Email.String,
		Name:      m.Name.String,
		LastName:  m.LastName.String,
		IsAdmin:   m.IsAdmin.Bool,
		IsBlocked: m.IsBlocked.Bool,
		CreatedAt: m.CreatedAt.Time,
	}
}

type userFullModel struct {
	ID        sql.NullInt64
	Email     sql.NullString
	Name      sql.NullString
	LastName  sql.NullString
	IsAdmin   sql.NullBool
	Password  sql.NullString
	IsBlocked sql.NullBool
	CreatedAt sql.NullTime
}

func (m *userFullModel) convert() *entity.UserFull {
	return &entity.UserFull{
		ID:        m.ID.Int64,
		Email:     m.Email.String,
		Name:      m.Name.String,
		LastName:  m.LastName.String,
		IsAdmin:   m.IsAdmin.Bool,
		IsBlocked: m.IsBlocked.Bool,
		Password:  m.Password.String,
		CreatedAt: m.CreatedAt.Time,
	}
}
