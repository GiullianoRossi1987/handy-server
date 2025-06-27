package types

import "time"

type EmailRecord struct {
	Id         int
	IdWorker   int
	IdCustomer int
	Email      string
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
