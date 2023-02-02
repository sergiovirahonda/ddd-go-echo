package service

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"gorm.io/gorm"
)

type ArticleEntity entity.Article

func (art *ArticleEntity) Get(
	id uuid.UUID,
	institutionId uuid.UUID) (*entity.Article, error) {
	db := infrastructure.DbManager()
	instance := model.Articles{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Article{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		Unit:          instance.Unit,
		CategoryID:    instance.CategoryID,
		DefaultPrice:  instance.DefaultPrice,
		InstitutionID: institutionId,
	}, nil
}

func (art *ArticleEntity) GetByIsin(
	isin string,
	institutionId uuid.UUID) (*entity.Article, error) {
	db := infrastructure.DbManager()
	instance := model.Articles{}
	result := db.Where(
		"isin = ? AND institution_id = ?",
		isin,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Article{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		Unit:          instance.Unit,
		CategoryID:    instance.CategoryID,
		DefaultPrice:  instance.DefaultPrice,
		InstitutionID: institutionId,
	}, nil
}

func (art *ArticleEntity) GetFromCategory(
	categoryId uuid.UUID,
	institutionId uuid.UUID) ([]entity.Article, error) {
	db := infrastructure.DbManager()
	articles := []model.Articles{}
	result := db.Where(
		"category_id = ? AND institution_id = ?",
		categoryId,
		institutionId,
	).Find(&articles)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.Article
	for _, instance := range articles {
		articleEntity := entity.Article{
			ID:            instance.ID,
			Isin:          instance.Isin,
			Name:          instance.Name,
			Description:   instance.Description,
			Unit:          instance.Unit,
			CategoryID:    instance.CategoryID,
			DefaultPrice:  instance.DefaultPrice,
			InstitutionID: institutionId,
		}
		instances = append(instances, articleEntity)
	}
	return instances, nil
}

func (art *ArticleEntity) GetAll(
	institutionId uuid.UUID) ([]entity.Article, error) {
	db := infrastructure.DbManager()
	articles := []model.Articles{}
	result := db.Where(
		"institution_id = ?",
		institutionId,
	).Find(&articles)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.Article
	for _, instance := range articles {
		articleEntity := entity.Article{
			ID:            instance.ID,
			Isin:          instance.Isin,
			Name:          instance.Name,
			Description:   instance.Description,
			Unit:          instance.Unit,
			CategoryID:    instance.CategoryID,
			DefaultPrice:  instance.DefaultPrice,
			InstitutionID: institutionId,
		}
		instances = append(instances, articleEntity)
	}
	return instances, nil
}

func (art *ArticleEntity) Update(
	instance entity.Article) (*entity.Article, error) {
	db := infrastructure.DbManager()
	article := model.Articles{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		instance.ID,
		instance.InstitutionID,
	).First(&article)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	article.Isin = instance.Isin
	article.Name = instance.Name
	article.Description = instance.Description
	article.Unit = instance.Unit
	article.CategoryID = instance.CategoryID
	article.DefaultPrice = instance.DefaultPrice
	db.Save(&article)
	return &entity.Article{
		ID:            article.ID,
		Isin:          article.Isin,
		Name:          article.Name,
		Description:   article.Description,
		Unit:          article.Unit,
		CategoryID:    article.CategoryID,
		DefaultPrice:  article.DefaultPrice,
		InstitutionID: instance.InstitutionID,
	}, nil
}

func (art *ArticleEntity) Delete(
	id uuid.UUID,
	institutionID uuid.UUID) error {
	db := infrastructure.DbManager()
	article := model.Articles{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionID,
	).First(&article)
	if result.Error != nil {
		return gorm.ErrRecordNotFound
	}
	db.Delete(&article)
	return nil
}
