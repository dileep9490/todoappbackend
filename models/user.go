package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;unique;primarykey"`
	Name     string    `json:"name" `
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password"`
	
}

