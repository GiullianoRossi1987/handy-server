package pkg

import (
	"context"
	"fmt"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerById(id int32, conn *pgxpool.Conn) (*types.CustomerRecord, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM customers WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	customer, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.CustomerRecord])
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerByUUID(uuid string, conn *pgxpool.Conn) (*types.CustomerRecord, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM customers WHERE uuid = $1",
		uuid,
	)
	if err != nil {
		return nil, err
	}
	customer, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.CustomerRecord])
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerByUserId(id int32, conn *pgxpool.Conn) (*types.CustomerRecord, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM customers WHERE id_user = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	customer, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.CustomerRecord])
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func AddCustomer(customer types.CustomerRecord, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		"INSERT INTO customers (id_user, uuid, fullname) ($1, $2, $3) RETURNING id;",
		customer.UserId,
		customer.UUID,
		customer.Fullname,
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

func DeactivateCustomer(uuid string, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE workers SET active = FALSE, name = '', avg_ratings = 0, updated_at = CURRENT_TIMESTAMP
		WHERE uuid = $1::text;`,
		uuid,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "deactivate",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           uuid,
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteCustomer(uuid string, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM workers WHERE uuid = $1::string;`,
		uuid,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           uuid,
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateCustomer(newDataRecord types.CustomerRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE customers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3;",
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

func UpdateCustomerRating(newDataRecord types.CustomerRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE customers SET avg_rating = $1 WHERE id = $2;",
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

func DoesCustomerExists(customerId int32, conn *pgxpool.Conn) (bool, error) {
	res, err := GetCustomerById(customerId, conn)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}

func DoesCustomerUUIDExists(uuid string, conn *pgxpool.Conn) (bool, error) {
	res, err := GetCustomerByUUID(uuid, conn)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}
