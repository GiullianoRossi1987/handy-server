package types

import (
	db "types/database/operations"
)

type ProductServiceBody struct {
	IdWorker          int32    `json:"worker" binding:"required"`
	Name              string   `json:"name" binding:"required"`
	Description       string   `json:"description" binding:"required"`
	Available         bool     `json:"available" binding:"required"`
	QuantityAvailable *float32 `json:"qtd_available" binding:"required"`
	Service           bool     `json:"is_service" binding:"required"`
}

func (ps *ProductServiceBody) ToRecord() *db.ProductService {
	if ps == nil {
		return nil
	}
	return &db.ProductService{
		IdWorker:          ps.IdWorker,
		Name:              ps.Name,
		Description:       ps.Description,
		Available:         ps.Available,
		QuantityAvailable: ps.QuantityAvailable,
		Service:           ps.Service,
	}
}
