package types

import (
	"encoding/xml"
	operations "types/database/operations"
	users "types/database/users"
)

type ComposedOrder struct {
	Customer   users.CustomerRecord `json:"customer"`
	Worker     users.WorkersRecord  `json:"worker"`
	UsingOrder operations.Order
}

func (co *ComposedOrder) ToXML() (string, error) {
	val, err := xml.Marshal(co)
	if err != nil {
		return "nil", err
	}
	return string(val), nil
}
