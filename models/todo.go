package models

import "github.com/google/uuid"

type Todo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"not null"`
	Completed bool      `json:"completed" gorm:"default:false"`
	User_uid  uuid.UUID `json:"user_id" gorm:"foreignkey:ID;type:uuid;not null"`
}
