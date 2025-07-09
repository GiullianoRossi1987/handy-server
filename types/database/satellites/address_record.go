package types

import "time"

type AddressRecord struct {
	Id            int
	IdWorker      *int
	IdCustomer    *int
	Address       string
	AddressNumber string
	City          string
	UF            string
	Country       string
	Main          bool
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
