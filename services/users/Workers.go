package services

import (
	"context"
	usr "pkg/users"
	requests "types/requests/users"
	responses "types/responses/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkerByUUID(pool *pgxpool.Pool, uuid string) (*responses.WorkerResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetWorkerByUUID(uuid, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	return responses.SerializeWorkerResponse(data), nil
}

func AddWorker(pool *pgxpool.Pool, req requests.UpdateUserRequest, usrid int32) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToWorkerRecord()
	record.UserId = int(usrid)
	id, err := usr.AddWorker(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
}

func DeactivateWorker(pool *pgxpool.Pool, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := usr.DeactivateWorker(uuid, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteWorker(pool *pgxpool.Pool, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := usr.DeleteWorker(uuid, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func UpdateWorker(pool *pgxpool.Pool, req requests.UpdateUserRequest, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	data, err := usr.GetWorkerByUUID(uuid, conn)
	if err != nil {
		return err
	}
	if data == nil {
		return nil
	}
	record := req.ToWorkerRecord()
	record.Id = data.Id
	record.UUID = uuid
	if err := usr.UpdateWorker(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
