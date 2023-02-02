package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Isin          string    `json:"isin"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Unit          string    `json:"unit"`
	CategoryID    uuid.UUID `gorm:"type:uuid;"`
	DefaultPrice  float32   `gorm:"type:float;default:null"`
	InstitutionID uuid.UUID `gorm:"type:uuid"`
}
