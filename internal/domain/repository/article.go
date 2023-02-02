package repository

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
)

type ArticleRepository interface {
	Get(ID uuid.UUID) (*entity.Article, error)
	GetByIsin(isin string) (*entity.Article, error)
	GetAll() (*[]entity.Article, error)
	Update(article entity.Article) (*entity.Article, error)
	Delete(article entity.Article) error
}
