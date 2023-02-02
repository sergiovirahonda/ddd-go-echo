package factory

import (
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
)

type ArticleEntity entity.Article

type ArticleFactory interface {
	Create(entity ArticleEntity) entity.Article
}

func (article *ArticleEntity) Create(art entity.Article) *entity.Article {
	db := infrastructure.DbManager()
	instance := model.Articles{
		ID:            art.ID,
		Isin:          art.Isin,
		Name:          art.Name,
		Description:   art.Description,
		Unit:          art.Unit,
		CategoryID:    art.CategoryID,
		DefaultPrice:  art.DefaultPrice,
		InstitutionID: art.InstitutionID,
	}
	db.Create(&instance)
	return &entity.Article{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		Unit:          instance.Unit,
		CategoryID:    instance.CategoryID,
		DefaultPrice:  instance.DefaultPrice,
		InstitutionID: instance.InstitutionID,
	}
}
