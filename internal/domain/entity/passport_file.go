package entity

import "time"

type PassportFile struct {
	ID         int64
	Path       string
	PassportID int64
	FileName   string

	CreateAt time.Time
}
