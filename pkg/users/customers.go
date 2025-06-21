package pkg

import (
	"context"
	types "types/database/users"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetCustomerById(id int32, conn *pgx.Conn) *types.CustomerRecord {
	var customer *types.CustomerRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM customers WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return customer
}

func GetCustomerByUserId(id int32, conn *pgx.Conn) *types.CustomerRecord {
	var customer *types.CustomerRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM customers WHERE id_user = $1", id)
	if err != nil {
		panic(err)
	}
	return customer
}

func AddCustomer(customer types.CustomerRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO customers (id_user, uuid, fullname) ($1, $2, $3);",
		customer.UserId, customer.UUID, customer.Fullname)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteCustomer(id int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(), "DELETE FROM customers WHERE id = $1;", id)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdateCustomer(newDataRecord types.CustomerRecord, conn *pgx.Conn) types.CustomerRecord {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM customers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3;",
		newDataRecord.Fullname, newDataRecord.Active, newDataRecord.Id)
	utils.CheckRowsAndError(commandTag, &err, 1)
	return newDataRecord
}

func UpdateCustomerRating(newDataRecord types.CustomerRecord, conn *pgx.Conn) types.CustomerRecord {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM customers SET avg_rating = $1 WHERE id = $2;",
		newDataRecord.Avg_ratings, newDataRecord.Id)
	utils.CheckRowsAndError(commandTag, &err, 1)
	return newDataRecord
}

func DoesCustomerExists(customerId int32, conn *pgx.Conn) bool {
	return GetCustomerById(customerId, conn) != nil
}
