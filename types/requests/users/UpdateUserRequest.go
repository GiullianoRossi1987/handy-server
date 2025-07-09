package types

import (
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
