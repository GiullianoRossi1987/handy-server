package types

import "time"

type ProductService struct {
	Id                int
	IdWorker          int
	Name              string
	Description       string
	Available         bool
	QuantityAvailable *float32
	Service           bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
