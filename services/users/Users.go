package services

import (
	usr "pkg/users"
	"services"
	requests "types/requests/users"
	responses "types/responses/users"
)

func GetUserByLogin(login string) (*responses.UserResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	conn.Close()
	return responses.SerializeUserResponse(data), nil
}

func AddUser(req requests.CreateUserRequest) (*responses.UserResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	conn.Close()
	return responses.SerializeUserResponse(data), nil
}

func UpdateUser(req requests.CreateUserRequest, id int) (*responses.UserResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	conn.Close()
	return responses.SerializeUserResponse(data), nil
}

func DeleteUser(id int) error {
	conn, err := services.GetConnByEnv()
	if err != nil {
		return err
	}
	if err := usr.DeleteUserById(id, conn); err != nil {
		return err
	}
	conn.Close()
	return nil
}
