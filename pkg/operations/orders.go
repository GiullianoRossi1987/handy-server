package pkg

import (
	"context"
	"fmt"
	types "types/database/operations"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOrderById(orderId int32, conn *pgxpool.Conn) (*types.Order, error) {
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

func GetCustomerOrders(customerId int32, conn *pgxpool.Conn) ([]types.Order, error) {
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

func GetWorkerOrders(workerId int32, conn *pgxpool.Conn) ([]types.Order, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM orders WHERE id_worker = $1;",
		workerId,
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

func GetProductServiceOrders(productSericeId int32, conn *pgxpool.Conn) ([]types.Order, error) {
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

func AddOrder(order types.Order, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO orders (
      id_product_service, 
      id_customer, 
      requested_at, 
      deployed_at,
			scheduled_to, 
      description,
      id_worker_addr, 
      id_customer_addr, 
      online, quantity, 
      quantity_by_time, 
      total_price, 
      customer_rating, 
      customer_feedback
      )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;`,
		order.IdProductService,
		order.IdCustomer,
		order.RequestedAt,
		order.ScheduleTo,
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
	).Scan(&id); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}
	return &id, nil
}

func DeployOrder(order_id int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE orders SET 
		deployed_at = CURRENT_TIMESTAMP, 
		updated_at = CURRENT_TIMESTAMP 
		WHERE id = $1;`,
		order_id,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "orders",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           string(order_id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateOrder(order types.Order, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE orders SET 
    description = $2, 
		scheduled_to = $3
    id_worker_addr = $4, 
    id_customer_addr = $5,
    quantity = $6,
    quantity_by_time = $7,
    total_price = $8
    updated_at = CURRENT_TIMESTAMP WHERE id = $1;`,
		order.Id,
		order.Description,
		order.ScheduleTo,
		order.IdWorkerAddr,
		order.IdCustomerAddr,
		order.Quantity,
		order.QuantityByTime,
		order.TotalPrice,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "orders",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", order.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteOrder(orderId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM orders WHERE id = $1 AND deployed_at = NULL;`,
		orderId,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "orders",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", orderId),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
