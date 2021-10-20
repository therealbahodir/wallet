package database 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
)

func DBConnection () (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PG_HOST := os.Getenv("DB_HOST")
	PG_USER := os.Getenv("DB_USER")
	PG_PASSWORD := os.Getenv("DB_PASSWORD")
	PG_PORT := os.Getenv("DB_PORT")
	PG_DBNAME := os.Getenv("DB_DBNAME")
	PG_SSLMODE := os.Getenv("DB_SSLMODE")
	PG_TIMEZONE := os.Getenv("DB_TIME_ZONE")

	PGConnection := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		PG_HOST,
		PG_USER,
		PG_PASSWORD,
		PG_PORT,
		PG_DBNAME,
		PG_SSLMODE,
		PG_TIMEZONE,
	)

	return gorm.Open(postgres.Open(PGConnection), &gorm.Config{})
}