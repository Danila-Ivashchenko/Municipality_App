package entity

import "time"

type User struct {
	ID        int64
	Email     string
	Name      string
	LastName  string
	IsAdmin   bool
	IsBlocked bool
	CreatedAt time.Time
}

type UserFull struct {
	ID        int64
	Email     string
	Name      string
	LastName  string
	IsAdmin   bool
	IsBlocked bool
	Password  string
	CreatedAt time.Time
}

type UserEx struct {
	User        User
	Permissions []Permission
}

func NewUserEx(user User, permissions []Permission) UserEx {
	return UserEx{
		User:        user,
		Permissions: permissions,
	}
}

func NewUserExPtr(user User, permissions []Permission) *UserEx {
	return &UserEx{
		User:        user,
		Permissions: permissions,
	}
}
