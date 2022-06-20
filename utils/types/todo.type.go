package types

import (
	"github.com/google/uuid"
)

type TodoType struct {
	Title    string    `json:"title"`
	User_uid uuid.UUID `json:"user_id"`
}

type TodoUpdate struct {
	Title     string `json:"title"`
	ID        uint   `json:"id"`
	Completed bool   `json:"completed"`
}
