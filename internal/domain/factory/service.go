package factory

import (
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
)

type ServiceEntity entity.Service

type ServiceFactory interface {
	Create(entity ServiceEntity) entity.Service
}

func (service *ServiceEntity) Create(serv entity.Service) *entity.Service {
	db := infrastructure.DbManager()
	instance := model.Services{
		ID:            serv.ID,
		Isin:          serv.Isin,
		Name:          serv.Name,
		Description:   serv.Description,
		Unit:          serv.Unit,
		CategoryID:    serv.CategoryID,
		DefaultPrice:  serv.DefaultPrice,
		InstitutionID: serv.InstitutionID,
	}
	db.Create(&instance)
	return &entity.Service{
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
