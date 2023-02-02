package factory

import (
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
)

type CategoryEntity entity.Category

type CategoryFactory interface {
	Create(entity CategoryEntity) entity.Category
}

func (category *CategoryEntity) Create(cat entity.Category) *entity.Category {
	db := infrastructure.DbManager()
	instance := model.Categories{
		ID:            cat.ID,
		Isin:          cat.Isin,
		Name:          cat.Name,
		Description:   cat.Description,
		InstitutionID: cat.InstitutionID,
	}
	db.Create(&instance)
	return &entity.Category{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		InstitutionID: instance.InstitutionID,
	}
}
