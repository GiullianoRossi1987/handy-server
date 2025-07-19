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

func GetPhoneById(pool *pgxpool.Pool, phoneId int32) (*serial.Phone, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record, err := satellites.GetPhoneById(phoneId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return serial.SerializePhone(record), nil
}

func AddPhone(pool *pgxpool.Pool, rq requests.PhoneBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := rq.ToRecord()
	id, err := satellites.AddPhone(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
}

func GetWorkerPhones(pool *pgxpool.Pool, uuid string) ([]serial.Phone, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetWorkerPhones(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.PhoneRecord) serial.Phone {
			return *serial.SerializePhone(&i)
		},
	)
	return serialzed_data, nil
}

func GetCustomerPhones(pool *pgxpool.Pool, uuid string) ([]serial.Phone, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetCustomerPhones(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.PhoneRecord) serial.Phone {
			return *serial.SerializePhone(&i)
		},
	)
	return serialzed_data, nil
}

// NOTE replace item exists verification to controller/route level
func UpdatePhone(pool *pgxpool.Pool, rq requests.PhoneBody, emailId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := rq.ToRecord()
	record.Id = emailId
	if err := satellites.UpdatePhone(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeletePhone(pool *pgxpool.Pool, phoneId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := satellites.DeletePhone(phoneId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
