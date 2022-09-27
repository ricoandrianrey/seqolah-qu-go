package repositories

import (
	"seqolah-qu/entities/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindOne()
	Create()
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) FindOne(conditions models.User) (models.User, error) {
	var user models.User

	err := r.db.Where(conditions).First(&user).Error

	return user, err
}

func (r *UserRepositoryImpl) Create(data models.User) (models.User, error) {
	user := data

	err := r.db.Create(&user).Error

	return user, err
}
