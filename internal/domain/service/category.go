package service

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"gorm.io/gorm"
)

type CategoryEntity entity.Category

func (art *CategoryEntity) Get(
	id uuid.UUID,
	institutionId uuid.UUID) (*entity.Category, error) {
	db := infrastructure.DbManager()
	instance := model.Categories{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Category{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		InstitutionID: institutionId,
	}, nil
}

func (art *CategoryEntity) GetByIsin(
	isin string,
	institutionId uuid.UUID) (*entity.Category, error) {
	db := infrastructure.DbManager()
	instance := model.Categories{}
	result := db.Where(
		"isin = ? AND institution_id = ?",
		isin,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Category{
		ID:            instance.ID,
		Isin:          instance.Isin,
		Name:          instance.Name,
		Description:   instance.Description,
		InstitutionID: institutionId,
	}, nil
}

func (art *CategoryEntity) GetAll(
	institutionId uuid.UUID) ([]entity.Category, error) {
	db := infrastructure.DbManager()
	categories := []model.Categories{}
	result := db.Where(
		"institution_id = ?",
		institutionId,
	).Find(&categories)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.Category
	for _, instance := range categories {
		categoryEntity := entity.Category{
			ID:            instance.ID,
			Isin:          instance.Isin,
			Name:          instance.Name,
			Description:   instance.Description,
			InstitutionID: institutionId,
		}
		instances = append(instances, categoryEntity)
	}
	return instances, nil
}

func (art *CategoryEntity) Update(
	instance entity.Category) (*entity.Category, error) {
	db := infrastructure.DbManager()
	category := model.Categories{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		instance.ID,
		instance.InstitutionID,
	).First(&category)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	category.Isin = instance.Isin
	category.Name = instance.Name
	category.Description = instance.Description
	db.Save(&category)
	return &entity.Category{
		ID:            category.ID,
		Isin:          category.Isin,
		Name:          category.Name,
		Description:   category.Description,
		InstitutionID: instance.InstitutionID,
	}, nil
}

func (art *CategoryEntity) Delete(
	id uuid.UUID,
	institutionID uuid.UUID) error {
	db := infrastructure.DbManager()
	category := model.Categories{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionID,
	).First(&category)
	if result.Error != nil {
		return gorm.ErrRecordNotFound
	}
	db.Delete(&category)
	return nil
}
