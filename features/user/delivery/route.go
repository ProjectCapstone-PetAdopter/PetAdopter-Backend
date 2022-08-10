package delivery

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/user/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, usr domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/oauth/signup", usr.SignUpGoogle())
	e.GET("/callback", usr.CallbackGoogle())
	e.POST("/login", usr.Login())
	e.DELETE("/users", usr.DeleteUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/users", usr.Register())
	e.PUT("/users", usr.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/users", usr.GetProfile(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
