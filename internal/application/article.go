package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/factory"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/service"
	"gorm.io/gorm"
)

type Article struct {
	Isin         string  `json:"isin"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Unit         string  `json:"unit"`
	CategoryID   string  `json:"category_id"`
	DefaultPrice float32 `json:"default_price"`
}

func GetArticles(user entity.User) (*[]entity.Article, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ArticleEntity{}
	entities, err := repository.GetAll(institution.ID)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func GetArticlesFromCategory(
	categoryID uuid.UUID,
	user entity.User,
) (*[]entity.Article, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ArticleEntity{}
	entities, err := repository.GetFromCategory(categoryID, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entities, nil
}

func GetArticleById(id uuid.UUID, user entity.User) (*entity.Article, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ArticleEntity{}
	entity, err := repository.Get(id, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func GetArticleByIsin(isin string, user entity.User) (*entity.Article, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ArticleEntity{}
	entity, err := repository.GetByIsin(isin, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func CreateArticle(article Article, user entity.User) (*entity.Article, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ArticleEntity{}
	factory := factory.ArticleEntity{}
	exists, err := repository.GetByIsin(article.Isin, institution.ID)
	if err != nil {
		categoryID, err := uuid.Parse(article.CategoryID)
		if err != nil {
			return nil, errors.New("invalid category id")
		}
		// still missing category validation
		entity := entity.Article{
			ID:            uuid.New(),
			Isin:          article.Isin,
			Name:          article.Name,
			Description:   article.Description,
			Unit:          article.Unit,
			CategoryID:    categoryID,
			DefaultPrice:  article.DefaultPrice,
			InstitutionID: institution.ID,
		}
		instance := factory.Create(entity)
		return instance, nil
	}
	errorMessage := "isin already exists for ID:" + exists.ID.String()
	return nil, errors.New(errorMessage)
}

func DeleteArticleById(id uuid.UUID, user entity.User) error {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return err
	}
	repository := service.ArticleEntity{}
	err = repository.Delete(id, institution.ID)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return nil
}
