package pkg

import (
	"context"
	"fmt"
	types "types/database/satellites"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPhoneById(phoneId int32, conn *pgxpool.Conn) (*types.PhoneRecord, error) {
	var phone *types.PhoneRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM phones WHERE id = $1;", phoneId).Scan(&phone)
	if err != nil {
		return nil, err
	}
	return phone, nil
}

func GetWorkerPhones(workerId int32, conn *pgxpool.Conn) ([]types.PhoneRecord, error) {
	row, err := conn.Query(context.Background(), "SELECT * FROM phones WHERE id_worker = $1;", workerId)
	if err != nil {
		return nil, err
	}
	phones, err := pgx.CollectRows(row, pgx.RowToStructByPos[types.PhoneRecord])
	if err != nil {
		return nil, err
	}
	return phones, nil
}

func GetCustomerPhones(customerId int32, conn *pgxpool.Conn) ([]types.PhoneRecord, error) {
	row, err := conn.Query(context.Background(), "SELECT * FROM phones WHERE id_customer = $1;", customerId)
	if err != nil {
		return nil, err
	}
	phones, err := pgx.CollectRows(row, pgx.RowToStructByPos[types.PhoneRecord])
	if err != nil {
		return nil, err
	}
	return phones, nil
}

func AddPhone(phone types.PhoneRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"INSERT INTO phones (id_worker, id_customer, phone_number, area_code) VALUES ($1, $2, $3, $4);",
		phone.IdWorker,
		phone.IdCustomer,
		phone.PhoneNumber,
		phone.AreaCode,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "phones",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", phone.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeletePhone(phoneId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM phones WHERE id = $1;",
		phoneId,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "phones",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", phoneId),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func UpdatePhone(phone types.PhoneRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE phones SET id_worker = $1, id_customer = $2, phone_number = $3, area_code = $4, updated_at = CURRENT_TIMESTAMP()  WHERE id = $1;",
		phone.IdWorker,
		phone.IdCustomer,
		phone.PhoneNumber,
		phone.AreaCode,
		phone.Id,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update",
			Table:                "phones",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", phone.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
