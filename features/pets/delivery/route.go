package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/pets/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePets(e *echo.Echo, pet domain.PetsHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/pets", pet.InsertPets(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/pets/:id", pet.UpdatePets(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/pets/:id", pet.DeletePets(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/pets", pet.GetAllPets())
	e.GET("/pets/:id", pet.GetPetsID())
	e.GET("/mypets", pet.GetmyPets(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
