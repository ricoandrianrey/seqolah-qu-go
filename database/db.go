package database

import (
	"os"
	"seqolah-qu/entities"

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
		&entities.School{},
		&entities.User{},
		&entities.UserSchool{},
		&entities.UserParent{},
		&entities.UserStudent{},
	)

	return db
}
