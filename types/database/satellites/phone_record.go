package types

import "time"

type PhoneRecord struct {
	Id          int32
	IdWorker    *int32
	IdCustomer  *int32
	PhoneNumber string
	AreaCode    string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
