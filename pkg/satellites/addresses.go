package pkg

import (
	"context"
	"fmt"
	types "types/database/satellites"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerAddresses(customerId int32, conn *pgxpool.Pool) ([]types.AddressRecord, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM addresses WHERE id_customer = $1;", customerId)
	if err != nil {
		return nil, err
	}
	addresses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.AddressRecord])
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetWorkerAddresses(workerId int32, conn *pgxpool.Pool) ([]types.AddressRecord, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM addresses WHERE id_worker = $1;", workerId)
	if err != nil {
		return nil, err
	}
	addresses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.AddressRecord])
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetAddressById(addressId int32, conn *pgxpool.Pool) (*types.AddressRecord, error) {
	var row *types.AddressRecord
	if err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM addresses WHERE id = $1;",
		addressId,
	).Scan(&row); err != nil {
		return nil, err
	}
	return row, nil
}

func AddAddress(address types.AddressRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"INSERT INTO addresses (id_worker, id_customer, address, address_number, city, uf, country) VALUES ($1,$2,$3,$4,$5,$6);",
		address.IdWorker,
		address.IdCustomer,
		address.Address,
		address.AddressNumber,
		address.City,
		address.UF,
		address.Country,
	)
	if commandTag.RowsAffected() != 1 {
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "Insert",
			Table:                "addresses",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
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

func DeleteAddress(addressId int32, conn *pgxpool.Pool) error {
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

func UpdatedAddress(address types.AddressRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE addresses id_worker = $1, id_customer = $2, address = $3, address_number = $4, city = $5, uf = $6, country = $8 WHERE id = $8;",
		address.IdWorker,
		address.IdCustomer,
		address.Address,
		address.AddressNumber,
		address.City,
		address.UF,
		address.Country,
	)
	if commandTag.RowsAffected() != 1 {
		err = fmt.Errorf("UPDATED MORE THAN ONE ADDRESS PER ACTION")
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
