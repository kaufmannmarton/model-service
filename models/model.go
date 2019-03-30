package models

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	email      string
	password   string
	isVerified bool

	name     *string
	bio      *string
	avatar   *string
	video    *string
	pictures *[]string
}
