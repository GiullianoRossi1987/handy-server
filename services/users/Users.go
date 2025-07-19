package services

import (
	"context"
	usr "pkg/users"
	"time"
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
	conn.Release()
	return responses.SerializeUserResponse(data), nil
}

func Login(pool *pgxpool.Pool, rq requests.LoginRequestBody) (*responses.LoginResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetUserByLogin(rq.Login, conn)
	if err != nil {
		return nil, err
	}
	loggable, err := usr.IsUserLoggable(int32(data.Id), conn)
	if err != nil {
		return nil, err
	}
	if data == nil && !loggable {
		return nil, nil
	}
	conn.Release()
	response := responses.LoginResponse{
		Login:          rq.Login,
		Success:        data.Password == rq.Password,
		AttemptedLogin: time.Now(),
	}
	return &response, nil
}

func AddUser(pool *pgxpool.Pool, req requests.CreateUserRequest) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	id, err := usr.AddUser(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
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
	data, err := usr.GetUserByLogin(record.Login, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return responses.SerializeUserResponse(data), nil
}
