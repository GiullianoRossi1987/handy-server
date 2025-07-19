package types

import (
	"encoding/json"
	"time"

	db "types/database/operations"
)

type ProductServiceResponse struct {
	Id                int32      `json:"id" binding:"required"`
	IdWorker          int32      `json:"worker_id" binding:"required"`
	Name              string     `json:"name" binding:"required"`
	Description       string     `json:"description,omitempty"`
	Available         bool       `json:"available" binding:"required"`
	QuantityAvailable *float32   `json:"quantity" binding:"required"`
	Service           bool       `json:"is_service" binding:"required"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}

func SerializeProductService(record *db.ProductService) *ProductServiceResponse {
	if record == nil {
		return nil
	}
	return &ProductServiceResponse{
		Id:                record.Id,
		IdWorker:          record.IdWorker,
		Name:              record.Name,
		Description:       record.Description,
		Available:         record.Available,
		QuantityAvailable: record.QuantityAvailable,
		Service:           record.Service,
		CreatedAt:         &record.CreatedAt,
		UpdatedAt:         &record.UpdatedAt,
	}
}

func (prd *ProductServiceResponse) ToJSON() (string, error) {
	val, err := json.Marshal(prd)
	if err != nil {
		return "nil", err
	}
	return string(val), nil
}
