package factory

import (
	ud "petadopter/features/user/data"
	udeli "petadopter/features/user/delivery"
	uc "petadopter/features/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	valid := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, valid)
	userHandler := udeli.New(userCase)
	udeli.RouteUser(e, userHandler)

}
