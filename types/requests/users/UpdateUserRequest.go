package types

import (
	"time"
	usr "types/database/users"
	serializables "types/serializables"

	"github.com/google/uuid"
)

// PUT REQUEST
type UpdateUserRequest struct {
	Fullname  string                  `json:"name" binding:"required"`
	Active    bool                    `json:"active" binding:"required"`
	Phones    []serializables.Phone   `json:"phones" binding:"required"`
	Emails    []serializables.Email   `json:"emails" binding:"required"`
	Addresses []serializables.Address `json:"addresses" binding:"required"`
}

func (uur *UpdateUserRequest) ToCustomerRecord() *usr.CustomerRecord {
	if uur == nil {
		return nil
	}
	return &usr.CustomerRecord{
		Id:        0,
		Fullname:  uur.Fullname,
		Active:    uur.Active,
		UUID:      uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (uur *UpdateUserRequest) ToWorkerRecord() *usr.WorkersRecord {
	if uur == nil {
		return nil
	}
	return &usr.WorkersRecord{
		Id:        0,
		Fullname:  uur.Fullname,
		Active:    uur.Active,
		UUID:      uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
