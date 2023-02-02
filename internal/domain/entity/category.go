package entity

import (
	"github.com/google/uuid"
)

type Category struct {
	ID            uuid.UUID
	Isin          string
	Name          string
	Description   string
	InstitutionID uuid.UUID
}
