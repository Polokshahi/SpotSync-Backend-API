package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	// Load .env
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found")
	}


	dsn := os.Getenv("DATABASE_URL")


	if dsn == "" {

		log.Fatal(
			"DATABASE_URL is missing in .env file",
		)
	}



	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)


	if err != nil {

		log.Fatal(
			"Failed to connect database: ",
			err,
		)
	}



	log.Println("Neon PostgreSQL connected successfully")


	return db
}