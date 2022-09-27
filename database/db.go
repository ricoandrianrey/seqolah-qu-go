package database

import (
	"os"
	"seqolah-qu/entities/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbUrl := os.Getenv("DB_CONNECTION_URL")
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.School{},
		&models.User{},
		&models.UserSchool{},
		&models.UserParent{},
		&models.UserStudent{},
	)

	return db
}
