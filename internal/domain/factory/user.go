package factory

import (
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
)

type UserEntity entity.User

type UserFactory interface {
	Create(entity UserEntity) entity.User
}

func (user *UserEntity) Create(u entity.User) *entity.User {
	db := infrastructure.DbManager()
	instance := model.Users{
		ID:          u.ID,
		Email:       u.Email,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Password:    u.Password,
		Institution: u.Institution,
		Role:        u.Role,
		Active:      u.Active,
	}
	db.Create(&instance)
	return &entity.User{
		ID:          instance.ID,
		Email:       instance.Email,
		FirstName:   instance.FirstName,
		LastName:    instance.LastName,
		Password:    instance.Password,
		Institution: instance.Institution,
		Role:        instance.Role,
		Active:      instance.Active,
	}
}
