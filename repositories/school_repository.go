package repositories

import (
	"seqolah-qu/entities/models"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	FindById(id int)
}

type SchoolRepositoryImpl struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepositoryImpl {
	return &SchoolRepositoryImpl{db}
}

func (r *SchoolRepositoryImpl) FindById(id int) (models.School, error) {
	var school models.School

	err := r.db.Find(&school, id).Error

	return school, err
}
