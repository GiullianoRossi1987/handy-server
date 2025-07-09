package types

import (
	"time"
	usr "types/database/users"
	serializables "types/serializables"
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
	return &usr.CustomerRecord{
		Id:        0,
		Fullname:  uur.Fullname,
		Active:    uur.Active,
		UUID:      "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
