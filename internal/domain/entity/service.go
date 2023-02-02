package entity

import (
	"github.com/google/uuid"
)

type Service struct {
	ID            uuid.UUID
	Isin          string
	Name          string
	Description   string
	Unit          string
	CategoryID    uuid.UUID
	DefaultPrice  float32
	InstitutionID uuid.UUID
}
