package types

import (
	"encoding/json"
	db "types/database"
	response_opr "types/responses/operations"
	response_usr "types/responses/users"
)

type ComposedOrderResponse struct {
	Customer   response_usr.CustomerResponseBody `json:"customer" binding:"required"`
	Worker     response_usr.WorkerResponseBody   `json:"worker" binding:"required"`
	UsingOrder response_opr.OrderResponse        `json:"order" binding:"required"`
}

func SerializeComposedOrderResponse(order_record db.ComposedOrder) ComposedOrderResponse {
	return ComposedOrderResponse{
		Customer:   *response_usr.SerializeCustomerResponse(order_record.Customer),
		Worker:     *response_usr.SerializeWorkerResponse(order_record.Worker),
		UsingOrder: response_opr.SerializeOrderRecord(order_record.UsingOrder),
	}
}

func (cor *ComposedOrderResponse) ToJSON() (string, error) {
	val, err := json.Marshal(cor)
	if err != nil {
		return "nil", err
	}
	return string(val), err
}
