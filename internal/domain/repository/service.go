package repository

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
)

type ServiceRepository interface {
	Get(ID uuid.UUID) (*entity.Service, error)
	GetByIsin(isin string) (*entity.Service, error)
	GetAll() (*[]entity.Service, error)
	Update(article entity.Service) (*entity.Service, error)
	Delete(article entity.Service) error
}
