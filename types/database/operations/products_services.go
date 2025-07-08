package types

import "time"

type ProductService struct {
	Id                int32
	IdWorker          int32
	Name              string
	Description       string
	Available         bool
	QuantityAvailable *float32
	Service           bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
