package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID
	Email       string
	FirstName   string
	LastName    string
	Password    string
	Institution uuid.UUID
	Role        string
	Active      bool
}
