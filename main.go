package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"tugas2/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err, "cant load .env file")
	}

	dbConfig := database.PostgresConfig{
		Host:     os.Getenv("POSTGRES_ADDR"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db := database.StartDB(dbConfig)

}
