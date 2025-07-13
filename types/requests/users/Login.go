package types

import (
	"time"
	usrs "types/database/users"
)

type LoginRequestBody struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lrb *LoginRequestBody) ToRecord() *usrs.UsersRecord {
	if lrb == nil {
		return nil
	}
	return &usrs.UsersRecord{
		Id:        0,
		Login:     lrb.Login,
		Password:  lrb.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
