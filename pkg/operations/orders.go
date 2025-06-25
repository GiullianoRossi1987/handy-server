package pkg

import (
	"context"
	"fmt"
	types "types/database/operations"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOrderById(orderId int, conn *pgxpool.Pool) (*types.Order, error) {
	var order *types.Order
	if err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM orders WHERE id = $1;",
		orderId,
	).Scan(&order); err != nil {
		return nil, err
	}
	return order, nil
}

func GetCustomerOrders(customerId int, conn *pgxpool.Pool) ([]types.Order, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM orders WHERE id_customer = $1;",
		customerId,
	)
	if err != nil {
		return nil, err
	}
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetProductServiceOrders(productSericeId int, conn *pgxpool.Pool) ([]types.Order, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM orders WHERE id_product_service = $1;",
		productSericeId,
	)
	if err != nil {
		return nil, err
	}
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func AddOrder(order types.Order, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`INSERT INTO orders (
			id_product_service, 
			id_customer, 
			requested_at, 
			deployed_at, 
			description, 
			id_worker_addr, 
			id_customer_addr, 
			online, quantity, 
			quantity_by_time, 
			total_price, 
			customer_rating, 
			customer_feedback
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
		order.IdProductService,
		order.IdCustomer,
		order.RequestedAt,
		order.DeployedAt,
		order.Description,
		order.IdWorkerAddr,
		order.IdCustomerAddr,
		order.Online,
		order.Quantity,
		order.QuantityByTime,
		order.TotalPrice,
		order.CustomerRating,
		order.CustomerFeedback,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "orders",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier: fmt.Sprintf(
				"%d;%d;%s",
				order.IdProductService,
				order.IdCustomer,
				order.RequestedAt.String(),
			),
		}
	}
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
