package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/factory"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/service"
	"gorm.io/gorm"
)

type Category struct {
	Isin        string `json:"isin"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetCategories(user entity.User) (*[]entity.Category, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.CategoryEntity{}
	entities, err := repository.GetAll(institution.ID)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func GetCategoryById(id uuid.UUID, user entity.User) (*entity.Category, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.CategoryEntity{}
	entity, err := repository.Get(id, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func GetCategoryByIsin(isin string, user entity.User) (*entity.Category, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.CategoryEntity{}
	entity, err := repository.GetByIsin(isin, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func CreateCategory(category Category, user entity.User) (*entity.Category, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.CategoryEntity{}
	factory := factory.CategoryEntity{}
	exists, err := repository.GetByIsin(category.Isin, institution.ID)
	if err != nil {
		entity := entity.Category{
			ID:            uuid.New(),
			Isin:          category.Isin,
			Name:          category.Name,
			Description:   category.Description,
			InstitutionID: institution.ID,
		}
		instance := factory.Create(entity)
		return instance, nil
	}
	errorMessage := "isin already exists for ID:" + exists.ID.String()
	return nil, errors.New(errorMessage)
}

func DeleteCategoryById(id uuid.UUID, user entity.User) error {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return err
	}
	categoryRepository := service.CategoryEntity{}
	category, err := categoryRepository.Get(id, institution.ID)
	if err != nil {
		return err
	}
	articles, err := GetArticlesFromCategory(category.ID, user)
	if err != nil {
		articles = &[]entity.Article{}
	}
	for _, article := range *articles {
		DeleteArticleById(article.ID, user)
	}

	err = categoryRepository.Delete(id, institution.ID)
	if err != nil {
		return err
	}
	return nil
}
