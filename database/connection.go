package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn   *gorm.DB
	user     string
	password string
	host     string
	db       string
	port     string
)

func LoadVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	db = os.Getenv("DB_DB")
	port = os.Getenv("DB_PORT")

}

// Connect creates a connection to database
func Connect() (err error) {
	LoadVar()

	port, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, port)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := DBConn.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
