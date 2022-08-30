package services

import (
	"seqolah-qu/entities"
	"seqolah-qu/repositories"
)

type SchoolService interface {
	FindSchoolById(id int)
}

type SchoolServiceImpl struct {
	schoolRepo *repositories.SchoolRepositoryImpl
}

func NewSchoolService(schoolRepo *repositories.SchoolRepositoryImpl) *SchoolServiceImpl {
	return &SchoolServiceImpl{schoolRepo}
}

func (s *SchoolServiceImpl) FindSchoolById(id int) *entities.School {
	school, _ := s.schoolRepo.FindById(id)
	return &school
}
