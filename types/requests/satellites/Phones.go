package types

import (
	"fmt"
	"time"
	tp "types/database/satellites"
	errors "types/errors"
	"utils"
)

type PhoneBody struct {
	IdWorker    *int32 `json:"worker" binding:"required"`
	IdCustomer  *int32 `json:"customer" binding:"required"`
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

func (b *PhoneBody) ToRecord() *tp.PhoneRecord {
	if b == nil {
		return nil
	}
	return &tp.PhoneRecord{
		Id:          0,
		IdWorker:    b.IdWorker,
		IdCustomer:  b.IdCustomer,
		PhoneNumber: b.PhoneNumber,
		AreaCode:    b.AreaCode,
		Active:      b.Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
