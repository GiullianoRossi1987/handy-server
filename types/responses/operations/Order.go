package types

import (
	"encoding/json"
	"time"

	db "types/database/operations"
)

type OrderResponse struct {
	Id               int32      `json:"id" binding:"required"`
	IdProductService int32      `json:"id_product_service" binding:"required"`
	IdCustomer       int32      `json:"customer_id" binding:"required"`
	RequestedAt      time.Time  `json:"requested_at" binding:"required"`
	DeployedAt       *time.Time `json:"deployed_at,omitempty"`
	ScheduleTo       *time.Time `json:"schedule_to,omitempty"`
	Description      string     `json:"description" binding:"required"`
	IdWorkerAddr     *int32     `json:"id_worker_addr,omitempty"`
	IdCustomerAddr   *int32     `json:"id_customer_addr,omitempty"`
	Online           bool       `json:"online" binding:"required"`
	Quantity         float32    `json:"qtd" binding:"required"`
	QuantityByTime   float32    `json:"qtd_at_time" binding:"required"`
	TotalPrice       float32    `json:"total" binding:"required"`
	CustomerRating   *float32   `json:"rating,omitempty"`
	CustomerFeedback *string    `json:"feedback,omitempty" binding:"required"`
	UpdatedAt        time.Time  `json:"updated_at,omitempty"`
}

func SerializeOrderRecord(record db.Order) OrderResponse {
	return OrderResponse{
		Id:               record.Id,
		IdProductService: record.IdProductService,
		IdCustomer:       record.IdCustomer,
		RequestedAt:      record.RequestedAt,
		ScheduleTo:       record.ScheduleTo,
		Description:      record.Description,
		IdWorkerAddr:     record.IdWorkerAddr,
		IdCustomerAddr:   record.IdCustomerAddr,
		Online:           record.Online,
		Quantity:         record.Quantity,
		QuantityByTime:   record.QuantityByTime,
		TotalPrice:       record.TotalPrice,
		CustomerRating:   record.CustomerRating,
		CustomerFeedback: record.CustomerFeedback,
		UpdatedAt:        record.UpdatedAt,
	}
}

func (or *OrderResponse) ToJSON() (string, error) {
	val, err := json.Marshal(or)
	if err != nil {
		return "nil", err
	}
	return string(val), nil
}
