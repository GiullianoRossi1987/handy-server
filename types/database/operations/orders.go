package types

import "time"

type Order struct {
	Id               int
	IdProductService int
	IdCustomer       int
	RequestedAt      time.Time
	DeployedAt       *time.Time
	Description      string
	IdWorkerAddr     *int
	IdCustomerAddr   *int
	Online           bool
	Quantity         float32
	QuantityByTime   float32
	TotalPrice       float32
	CustomerRating   *float32
	CustomerFeedback *string
}
