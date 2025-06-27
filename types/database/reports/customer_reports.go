package types

import "time"

type CustomerReport struct {
	Id          int
	Id_Customer int
	Tags        []string
	Description string
	Revoked     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
