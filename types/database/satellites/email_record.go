package types

import "time"

type EmailRecord struct {
	Id         int32
	IdWorker   *int32
	IdCustomer *int32
	Email      string
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
