package types

import "time"

type PhoneRecord struct {
	Id          int
	IdWorker    int
	IdCustomer  int
	PhoneNumber string
	AreaCode    string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
