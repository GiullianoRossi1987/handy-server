package types

import (
	"fmt"
	errors "types/errors"
	"utils"
)

type PhoneBody struct {
	IdWorker    *int   `json:"worker" binding:"required"`
	IdCustomer  *int   `json:"customer" binding:"required"`
	PhoneNumber string `json:"number" binding:"required"`
	AreaCode    string `json:"area_code" binding:"required"`
	Active      bool   `json:"active" binding:"required"`
}

func (b *PhoneBody) Validate(operation *string) error {
	if b.IdCustomer == b.IdWorker && b.IdWorker == nil {
		return &errors.NullUserAttachmentPointError{
			Satellite:  errors.Email,
			Operation:  utils.Coalesce(operation, ""),
			Identifier: fmt.Sprintf("(%s) %s", b.AreaCode, b.PhoneNumber),
		}
	}
	return nil
}
