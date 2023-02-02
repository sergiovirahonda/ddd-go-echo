package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/factory"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/service"
	"gorm.io/gorm"
)

type Service struct {
	Isin         string  `json:"isin"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Unit         string  `json:"unit"`
	CategoryID   string  `json:"category_id"`
	DefaultPrice float32 `json:"default_price"`
}

func GetServices(user entity.User) (*[]entity.Service, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ServiceEntity{}
	entities, err := repository.GetAll(institution.ID)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func GetServicesFromCategory(
	categoryID uuid.UUID,
	user entity.User,
) (*[]entity.Service, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ServiceEntity{}
	entities, err := repository.GetFromCategory(categoryID, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entities, nil
}

func GetServiceById(id uuid.UUID, user entity.User) (*entity.Service, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ServiceEntity{}
	entity, err := repository.Get(id, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func GetServiceByIsin(isin string, user entity.User) (*entity.Service, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ServiceEntity{}
	entity, err := repository.GetByIsin(isin, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func CreateService(serv Service, user entity.User) (*entity.Service, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.ServiceEntity{}
	factory := factory.ServiceEntity{}
	exists, err := repository.GetByIsin(serv.Isin, institution.ID)
	if err != nil {
		categoryID, err := uuid.Parse(serv.CategoryID)
		if err != nil {
			return nil, errors.New("invalid category id")
		}
		// still missing category validation
		entity := entity.Service{
			ID:            uuid.New(),
			Isin:          serv.Isin,
			Name:          serv.Name,
			Description:   serv.Description,
			Unit:          serv.Unit,
			CategoryID:    categoryID,
			DefaultPrice:  serv.DefaultPrice,
			InstitutionID: institution.ID,
		}
		instance := factory.Create(entity)
		return instance, nil
	}
	errorMessage := "isin already exists for ID:" + exists.ID.String()
	return nil, errors.New(errorMessage)
}

func DeleteServiceById(id uuid.UUID, user entity.User) error {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return err
	}
	repository := service.ServiceEntity{}
	err = repository.Delete(id, institution.ID)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return nil
}
