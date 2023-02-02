package service

import (
	"github.com/google/uuid"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"gorm.io/gorm"
)

type UserEntity entity.User

func (u *UserEntity) Get(
	id uuid.UUID,
	institutionId uuid.UUID) (*entity.User, error) {
	db := infrastructure.DbManager()
	instance := model.Users{}
	result := db.Where(
		"id = ? AND institution = ?",
		id,
		institutionId,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.User{
		ID:          instance.ID,
		Email:       instance.Email,
		FirstName:   instance.FirstName,
		LastName:    instance.LastName,
		Password:    instance.Password,
		Institution: instance.Institution,
		Role:        instance.Role,
		Active:      instance.Active,
	}, nil
}

func (u *UserEntity) GetByEmail(email string) (*entity.User, error) {
	db := infrastructure.DbManager()
	instance := model.Users{}
	result := db.Where(
		"email = ?",
		email,
	).First(&instance)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &entity.User{
		ID:          instance.ID,
		Email:       instance.Email,
		FirstName:   instance.FirstName,
		LastName:    instance.LastName,
		Password:    instance.Password,
		Institution: instance.Institution,
		Role:        instance.Role,
		Active:      instance.Active,
	}, nil
}

func (u *UserEntity) GetAll(
	institutionId uuid.UUID) ([]entity.User, error) {
	db := infrastructure.DbManager()
	users := []model.Users{}
	result := db.Where(
		"institution = ?",
		institutionId,
	).Find(&users)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var instances []entity.User
	for _, instance := range users {
		userEntity := entity.User{
			ID:          instance.ID,
			Email:       instance.Email,
			FirstName:   instance.FirstName,
			LastName:    instance.LastName,
			Password:    instance.Password,
			Institution: instance.Institution,
			Role:        instance.Role,
			Active:      instance.Active,
		}
		instances = append(instances, userEntity)
	}
	return instances, nil
}

func (u *UserEntity) Update(
	instance entity.User) (*entity.User, error) {
	db := infrastructure.DbManager()
	user := model.Users{}
	result := db.Where(
		"id = ? AND institution_id = ?",
		instance.ID,
		instance.Institution,
	).First(&user)
	if result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	user.Email = instance.Email
	user.FirstName = instance.FirstName
	user.LastName = instance.LastName
	user.Password = instance.Password
	user.Institution = instance.Institution
	user.Role = instance.Role
	user.Active = instance.Active
	db.Save(&user)
	return &entity.User{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Password:    user.Password,
		Institution: user.Institution,
		Role:        user.Role,
		Active:      user.Active,
	}, nil
}

func (u *UserEntity) Delete(
	id uuid.UUID,
	institutionID uuid.UUID) error {
	db := infrastructure.DbManager()
	user := model.Users{}
	result := db.Where(
		"id = ? AND institution = ?",
		id,
		institutionID,
	).First(&user)
	if result.Error != nil {
		return gorm.ErrRecordNotFound
	}
	db.Delete(&user)
	return nil
}
