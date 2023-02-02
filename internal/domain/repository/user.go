package repository

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
)

type UserRepository interface {
	Get(ID uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() (*[]entity.User, error)
	Update(article entity.Article) (*entity.User, error)
	Delete(article entity.User) error
}
