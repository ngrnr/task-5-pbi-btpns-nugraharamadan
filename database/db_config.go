package database

import (
	"fmt"
	"log"

	"github.com/ngrnr/task-5-pbi-btpns-nugraharamadan/helpers"
	"github.com/ngrnr/task-5-pbi-btpns-nugraharamadan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	var path string
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "testing" {
		path = "../.env"
	}
	if stage != "testing" {
		path = ".env"
	}

	//
	helpers.LoadEnv(path)

	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "123"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 5433),
		helpers.GetAsString("DB_NAME", "postgres"),
	)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MigrateDB() {
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "development" ||
		stage == "production" {
		db.Debug().AutoMigrate(models.User{}, models.Photo{})
	}
}

func GetDB() *gorm.DB {
	return db
}
