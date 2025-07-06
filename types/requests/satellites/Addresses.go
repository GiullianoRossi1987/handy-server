package types

import (
	"fmt"
	errors "types/errors"
	"utils"
)

type AddressBody struct {
	IdWorker      *int32 `json:"worker" binding:"required"`
	IdCustomer    *int32 `json:"customer" binding:"required"`
	Address       string `json:"address" binding:"required"`
	AddressNumber string `json:"number" binding:"required"`
	City          string `json:"city" binding:"required"`
	UF            string `json:"uf" binding:"required"`
	Country       string `json:"country" binding:"required"`
	Main          bool   `json:"main" binding:"required"`
	Active        bool   `json:"active" binding:"required"`
}

func (b *AddressBody) Validate(operation *string) error {
	if b.IdCustomer == b.IdWorker && b.IdWorker == nil {
		return &errors.NullUserAttachmentPointError{
			Satellite: errors.Email,
			Operation: utils.Coalesce(operation, ""),
			Identifier: fmt.Sprintf(
				"%s - %s (%s, %s)",
				b.Address,
				b.AddressNumber,
				b.UF,
				b.Country,
			),
		}
	}
	return nil
}
