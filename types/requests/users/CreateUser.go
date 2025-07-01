package types

type CreateUserRequest struct {
	Login         string `json:"login" binding:"required"`
	Password      string `json:"password" binding:"required"` // THE ENCODING SHOULD HAPPEN IN THE CLIENT SIDE TO AVOID ANNOYANCE
	AsProprietary bool   `json:"as" binding:"required"`
}
