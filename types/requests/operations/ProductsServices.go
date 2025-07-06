package types

type ProductServiceBody struct {
	IdWorker          int      `json:"worker" binding:"required"`
	Name              string   `json:"name" binding:"required"`
	Description       string   `json:"description" binding:"required"`
	Available         bool     `json:"available" binding:"required"`
	QuantityAvailable *float32 `json:"qtd_available" binding:"required"`
	Service           bool     `json:"is_service" binding:"required"`
}
