package services

import (
	"seqolah-qu/entities/dto"
	"seqolah-qu/entities/models"
	"seqolah-qu/repositories"
)

type SchoolService interface {
	FindSchoolById(id int)
	CreateSchool(data dto.School)
	UpdateSchool(data dto.School, id int)
	DeleteSchool(id int)
}

type SchoolServiceImpl struct {
	schoolRepo *repositories.SchoolRepositoryImpl
}

func NewSchoolService(schoolRepo *repositories.SchoolRepositoryImpl) *SchoolServiceImpl {
	return &SchoolServiceImpl{schoolRepo}
}

func (s *SchoolServiceImpl) FindSchoolById(id int) models.School {
	school, _ := s.schoolRepo.FindById(id)
	return school
}

func (s *SchoolServiceImpl) CreateSchool(data dto.School) models.School {
	newSchool := models.School{Address: data.Address, AddressDetail: data.AddressDetail, Logo: data.Logo, Email: data.Email, Accreditation: data.Accreditation, Phone: data.Phone, Curriculum: data.Curriculum}
	school, _ := s.schoolRepo.Create(newSchool)
	return school
}

func (s *SchoolServiceImpl) UpdateSchool(data dto.School, id int) int {
	newSchool := models.School{Address: data.Address, AddressDetail: data.AddressDetail, Logo: data.Logo, Email: data.Email, Accreditation: data.Accreditation, Phone: data.Phone, Curriculum: data.Curriculum}
	affected, _ := s.schoolRepo.Update(newSchool, id)
	return affected
}

func (s *SchoolServiceImpl) DeleteSchool(id int) int {
	affected, _ := s.schoolRepo.Delete(id)
	return affected
}
