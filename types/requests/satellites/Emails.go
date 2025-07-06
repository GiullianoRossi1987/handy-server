package types

import (
	errors "types/errors"
	utils "utils"
)

type EmailBody struct {
	IdWorker   *int32 `json:"worker,omitempty"`
	IdCustomer *int32 `json:"customer,omitempty"`
	Email      string `json:"email" binding:"required"`
	Active     bool   `json:"active" binding:"required"`
}

func (b *EmailBody) Validate(operation *string) error {
	if b.IdCustomer == b.IdWorker && b.IdWorker == nil {
		return &errors.NullUserAttachmentPointError{
			Satellite:  errors.Email,
			Operation:  utils.Coalesce(operation, ""),
			Identifier: b.Email,
		}
	}
	return nil
}
