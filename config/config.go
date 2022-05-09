package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
	SSL      string
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ConnectDB() *gorm.DB {
	var cred Config

	cred.Username = os.Getenv("DB_USER")
	cred.Password = os.Getenv("DB_PASS")
	cred.Host = os.Getenv("DB_HOST")
	cred.DBName = os.Getenv("DB_NAME")
	cred.Port = os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cred.Username, cred.Password, cred.Host, cred.Port, cred.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	FailOnError(err)

	return db

}
