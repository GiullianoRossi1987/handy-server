package types

import "time"

type UsersRecord struct {
	Id        int
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
