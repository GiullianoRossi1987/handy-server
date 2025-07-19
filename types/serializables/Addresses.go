package types

import (
	"encoding/json"
	"time"
	db "types/database/satellites"
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
	Main          bool       `json:"main" binding:"required"`
	Active        bool       `json:"active"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

func (addr *Address) ToJSON() (string, error) {
	val, err := json.Marshal(addr)
	if err != nil {
		return "nil", err
	}
	return string(val), err
}

func SerializeAddress(record *db.AddressRecord) *Address {
	if record == nil {
		return nil
	}
	return &Address{
		Id:            int32(record.Id),
		IdWorker:      new(int32),
		IdCustomer:    new(int32),
		Address:       "",
		AddressNumber: "",
		City:          "",
		UF:            "",
		Country:       "",
		Main:          false,
		Active:        false,
		CreatedAt:     &time.Time{},
		UpdatedAt:     &time.Time{},
	}
}
