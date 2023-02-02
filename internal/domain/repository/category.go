package repository

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
)

type CategoryRepository interface {
	Get(ID uuid.UUID) (*entity.Category, error)
	GetByIsin(isin string) (*entity.Category, error)
	GetAll() (*[]entity.Category, error)
	Update(category entity.Category) (*entity.Category, error)
	Delete(category entity.Category) error
}
