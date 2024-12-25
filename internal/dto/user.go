package dto

type CreateUser struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}
