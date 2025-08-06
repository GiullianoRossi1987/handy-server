package types

import (
	"time"
	db "types/database/operations"
	"utils"

	"github.com/google/uuid"
)

type OrderBody struct {
	IdProductService int32      `json:"product_service" binding:"required"`
	IdCustomer       int32      `json:"customer" binding:"required"`
	RequestedAt      time.Time  `json:"requested_at" binding:"required"`
	ScheduleTo       *time.Time `json:"schedule_to" binding:"required"`
	DeployedAt       *time.Time `json:"deployed_at" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	IdWorkerAddr     *int32     `json:"worker_address" binding:"required"`
	IdCustomerAddr   *int32     `json:"customer_address" binding:"required"`
	Online           bool       `json:"online" binding:"required"`
	Quantity         float32    `json:"qtd" binding:"required"`
	QuantityByTime   float32    `json:"qtd_at_time" binding:"required"`
	TotalPrice       float32    `json:"total" binding:"required"`
	CustomerRating   *float32   `json:"rating" binding:"required"`
	CustomerFeedback *string    `json:"feedback" binding:"required"`
	CartUUID         *string    `json:"cartId"`
}

func (o *OrderBody) ToRecord() *db.Order {
	if o == nil {
		return nil
	}
	return &db.Order{
		IdProductService: o.IdProductService,
		IdCustomer:       o.IdCustomer,
		RequestedAt:      o.RequestedAt,
		ScheduleTo:       o.ScheduleTo,
		DeployedAt:       o.DeployedAt,
		Description:      o.Description,
		IdWorkerAddr:     o.IdWorkerAddr,
		IdCustomerAddr:   o.IdCustomerAddr,
		Online:           o.Online,
		Quantity:         o.Quantity,
		QuantityByTime:   o.QuantityByTime,
		TotalPrice:       o.TotalPrice,
		CustomerRating:   o.CustomerRating,
		CustomerFeedback: o.CustomerFeedback,
		CartUUID:         utils.Coalesce(o.CartUUID, uuid.NewString()),
	}
}
