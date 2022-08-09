package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/adoption/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteAdopt(e *echo.Echo, adopt domain.AdoptionHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/adoptions", adopt.InsertAdoption(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/adoptions/:id", adopt.UpdateAdoption(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/adoptions/:id", adopt.DeleteAdoption(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/adoptions", adopt.GetAllAdoption())
	e.GET("/adoptions/:id", adopt.GetAdoptionID())
	e.GET("/myadoptions", adopt.GetMYAdopt())
}
