package services

import (
	op "pkg/operations"
	types "types/database/operations"
	requests "types/requests/operations"
	responses "types/responses/operations"
	"utils"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOrderById(pool *pgxpool.Pool, id int32) (*responses.OrderResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := op.GetOrderById(id, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return responses.SerializeOrderRecord(data), nil
}

func GetWorkerOrders(pool *pgxpool.Pool, workerId int32) ([]responses.OrderResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := op.GetWorkerOrders(workerId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	data := utils.MapCar(
		records,
		func(record types.Order) responses.OrderResponse {
			return *responses.SerializeOrderRecord(&record)
		},
	)
	return data, nil
}

func GetCustomerOrders(pool *pgxpool.Pool, customerId int32) ([]responses.OrderResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := op.GetCustomerOrders(customerId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	data := utils.MapCar(
		records,
		func(record types.Order) responses.OrderResponse {
			return *responses.SerializeOrderRecord(&record)
		},
	)
	return data, nil
}

func GetCartOrders(pool *pgxpool.Pool, cartUUID string) ([]responses.OrderResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := op.GetCartOrders(cartUUID, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	data := utils.MapCar(
		records,
		func(record types.Order) responses.OrderResponse {
			return *responses.SerializeOrderRecord(&record)
		},
	)
	return data, nil
}

func PlaceOrder(pool *pgxpool.Pool, req requests.OrderBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	id, err := op.AddOrder(*record, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return id, nil
}

func UpdateOrder(pool *pgxpool.Pool, req requests.OrderBody, orderId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := req.ToRecord()
	record.Id = orderId
	if err := op.UpdateOrder(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteOrder(pool *pgxpool.Pool, orderId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := op.DeleteOrder(orderId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
