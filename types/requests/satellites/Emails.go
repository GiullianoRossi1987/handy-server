package types

import (
	"time"
	tp "types/database/satellites"
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

func (b *EmailBody) ToRecord() *tp.EmailRecord {
	if b == nil {
		return nil
	}
	return &tp.EmailRecord{
		Id:         0,
		IdWorker:   b.IdWorker,
		IdCustomer: b.IdCustomer,
		Email:      b.Email,
		Active:     b.Active,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
