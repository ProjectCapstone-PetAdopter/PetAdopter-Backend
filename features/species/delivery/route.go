package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/species/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteSpecies(e *echo.Echo, species domain.SpeciesHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/species", species.AddSpecies(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/species/:id", species.UpdateDataSpecies(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/species/:id", species.DeleteDataSpecies(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/species", species.GetSpecies())
}
