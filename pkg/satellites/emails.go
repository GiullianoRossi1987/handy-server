package pkg

import (
	"context"
	"fmt"
	types "types/database/satellites"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetEmailById(emailId int32, conn *pgxpool.Conn) (*types.EmailRecord, error) {
	var email *types.EmailRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM emails WHERE id = $1;", emailId).Scan(email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func GetCustomerEmails(uuid string, conn *pgxpool.Conn) ([]types.EmailRecord, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT e.* FROM emails AS e INNER JOIN customers AS c ON c.id = e.id_worker WHERE c.uuid = $1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}
	emails, errCol := pgx.CollectRows(rows, pgx.RowToStructByPos[types.EmailRecord])
	if errCol != nil {
		return nil, err
	}
	return emails, nil
}

func GetWorkerEmails(uuid string, conn *pgxpool.Conn) ([]types.EmailRecord, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT e.* FROM emails AS e INNER JOIN workers AS w ON w.id = e.id_worker WHERE w.uuid = $1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}
	emails, errCol := pgx.CollectRows(rows, pgx.RowToStructByPos[types.EmailRecord])
	if errCol != nil {
		return nil, err
	}
	return emails, nil
}

func AddEmail(email types.EmailRecord, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO emails (id_worker, id_customer, email) VALUES ($1, $2, $3) RETURNING id;`,
		email.IdWorker, email.IdCustomer, email.Email,
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

func DeleteEmail(emailId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM emails WHERE id = $1;",
		emailId,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "emails",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", emailId),
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
	return err
}

func UpdateEmail(email types.EmailRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE emails SET id_worker = $1, id_customer = $2, email = $3, is_active = $4, updated_at = CURRENT_TIMESTAMP() WHERE id = $5;",
		email.IdWorker,
		email.IdCustomer,
		email.Email,
		email.Active,
		email.Id,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "emails",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", email.Id),
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

func DeleteEmailsFromCustomer(customerId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM emails WHERE customer_id = $1`,
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
			Table:                "emails",
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

func DeleteEmailsFromWorker(workerId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM emails WHERE worker_id = $1`,
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
			Table:                "phones",
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
