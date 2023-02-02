package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Password    string
	Institution uuid.UUID `gorm:"type:uuid;"`
	Role        string    `json:"role"`
	Active      bool      `json:"active"`
}
