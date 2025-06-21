package pkg

import (
	"context"
	types "types/database/satellites"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetPhoneById(phoneId int32, conn *pgx.Conn) *types.PhoneRecord {
	var phone *types.PhoneRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM phones WHERE id = $1;", phoneId).Scan(&phone)
	if err != nil {
		panic(err)
	}
	return phone
}

func GetWorkerPhones(workerId int32, conn *pgx.Conn) []types.PhoneRecord {
	row, errQ := conn.Query(context.Background(), "SELECT * FROM phones WHERE id_worker = $1;", workerId)
	if errQ != nil {
		panic(errQ)
	}
	phones, errC := pgx.CollectRows(row, pgx.RowTo[types.PhoneRecord])
	if errC != nil {
		panic(errC)
	}
	return phones
}

func GetCustomerPhones(customerId int32, conn *pgx.Conn) []types.PhoneRecord {
	row, errQ := conn.Query(context.Background(), "SELECT * FROM phones WHERE id_customer = $1;", customerId)
	if errQ != nil {
		panic(errQ)
	}
	phones, errC := pgx.CollectRows(row, pgx.RowTo[types.PhoneRecord])
	if errC != nil {
		panic(errC)
	}
	return phones
}

func AddPhone(phone types.PhoneRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO phones (id_worker, id_customer, phone_number, area_code) VALUES ($1, $2, $3, $4);",
		phone.IdWorker, phone.IdCustomer, phone.PhoneNumber, phone.AreaCode,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeletePhone(phoneId int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM phones WHERE id = $1;",
		phoneId,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdatePhone(phone types.PhoneRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE phones SET id_worker = $1, id_customer = $2, phone_number = $3, area_code = $4, updated_at = CURRENT_TIMESTAMP()  WHERE id = $1;",
		phone.IdWorker, phone.IdCustomer, phone.PhoneNumber, phone.AreaCode, phone.Id,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}
