package types

import (
	"time"
)

type OrderBody struct {
	IdProductService int        `json:"product_service" binding:"required"`
	IdCustomer       int        `json:"customer" binding:"required"`
	RequestedAt      time.Time  `json:"requested_at" binding:"required"`
	ScheduleTo       *time.Time `json:"schedule_to" binding:"required"`
	DeployedAt       *time.Time `json:"deployed_at" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	IdWorkerAddr     *int       `json:"worker_address" binding:"required"`
	IdCustomerAddr   *int       `json:"customer_address" binding:"required"`
	Online           bool       `json:"online" binding:"required"`
	Quantity         float32    `json:"qtd" binding:"required"`
	QuantityByTime   float32    `json:"qtd_at_time" binding:"required"`
	TotalPrice       float32    `json:"total" binding:"required"`
	CustomerRating   *float32   `json:"rating" binding:"required"`
	CustomerFeedback *string    `json:"feedback" binding:"required"`
}
