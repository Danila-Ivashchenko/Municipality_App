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
