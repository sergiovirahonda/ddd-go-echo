package service

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"gorm.io/gorm"
)

type InstitutionEntity entity.Institution

func (inst *InstitutionEntity) Get(id uuid.UUID) (*entity.Institution, error) {
	db := infrastructure.DbManager()
	instance := model.Institutions{}
	result := db.Where("id = ?", id).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Institution{
		ID:   instance.ID,
		Name: instance.Name,
	}, nil
}

func (inst *InstitutionEntity) GetByName(
	name string) (*entity.Institution, error) {
	db := infrastructure.DbManager()
	instance := model.Institutions{}
	result := db.Where("name = ?", name).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Institution{
		ID:   instance.ID,
		Name: instance.Name,
	}, nil
}

func (inst *InstitutionEntity) GetAll() ([]entity.Institution, error) {
	db := infrastructure.DbManager()
	institutions := []model.Institutions{}
	result := db.Find(&institutions)
	if result.Error != nil {
		return []entity.Institution{}, nil
	}
	var instances []entity.Institution
	for _, instance := range institutions {
		institutionEntity := entity.Institution{
			ID:   instance.ID,
			Name: instance.Name,
		}
		instances = append(instances, institutionEntity)
	}
	return instances, nil
}

func (inst *InstitutionEntity) Update(
	instance entity.Institution) (*entity.Institution, error) {
	db := infrastructure.DbManager()
	institution := model.Institutions{}
	result := db.Where(
		"id = ?",
		instance.ID,
	).First(&institution)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	institution.Name = instance.Name
	db.Save(&institution)
	return &entity.Institution{
		ID:   institution.ID,
		Name: institution.Name,
	}, nil
}
