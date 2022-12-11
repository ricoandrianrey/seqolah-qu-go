package repositories

import (
	"seqolah-qu/entities/models"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	FindById(id int)
	Create(data models.School)
	Update(data models.School, id int)
	Delete(id int)
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

func (r *SchoolRepositoryImpl) Create(data models.School) (models.School, error) {
	err := r.db.Create(&data).Error

	return data, err
}

func (r *SchoolRepositoryImpl) Update(data models.School, id int) (int, error) {
	resp := r.db.Where("ID = ?", uint(id)).Updates(&data)

	return int(resp.RowsAffected), resp.Error
}

func (r *SchoolRepositoryImpl) Delete(id int) (int, error) {
	var school models.School
	resp := r.db.Delete(&school, id)

	return int(resp.RowsAffected), resp.Error
}
