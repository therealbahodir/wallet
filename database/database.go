package database 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
	"time"
	"errors"
)

type User struct {
	UserId string
	Digest string
	CreatedAt time.Time
	IsIdentified bool
	Balance float64
}

type Replenishment struct {
	UserId string
	Amount float64 
	ReceivedAt time.Time
}


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


func IsExisting (id, digest string) error {
	
	db, err:= DBConnection()
	if err != nil {
		log.Print(err)
		return err
	}

	var user User
	db.Where("user_id = ? AND digest = ?", id, digest).Find(&user)

	if user.UserId == id && user.Digest == digest {
		return nil
	} 

	err = errors.New("Incorrect user_id or digest")
	return err
}

func TopUpBalance (id, digest string, amount float64) error {

	db, err:= DBConnection()
	if err != nil {
		log.Print(err)
		return err
	}

	var user User
	db.Table("users").Where("user_id = ? AND digest = ?", id, digest).Find(&user)
	
	err = nil
	if user.UserId == id && user.Digest == digest {

		if user.IsIdentified == true && user.Balance <= 100000 - amount {
			user.Balance += amount
			db.Where("user_id = ? AND digest = ?", id, digest).Save(&user)
			EnterReplenishment(id, amount)
		} else if user.IsIdentified == false && user.Balance <= 10000 - amount{
			user.Balance += amount
			db.Where("user_id = ? AND digest = ?", id, digest).Save(&user)
			EnterReplenishment(id, amount)
		} else if user.IsIdentified == false && user.Balance >= 10000 - amount{
			err = errors.New("You cannot have more than 10000 tjs")
		} else if user.IsIdentified == true && user.Balance >= 100000 - amount {
			err = errors.New("You cannot have more than 100000 tjs")
		}

	} else {
		err = errors.New("Incorrect user_id or digest")
	}

	return err
}


func EnterReplenishment (id string, amount float64) error {
	db, err:= DBConnection()
	if err != nil {
		log.Print(err)
		return err
	}

	var replenishment Replenishment

	replenishment.UserId = id
	replenishment.Amount = amount
	replenishment.ReceivedAt = time.Now()

	result := db.Create(&replenishment)

	if result.Error != nil {
		log.Print(err)
		return err
	}
	return nil
}

func ReplenishmentsInfo (id string) (count int, sum float64) {
	
	db, err := DBConnection()
	if err != nil {
		log.Print(err)
		return
	}
	var replenishments []Replenishment

	db.Table("replenishments").Where("extract(month from received_at) = extract(month from current_timestamp) AND user_id = ?", id).Find(&replenishments)


	for _, repl := range replenishments{
		count += 1 
		sum += repl.Amount
	}
	return count, sum
}