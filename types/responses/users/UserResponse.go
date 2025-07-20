package types

import (
	"time"
	db "types/database/users"
)

type UserResponseBody struct {
	Id        int       `json:"id" binding:"required"`
	Login     string    `json:"login" binding:"required"`
	Password  string    `json:"password,omitempty" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

func SerializeUserResponse(record *db.UsersRecord) *UserResponseBody {
	if record == nil {
		return nil
	}
	return &UserResponseBody{
		Id:        record.Id,
		Login:     record.Login,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}
