package types

import (
	"encoding/json"
	"time"
)

type LoginResponse struct {
	Login          string    `json:"login" binding:"required"`
	Success        bool      `json:"success" binding:"required"` // True w/ 200 Status | False with 401
	UserId         int       `json:"id" binding:"required"`
	AttemptedLogin time.Time `json:"attempted_at" binding:"required"`
}

func (lr *LoginResponse) ToJSON() (string, error) {
	val, err := json.Marshal(lr)
	if err != nil {
		return "", nil
	}
	return string(val), nil
}
