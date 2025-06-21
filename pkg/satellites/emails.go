package pkg

import (
	"context"
	types "types/database/satellites"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetEmailById(emailId int32, conn *pgx.Conn) *types.EmailRecord {
	var email *types.EmailRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM emails WHERE id = $1;", emailId).Scan(email)
	if err != nil {
		panic(err)
	}
	return email
}

func GetCustomerEmails(customerId int32, conn *pgx.Conn) []types.EmailRecord {
	rows, err := conn.Query(context.Background(), "SELECT * FROM emails WHERE id_customer = $1;", customerId)
	if err != nil {
		panic(err)
	}
	emails, errCol := pgx.CollectRows(rows, pgx.RowTo[types.EmailRecord])
	if errCol != nil {
		panic(errCol)
	}
	return emails
}

func GetWorkerEmails(workerId int32, conn *pgx.Conn) []types.EmailRecord {
	rows, err := conn.Query(context.Background(), "SELECT * FROM emails WHERE id_worker = $1;", workerId)
	if err != nil {
		panic(err)
	}
	emails, errCol := pgx.CollectRows(rows, pgx.RowTo[types.EmailRecord])
	if errCol != nil {
		panic(errCol)
	}
	return emails
}

func AddEmail(email types.EmailRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO emails (id_worker, id_customer, email) VALUES ($1, $2, $3);",
		email.IdWorker, email.IdCustomer, email.Email,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteEmail(emailId int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM emails WHERE id = $1;",
		emailId,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdateEmail(email types.EmailRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE emails SET id_worker = $1, id_customer = $2, email = $3, is_active = $4, updated_at = CURRENT_TIMESTAMP() WHERE id = $5;",
		email.IdWorker, email.IdCustomer, email.Email, email.Active, email.Id,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}
