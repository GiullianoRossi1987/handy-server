package types

import (
	"time"
	db "types/database/users"
	serial "types/serializables"
)

type CustomerResponseBody struct {
	Id        int              `json:"id" binding:"required"`
	UserId    int              `json:"user_id" binding:"required"`
	Uuid      string           `json:"uuid" binding:"required"`
	Name      string           `json:"name" binding:"required"`
	Active    bool             `json:"active" binding:"required"`
	AvgRating float32          `json:"rating" binding:"required"`
	CreatedAt time.Time        `json:"created_at" binding:"required"`
	UpdatedAt time.Time        `json:"updated_at" binding:"required"`
	Phones    []serial.Phone   `json:"Phones,omitempty" binding:"required"`
	Emails    []serial.Email   `json:"Emails,omitempty" binding:"required"`
	Addresses []serial.Address `json:"Addresses,omitempty" binding:"required"`
}

func SerializeCustomerResponse(record db.CustomerRecord) CustomerResponseBody {
	return CustomerResponseBody{
		Id:        record.Id,
		UserId:    record.UserId,
		Uuid:      record.UUID,
		Name:      record.Fullname,
		Active:    record.Active,
		AvgRating: record.Avg_ratings,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}
