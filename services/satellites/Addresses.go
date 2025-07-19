package services

import (
	satellites "pkg/satellites"
	types "types/database/satellites"
	requests "types/requests/satellites"
	serial "types/serializables"
	"utils"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAddressById(pool *pgxpool.Pool, addrId int32) (*serial.Address, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record, err := satellites.GetAddressById(addrId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return serial.SerializeAddress(record), nil
}

func AddAddress(pool *pgxpool.Pool, rq requests.AddressBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := rq.ToRecord()
	id, err := satellites.AddAddress(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
}

func GetWorkerAddresses(pool *pgxpool.Pool, uuid string) ([]serial.Address, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetWorkerAddresses(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.AddressRecord) serial.Address {
			return *serial.SerializeAddress(&i)
		},
	)
	return serialzed_data, nil
}

func GetCustomerAddresses(pool *pgxpool.Pool, uuid string) ([]serial.Address, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetCustomerAddresses(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.AddressRecord) serial.Address {
			return *serial.SerializeAddress(&i)
		},
	)
	return serialzed_data, nil
}

// NOTE replace item exists verification to controller/route level
func UpdateAddress(pool *pgxpool.Pool, rq requests.AddressBody, addressId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := rq.ToRecord()
	record.Id = addressId
	if err := satellites.UpdatedAddress(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteAddress(pool *pgxpool.Pool, addressId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := satellites.DeleteAddress(addressId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
