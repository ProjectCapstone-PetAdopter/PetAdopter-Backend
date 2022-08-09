package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "petadopter/features/user/data"
	udeli "petadopter/features/user/delivery"
	uc "petadopter/features/user/usecase"

	pd "petadopter/features/pets/data"
	petsDelivery "petadopter/features/pets/delivery"
	pu "petadopter/features/pets/usecase"

	ad "petadopter/features/adoption/data"
	adoptionDelivery "petadopter/features/adoption/delivery"
	au "petadopter/features/adoption/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	valid := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, valid)
	userHandler := udeli.New(userCase)
	udeli.RouteUser(e, userHandler)

	petsData := pd.New(db)
	petsCase := pu.New(petsData, valid)
	petsHandler := petsDelivery.New(petsCase)
	petsDelivery.RoutePets(e, petsHandler)

	adoptData := ad.New(db)
	adoptCase := au.New(adoptData, valid)
	adoptHandler := adoptionDelivery.New(adoptCase)
	adoptionDelivery.RouteAdopt(e, adoptHandler)

}
