package factory

import (
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
)

type InstitutionEntity entity.Institution

type InstitutionFactory interface {
	Create(entity InstitutionEntity) entity.Institution
}

func (institution *InstitutionEntity) Create(
	inst entity.Institution) *entity.Institution {
	db := infrastructure.DbManager()
	instance := model.Institutions{
		ID:   inst.ID,
		Name: inst.Name,
	}
	db.Create(&instance)
	return &entity.Institution{
		ID:   instance.ID,
		Name: instance.Name,
	}
}
