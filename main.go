package main

import (
	"fmt"

	"petadopter/config"
	"petadopter/factory"
	middlewares "petadopter/features/middlewares"
	"petadopter/utils/database/mysql"
	"petadopter/utils/google"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()
	authconn := google.InitOauth()
	storageconn := google.InitStorage("pet-adopter-358806-9e20643cb88d.json", "be10-petdopter", "pet-adopter-358806")
	factory.InitFactory(e, db, authconn, storageconn)

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(address))
}
