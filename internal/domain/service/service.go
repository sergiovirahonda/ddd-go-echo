package service

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"gorm.io/gorm"
)

type ServiceEntity entity.Service

func (serv *ServiceEntity) Get(
	id uuid.UUID,
	institutionId uuid.UUID) (*entity.Service, error) {
	db := infrastructure.DbManager()
	instance := model.Services{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Service{
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

func (serv *ServiceEntity) GetByIsin(
	isin string,
	institutionId uuid.UUID) (*entity.Service, error) {
	db := infrastructure.DbManager()
	instance := model.Services{}
	result := db.Where(
		"isin = ? AND institution_id = ?",
		isin,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.Service{
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

func (serv *ServiceEntity) GetFromCategory(
	categoryId uuid.UUID,
	institutionId uuid.UUID) ([]entity.Service, error) {
	db := infrastructure.DbManager()
	services := []model.Services{}
	result := db.Where(
		"category_id = ? AND institution_id = ?",
		categoryId,
		institutionId,
	).Find(&services)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.Service
	for _, instance := range services {
		serviceEntity := entity.Service{
			ID:            instance.ID,
			Isin:          instance.Isin,
			Name:          instance.Name,
			Description:   instance.Description,
			Unit:          instance.Unit,
			CategoryID:    instance.CategoryID,
			DefaultPrice:  instance.DefaultPrice,
			InstitutionID: institutionId,
		}
		instances = append(instances, serviceEntity)
	}
	return instances, nil
}

func (serv *ServiceEntity) GetAll(
	institutionId uuid.UUID) ([]entity.Service, error) {
	db := infrastructure.DbManager()
	services := []model.Services{}
	result := db.Where(
		"institution_id = ?",
		institutionId,
	).Find(&services)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.Service
	for _, instance := range services {
		serviceEntity := entity.Service{
			ID:            instance.ID,
			Isin:          instance.Isin,
			Name:          instance.Name,
			Description:   instance.Description,
			Unit:          instance.Unit,
			CategoryID:    instance.CategoryID,
			DefaultPrice:  instance.DefaultPrice,
			InstitutionID: institutionId,
		}
		instances = append(instances, serviceEntity)
	}
	return instances, nil
}

func (serv *ServiceEntity) Update(
	instance entity.Service) (*entity.Service, error) {
	db := infrastructure.DbManager()
	service := model.Services{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		instance.ID,
		instance.InstitutionID,
	).First(&service)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	service.Isin = instance.Isin
	service.Name = instance.Name
	service.Description = instance.Description
	service.Unit = instance.Unit
	service.CategoryID = instance.CategoryID
	service.DefaultPrice = instance.DefaultPrice
	db.Save(&service)
	return &entity.Service{
		ID:            service.ID,
		Isin:          service.Isin,
		Name:          service.Name,
		Description:   service.Description,
		Unit:          service.Unit,
		CategoryID:    service.CategoryID,
		DefaultPrice:  service.DefaultPrice,
		InstitutionID: instance.InstitutionID,
	}, nil
}

func (serv *ServiceEntity) Delete(
	id uuid.UUID,
	institutionID uuid.UUID) error {
	db := infrastructure.DbManager()
	service := model.Services{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		id,
		institutionID,
	).First(&service)
	if result.Error != nil {
		return gorm.ErrRecordNotFound
	}
	db.Delete(&service)
	return nil
}
