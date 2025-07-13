package types

import (
	"time"
	types "types/database/users"
)

type CreateUserRequest struct {
	Login         string // `json:"login" binding:"required"`
	Password      string // `json:"password" binding:"required"` // THE ENCODING SHOULD HAPPEN IN THE CLIENT SIDE TO AVOID ANNOYANCE
	AsProprietary bool   // `json:"as" binding:"required"` THIS VARIABLE IS ONLY USEFUL INSIDE THE CLIENT, IT CAN BE SENT HERE TO REDUCE CLIENT SIDE OPERATIONS
}

func (rq *CreateUserRequest) ToRecord() *types.UsersRecord {
	if rq == nil {
		return nil
	}
	return &types.UsersRecord{
		Id:        0,
		Login:     rq.Login,
		Password:  rq.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
