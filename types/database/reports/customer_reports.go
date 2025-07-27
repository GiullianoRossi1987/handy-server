package types

import "time"

type CustomerReport struct {
	Id          int32
	Id_Customer int32
	Tags        []string
	Rating      float32
	Description string
	Revoked     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
