package types

import "time"

type AddressRecord struct {
	Id            int32
	IdWorker      *int32
	IdCustomer    *int32
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
