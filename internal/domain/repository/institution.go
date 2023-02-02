package repository

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
)

type InstitutionRepository interface {
	Get(ID uuid.UUID) (*entity.Institution, error)
	GetAll() (*[]entity.Institution, error)
	Update(institution entity.Institution) (*entity.Institution, error)
}
