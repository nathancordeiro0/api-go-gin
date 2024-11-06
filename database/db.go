package database

import (
	"log"
	"os"

	"github.com/nathancordeiro0/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_CONN_STR")))

	if err != nil {
		log.Panic("Error:", err)
	}

	DB.AutoMigrate(&models.Student{})
}
