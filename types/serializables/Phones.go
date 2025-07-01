package types

import (
	"encoding/json"
	"time"
)

type Phone struct {
	Id         int32      `json:"id,omitempty" binding:"required"`
	IdWorker   *int32     `json:"workerId,omitempty"`
	IdCustomer *int32     `json:"customerId,omitempty"`
	PhoneNumer string     `json:"number" binding:"required"`
	AreaCode   string     `json:"area_code" binding:"required"`
	Active     bool       `json:"active"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (p *Phone) ToJSON() (string, error) {
	val, err := json.Marshal(p)
	return string(val), err
}
