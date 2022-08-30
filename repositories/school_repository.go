package repositories

import (
	"seqolah-qu/entities"

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

func (r *SchoolRepositoryImpl) FindById(id int) (entities.School, error) {
	var school entities.School

	err := r.db.Find(&school, id).Error

	return school, err
}
