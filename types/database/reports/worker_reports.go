package types

import "time"

type WorkerReport struct {
	Id          int
	Id_Worker   int
	Tags        []string
	Description string
	Revoked     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
