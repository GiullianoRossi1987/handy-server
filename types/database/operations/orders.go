package types

import "time"

type Order struct {
	Id               int32
	IdProductService int32
	IdCustomer       int32
	RequestedAt      time.Time
	DeployedAt       *time.Time
	ScheduleTo       *time.Time
	Description      string
	IdWorkerAddr     *int32
	IdCustomerAddr   *int32
	Online           bool
	Quantity         float32
	QuantityByTime   float32
	TotalPrice       float32
	CustomerRating   *float32
	CustomerFeedback *string
	UpdatedAt        time.Time
	CartUUID         string
}
