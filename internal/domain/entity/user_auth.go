package entity

import "time"

type UserAuthToken struct {
	ID     int64
	Token  string
	UserID int64

	CreatedAt time.Time
	ExpiresAt time.Time
}

const (
	DefaultExpireTime = time.Hour * 24 * 30
	UserTokensCount   = 4
)
