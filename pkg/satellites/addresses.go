package pkg

import (
	"context"
	types "types/database/satellites"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetCustomerAddresses(customerId int32, conn *pgx.Conn) []types.AddressRecord {
	rows, err := conn.Query(context.Background(), "SELECT * FROM addresses WHERE id_customer = $1;", customerId)
	if err != nil {
		panic(err)
	}
	addresses, errCollect := pgx.CollectRows(rows, pgx.RowTo[types.AddressRecord])
	if errCollect != nil {
		panic(errCollect)
	}
	return addresses
}

func GetWorkerAddresses(workerId int32, conn *pgx.Conn) []types.AddressRecord {
	rows, err := conn.Query(context.Background(), "SELECT * FROM addresses WHERE id_worker = $1;", workerId)
	if err != nil {
		panic(err)
	}
	addresses, errCollect := pgx.CollectRows(rows, pgx.RowTo[types.AddressRecord])
	if errCollect != nil {
		panic(errCollect)
	}
	return addresses
}

func GetAddressById(addressId int32, conn *pgx.Conn) *types.AddressRecord {
	var row *types.AddressRecord
	err := conn.QueryRow(context.Background(), "SELECT * FROM addresses WHERE id = $1;", addressId).Scan(&row)
	if err != nil {
		panic(err)
	}
	return row
}

func AddAddress(address types.AddressRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO addresses (id_worker, id_customer, address, address_number, city, uf, country) VALUES ($1,$2,$3,$4,$5,$6);",
		address.IdWorker, address.IdCustomer, address.Address, address.AddressNumber, address.City, address.UF, address.Country)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteAddress(addressId int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM addresses WHERE id = $1;",
		addressId)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdatedAddress(address types.AddressRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE addresses id_worker = $1, id_customer = $2, address = $3, address_number = $4, city = $5, uf = $6, country = $8 WHERE id = $8;",
		address.IdWorker, address.IdCustomer, address.Address, address.AddressNumber, address.City, address.UF, address.Country)
	utils.CheckRowsAndError(commandTag, &err, 1)
}
