package main

import (
	"fmt"

	"petadopter/config"
	"petadopter/factory"
	middlewares "petadopter/features/middlewares"
	"petadopter/utils/database/mysql"
	auth "petadopter/utils/google"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()
	authconn := auth.InitOauth()
	factory.InitFactory(e, db, authconn)

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(address))
}
