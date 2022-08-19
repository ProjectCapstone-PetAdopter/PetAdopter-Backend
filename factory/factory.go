package factory

import (
	"petadopter/utils/google"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
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

	sd "petadopter/features/species/data"
	speciesDelivery "petadopter/features/species/delivery"
	su "petadopter/features/species/usecase"

	md "petadopter/features/meeting/data"
	meetingDelivery "petadopter/features/meeting/delivery"
	mu "petadopter/features/meeting/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB, oauth2 *oauth2.Config, client *google.ClientUploader) {
	valid := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, valid, client)
	userHandler := udeli.New(userCase, oauth2)
	udeli.RouteUser(e, userHandler)

	petsData := pd.New(db)
	petsCase := pu.New(petsData, valid, client)
	petsHandler := petsDelivery.New(petsCase)
	petsDelivery.RoutePets(e, petsHandler)

	adoptData := ad.New(db)
	adoptCase := au.New(adoptData, valid)
	adoptHandler := adoptionDelivery.New(adoptCase)
	adoptionDelivery.RouteAdopt(e, adoptHandler)

	speciesData := sd.New(db)
	speciesCase := su.New(speciesData, valid)
	speciesHandler := speciesDelivery.New(speciesCase)
	speciesDelivery.RouteSpecies(e, speciesHandler)

	meetingData := md.New(db)
	meetingCase := mu.New(meetingData, valid)
	meetingHandler := meetingDelivery.New(meetingCase, oauth2)
	meetingDelivery.RouteMeeting(e, meetingHandler)

}
