package interfaces

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/inventory-manager/internal/application"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	type LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	payload := LoginPayload{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&payload)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user, err := application.GetUserByEmail(payload.Email)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if user.Password != payload.Password {
		return echo.ErrUnauthorized
	}
	id := user.ID.String()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = id
	claims["admin"] = user.Role == "admin"
	claims["institution"] = user.Institution
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Users(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	users, _ := application.GetUsers(*user)
	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	instance, err := application.GetUserById(userID, *user)
	if user.Institution != instance.Institution {
		forbiddenError := make(map[string]string)
		forbiddenError["error"] = "action-not-allowed"
		forbiddenError["details"] = "Cannot retrieve user from other institution"
		return echo.NewHTTPError(http.StatusForbidden)
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, instance)
}

func CreateUser(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	if user.Role != "admin" {
		forbiddenError := make(map[string]string)
		forbiddenError["error"] = "action-not-allowed"
		forbiddenError["details"] = "Only admin users can create users"
		return echo.NewHTTPError(http.StatusForbidden, forbiddenError)
	}
	newUser := application.User{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	institutionId, err := uuid.Parse(newUser.Institution)
	if err != nil {
		log.Fatalf("Failed to parse Institution UUID %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	instance, err := application.CreateUser(newUser, institutionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusCreated, instance)
}
