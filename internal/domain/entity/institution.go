package entity

import (
	"github.com/google/uuid"
)

type Institution struct {
	ID   uuid.UUID
	Name string
}
