package services

import (
	usr "pkg/users"
	requests "types/requests/users"
	responses "types/responses/users"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO Change usage of ID to UUID's in API based queries
func GetCustomerById(pool *pgxpool.Pool, id int32) (*responses.CustomerResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetCustomerById(id, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	conn.Release()
	return responses.SerializeCustomerResponse(data), nil
}

func GetCustomerByUUID(pool *pgxpool.Pool, uuid string) (*responses.CustomerResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetCustomerByUUID(uuid, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	conn.Release()
	return responses.SerializeCustomerResponse(data), nil
}

func AddCustomer(pool *pgxpool.Pool, req requests.UpdateUserRequest, usr_id int32) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToCustomerRecord()
	record.UserId = int(usr_id)
	id, err := usr.AddCustomer(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
}

func UpdateCustomer(pool *pgxpool.Pool, req requests.UpdateUserRequest, uuid string) (*responses.CustomerResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetCustomerByUUID(uuid, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	record := req.ToCustomerRecord()
	record.Id = data.Id
	record.UUID = uuid
	if err := usr.UpdateCustomer(*record, conn); err != nil {
		return nil, err
	}
	conn.Release()
	return GetCustomerByUUID(pool, uuid)
}

func DeactivateCustomer(pool *pgxpool.Pool, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := usr.DeactivateCustomer(uuid, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteCustomer(pool *pgxpool.Pool, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := usr.DeleteCustomer(uuid, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
