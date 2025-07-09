package services

import (
	usr "pkg/users"
	"services"
	requests "types/requests/users"
	responses "types/responses/users"

	"github.com/google/uuid"
)

func GetCustomerById(id int32) (*responses.CustomerResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	return responses.SerializeCustomerResponse(*data), nil
}

func GetCustomerByUUID(uuid string) (*responses.CustomerResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	return responses.SerializeCustomerResponse(*data), nil
}

func AddCustomer(req requests.UpdateUserRequest) (*responses.CustomerResponseBody, error) {
	conn, err := services.GetConnByEnv()
	if err != nil {
		return nil, err
	}
	record := req.ToCustomerRecord()
	record.UUID = uuid.NewString()
	if err := usr.AddCustomer(*record, conn); err != nil {
		return nil, err
	}
	conn.Close()
	return GetCustomerByUUID(record.UUID)
}

func UpdateCustomer(req requests.UpdateUserRequest, uuid string) (*responses.CustomerResponseBody, error) {
	conn, err := services.GetConnByEnv()
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
	conn.Close()
	return GetCustomerByUUID(uuid)
}
