package types

import "time"

type WorkerReport struct {
	Id          int32
	Id_Worker   int32
	Tags        []string
	Rating      float32
	Description string
	Revoked     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
