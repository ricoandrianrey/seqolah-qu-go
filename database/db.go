package database

import (
	"seqolah-qu/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=0.0.0.0 user=postgres password=postgres dbname=seqolah-qu port=32771 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
