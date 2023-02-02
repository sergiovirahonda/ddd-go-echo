package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Isin          string    `json:"isin"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	InstitutionID uuid.UUID `gorm:"type:uuid"`
}
