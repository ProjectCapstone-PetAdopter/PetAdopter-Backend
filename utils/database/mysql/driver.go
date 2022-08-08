package mysql

import (
	"fmt"
	"log"
	"petadopter/config"
	userdata "petadopter/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Appconfig) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username,
		cfg.Password, cfg.Address, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Println("Cannot connect to db")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(userdata.User{})
}
