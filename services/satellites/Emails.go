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

func GetEmailById(pool *pgxpool.Pool, emailId int32) (*serial.Email, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record, err := satellites.GetEmailById(emailId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return serial.SerializeEmail(record), nil
}

func AddEmail(pool *pgxpool.Pool, rq requests.EmailBody) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := rq.ToRecord()
	if err := satellites.AddEmail(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func GetWorkerEmails(pool *pgxpool.Pool, uuid string) ([]serial.Email, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetWorkerEmails(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.EmailRecord) serial.Email {
			return *serial.SerializeEmail(&i)
		},
	)
	return serialzed_data, nil
}

func GetCustomerEmails(pool *pgxpool.Pool, uuid string) ([]serial.Email, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := satellites.GetCustomerEmails(uuid, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	serialzed_data := utils.MapCar(
		data,
		func(i types.EmailRecord) serial.Email {
			return *serial.SerializeEmail(&i)
		},
	)
	return serialzed_data, nil
}

// NOTE replace item exists verification to controller/route level
func UpdateEmail(pool *pgxpool.Pool, rq requests.EmailBody, emailId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := rq.ToRecord()
	record.Id = emailId
	if err := satellites.UpdateEmail(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteEmail(pool *pgxpool.Pool, emailId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := satellites.DeleteEmail(emailId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
