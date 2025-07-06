package types

import "fmt"

type ValidSatellite string

const (
	Email   ValidSatellite = "email"
	Phone   ValidSatellite = "phone"
	Address ValidSatellite = "address"
)

type NullUserAttachmentPointError struct {
	Satellite  ValidSatellite
	Operation  string
	Identifier string
}

func (e *NullUserAttachmentPointError) Error() string {
	return fmt.Sprintf(
		"Received a bad Satellite Request at the Satellite: %s operation: %s Indentifier: ¨%s",
		e.Satellite,
		e.Operation,
		e.Identifier,
	)
}
