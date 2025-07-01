package types

import serializables "types/serializables"

type UpdateUserRequest struct {
	Uuid      string                  `json:"uuid" binding:"required"`
	Fullname  string                  `json:"name"`
	Active    bool                    `json:"active"`
	Phones    []serializables.Phone   `json:"phones" binding:"required"`
	Emails    []serializables.Email   `json:"emails" binding:"required"`
	Addresses []serializables.Address `json:"addresses" binding:"required"`
}
