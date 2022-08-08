package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePets(e *echo.Echo, bc domain.PetsHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/pets", bc.InsertPets(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/pets/:id", bc.UpdatePets(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/pets/:id", bc.DeletePets(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/pets", bc.GetAllPets())
	e.GET("/pets/:id", bc.GetPetsID())
}
