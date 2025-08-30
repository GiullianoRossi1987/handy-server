package main

import (
	types "types/responses"
	op "types/responses/operations"
	usr "types/responses/users"
	"utils"
)

func mainPdf() {
	order := op.OrderResponse{}
	customer := usr.CustomerResponseBody{}
	worker := usr.WorkerResponseBody{}
	content := types.ComposedOrderResponse{
		UsingOrder: order,
		Worker:     worker,
		Customer:   customer,
	}
	utils.OrderToPDF(&content)
}
