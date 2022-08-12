package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/species/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteMeeting(e *echo.Echo, meetings domain.MeetingHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/meetings", meetings.InsertMeeting(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/meetings/:id", meetings.UpdateDataMeeting(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/meetings/:id", meetings.DeleteDataMeeting(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/meetings", meetings.GetAdopt(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
