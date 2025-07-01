package types

import (
	"encoding/json"
	"time"
)

type Address struct {
	Id            int32      `json:"id,omitempty" binding:"required"`
	IdWorker      *int32     `json:"workerId,omitempty"`
	IdCustomer    *int32     `json:"customerId,omitempty"`
	Address       string     `json:"address" binding:"required"`
	AddressNumber string     `json:"number" binding:"required"`
	City          string     `json:"city" binding:"required"`
	UF            string     `json:"uf" binding:"required"`
	Country       string     `json:"country" binding:"required"`
	Active        bool       `json:"active"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

func (email *Address) ToJSON() (string, error) {
	val, err := json.Marshal(email)
	return string(val), err
}
