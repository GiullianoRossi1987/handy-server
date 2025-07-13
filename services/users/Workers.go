package services

import (
	"context"
	usr "pkg/users"
	requests "types/requests/users"
	responses "types/responses/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO Implement proper satellite exclusion for user exclusion
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

func AddWorker(pool *pgxpool.Pool, req requests.UpdateUserRequest) (*responses.WorkerResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToWorkerRecord()
	if err := usr.AddWorker(*record, conn); err != nil {
		return nil, err
	}
	conn.Release()
	return GetWorkerByUUID(pool, record.UUID)
}

func DeleteWorker(pool *pgxpool.Pool, uuid string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	// if err := usr.DeleteWorker()
	// [ ] Implement DeleteWorker by UUID
	conn.Release()
	return nil
}
