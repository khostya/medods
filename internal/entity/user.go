package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"_id"`
	Password string    `json:"password"`
	Username string    `json:"username"`
}
