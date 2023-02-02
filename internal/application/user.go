package application

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/entity"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/factory"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/service"
	"gorm.io/gorm"
)

type User struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	Institution string `json:"institution"`
	Role        string `json:"role"`
	Active      bool   `json:"active"`
}

func GetUserFromContext(c echo.Context) (*entity.User, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uId := claims["id"].(string)
	instId := claims["institution"].(string)
	userId, err := uuid.Parse(uId)
	if err != nil {
		return nil, err
	}
	institutionId, err := uuid.Parse(instId)
	if err != nil {
		return nil, err
	}
	repository := service.UserEntity{}
	instance, err := repository.Get(userId, institutionId)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func GetUsers(user entity.User) (*[]entity.User, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.UserEntity{}
	entities, err := repository.GetAll(institution.ID)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func GetUserById(id uuid.UUID, user entity.User) (*entity.User, error) {
	institution, err := GetInstitutionFromUser(user)
	if err != nil {
		return nil, err
	}
	repository := service.UserEntity{}
	entity, err := repository.Get(id, institution.ID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func GetUserByEmail(email string) (*entity.User, error) {
	repository := service.UserEntity{}
	entity, err := repository.GetByEmail(email)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return entity, nil
}

func CreateUser(user User, institutionID uuid.UUID) (*entity.User, error) {
	repository := service.UserEntity{}
	factory := factory.UserEntity{}
	exists, err := repository.GetByEmail(user.Email)
	if err != nil {
		entity := entity.User{
			ID:          uuid.New(),
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Password:    user.Password,
			Institution: institutionID,
			Role:        user.Role,
			Active:      user.Active,
		}
		instance := factory.Create(entity)
		return instance, nil
	}
	errorMessage := "email already exists for ID:" + exists.ID.String()
	return nil, errors.New(errorMessage)
}
