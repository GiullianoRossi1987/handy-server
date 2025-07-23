package services

import (
	"context"
	"fmt"
	usr "pkg/users"
	"time"
	requests "types/requests/users"
	responses "types/responses/users"
	"utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByLogin(pool *pgxpool.Pool, login string) (*responses.UserResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetUserByLogin(login, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
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
		Success:        utils.ValidatePassword(rq.Password, data.Password),
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
	record.Password, err = utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	id, err := usr.AddUser(*record, conn)
	if err != nil {
		return nil, err
	}
	conn.Release()
	return id, nil
}

func UpdateUser(pool *pgxpool.Pool, req requests.CreateUserRequest, id int) (*int, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	record.Id = id
	fmt.Println(record.Id)
	record.Password, err = utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	if err := usr.UpdateUserById(*record, conn); err != nil {
		return nil, err
	}
	conn.Release()
	return &id, nil
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

func GetUserById(pool *pgxpool.Pool, id int) (*responses.UserResponseBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := usr.GetUserById(id, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return responses.SerializeUserResponse(data), nil
}
