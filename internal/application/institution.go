package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/factory"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/service"
	"gorm.io/gorm"
)

type Institution struct {
	Name string `json:"name"`
}

func GetInstitutions(user entity.User) (*[]entity.Institution, error) {
	repository := service.InstitutionEntity{}
	entities, err := repository.GetAll()
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func GetInstitutionById(id uuid.UUID, user entity.User) (*entity.Category, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.CategoryEntity{}
	entity, err := repository.Get(id, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	if institution.ID != entity.ID {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func CreateInstitution(
	institution Institution,
	user entity.User) (*entity.Institution, error) {
	repository := service.InstitutionEntity{}
	factory := factory.InstitutionEntity{}
	exists, err := repository.GetByName(institution.Name)
	if err != nil {
		entity := entity.Institution{
			ID:   uuid.New(),
			Name: institution.Name,
		}
		instance := factory.Create(entity)
		return instance, nil
	}
	errorMessage := "name already exists for ID:" + exists.ID.String()
	return nil, errors.New(errorMessage)
}

func GetInstitutionFromUser(user entity.User) (*entity.Institution, error) {
	institutionId := user.Institution
	repository := service.InstitutionEntity{}
	institution, err := repository.Get(institutionId)
	if err != nil {
		return nil, err
	}
	return institution, nil
}
