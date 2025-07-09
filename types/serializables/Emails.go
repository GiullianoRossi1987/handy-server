package types

import (
	"encoding/json"
	"time"
	db "types/database/satellites"
)

type Email struct {
	Id         int32      `json:"id,omitempty" binding:"required"`
	IdWorker   *int32     `json:"workerId,omitempty"`
	IdCustomer *int32     `json:"customerId,omitempty"`
	Email      string     `json:"email" binding:"required"`
	Active     bool       `json:"active"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (email *Email) ToJSON() (string, error) {
	val, err := json.Marshal(email)
	return string(val), err
}

func SerializeEmail(record db.EmailRecord) Email {
	return Email{
		Id:         record.Id,
		IdWorker:   &record.IdWorker,
		IdCustomer: &record.IdCustomer,
		Email:      record.Email,
		Active:     record.Active,
		CreatedAt:  &record.CreatedAt,
		UpdatedAt:  &record.UpdatedAt,
	}
}
