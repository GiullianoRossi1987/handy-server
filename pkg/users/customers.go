package pkg

import (
	"context"
	"fmt"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerById(id int32, conn *pgxpool.Pool) (*types.CustomerRecord, error) {
	var customer *types.CustomerRecord
	if err := conn.QueryRow(context.Background(), "SELECT * FROM customers WHERE id = $1", id).Scan(customer); err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerByUUID(uuid string, conn *pgxpool.Pool) (*types.CustomerRecord, error) {
	var customer *types.CustomerRecord
	if err := conn.QueryRow(context.Background(), "SELECT * FROM customers WHERE uuid = $1", uuid).Scan(customer); err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerByUserId(id int32, conn *pgxpool.Pool) (*types.CustomerRecord, error) {
	var customer *types.CustomerRecord
	if err := conn.QueryRow(context.Background(), "SELECT * FROM customers WHERE id_user = $1", id).Scan(customer); err != nil {
		return nil, err
	}
	return customer, nil
}

func AddCustomer(customer types.CustomerRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"INSERT INTO customers (id_user, uuid, fullname) ($1, $2, $3);",
		customer.UserId,
		customer.UUID,
		customer.Fullname,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "customers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", customer.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteCustomer(id int32, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil
	}
	commandTag, err := conn.Exec(context.Background(), "DELETE FROM customers WHERE id = $1;", id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "customers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateCustomer(newDataRecord types.CustomerRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE FROM customers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3;",
		newDataRecord.Fullname,
		newDataRecord.Active,
		newDataRecord.Id,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update",
			Table:                "customers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", newDataRecord.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateCustomerRating(newDataRecord types.CustomerRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE FROM customers SET avg_rating = $1 WHERE id = $2;",
		newDataRecord.Avg_ratings,
		newDataRecord.Id,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update rating",
			Table:                "customers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", newDataRecord.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DoesCustomerExists(customerId int32, conn *pgxpool.Pool) (bool, error) {
	res, err := GetCustomerById(customerId, conn)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}

func DoesCustomerUUIDExists(uuid string, conn *pgxpool.Pool) (bool, error) {
	res, err := GetCustomerByUUID(uuid, conn)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}
