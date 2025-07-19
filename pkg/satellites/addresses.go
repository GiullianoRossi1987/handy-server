package pkg

import (
	"context"
	"fmt"
	types "types/database/satellites"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerAddresses(uuid string, conn *pgxpool.Conn) ([]types.AddressRecord, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT a.* FROM addresses AS a INNER JOIN customers AS c ON c.id = a.customer_id WHERE c.uuid = $1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}
	addresses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.AddressRecord])
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetWorkerAddresses(uuid string, conn *pgxpool.Conn) ([]types.AddressRecord, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT a.* FROM addresses AS a INNER JOIN workers AS w ON w.id = a.customer_id WHERE w.uuid = $1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}
	addresses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.AddressRecord])
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetAddressById(addressId int32, conn *pgxpool.Conn) (*types.AddressRecord, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM addresses WHERE id = $1;",
		addressId,
	)
	if err != nil {
		return nil, err
	}
	address, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.AddressRecord])
	if err != nil {
		return nil, err
	}
	return address, nil
}

func AddAddress(address types.AddressRecord, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO addresses (
			id_worker, 
			id_customer, 
			address, 
			address_number, 
			city, 
			main,
			uf, 
			country) 
		VALUES ($1,$2,$3,$4,$5,$6, $7) RETURNING id;`,
		address.IdWorker,
		address.IdCustomer,
		address.Address,
		address.AddressNumber,
		address.City,
		address.Main,
		address.UF,
		address.Country,
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

func DeleteAddress(addressId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM addresses WHERE id = $1;",
		addressId)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "Delete",
			Table:                "addresses",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", addressId),
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

func UpdatedAddress(address types.AddressRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE addresses SET 
		 id_worker = $2, 
		 id_customer = $3, 
		 address = $4, 
		 address_number = $5, 
		 city = $6, 
		 uf = $7, 
		 country = $9,
		 main = $10
		 WHERE id = $1;`,
		address.Id,
		address.IdWorker,
		address.IdCustomer,
		address.Address,
		address.AddressNumber,
		address.City,
		address.UF,
		address.Country,
		address.Main,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "Insert",
			Table:                "addresses",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", address.Id),
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

func DeleteAddrsFromCustomer(customerId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM addresses WHERE customer_id = $1`,
		customerId,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete customer's",
			Table:                "addresses",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", customerId),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteAddrsFromWorker(workerId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM addresses WHERE worker_id = $1`,
		workerId,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete customer's",
			Table:                "addresses",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", workerId),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
