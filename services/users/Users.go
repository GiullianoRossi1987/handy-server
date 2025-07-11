package services

import (
	"context"
	usr "pkg/users"
	requests "types/requests/users"
	responses "types/responses/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByLogin(pool *pgxpool.Pool, login string) (*responses.UserResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetUserByLogin(login, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	conn.Release()
	return responses.SerializeUserResponse(data), nil
}

func AddUser(pool *pgxpool.Pool, req requests.CreateUserRequest) (*responses.UserResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	if err := usr.AddUser(*record, conn); err != nil {
		return nil, err
	}
	data, err := usr.GetUserByLogin(record.Password, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	conn.Release()
	return responses.SerializeUserResponse(data), nil
}

func UpdateUser(pool *pgxpool.Pool, req requests.CreateUserRequest, id int) (*responses.UserResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	record.Id = id
	if err := usr.UpdateUserById(*record, conn); err != nil {
		return nil, err
	}
	data, err := usr.GetUserByLogin(record.Password, conn)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	conn.Release()
	return responses.SerializeUserResponse(data), nil
}

func DeleteUser(pool *pgxpool.Pool, id int) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := usr.DeleteUserById(id, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
