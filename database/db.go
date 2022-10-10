package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tugas2/model"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func StartDB(config PostgresConfig) *gorm.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("cant open database ", err)
	}

	db.Debug().AutoMigrate(model.Order{}, model.Item{})

	return db
}
