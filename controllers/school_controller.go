package controllers

import (
	"context"
	"seqolah-qu/entities/dto"
	"seqolah-qu/services"
	"seqolah-qu/types"
	"strconv"
)

type SchoolController struct {
	schoolService *services.SchoolServiceImpl
}

func NewSchoolController(schoolService *services.SchoolServiceImpl) *SchoolController {
	return &SchoolController{schoolService}
}

func (c *SchoolController) FindSchoolById(data types.Request, ctx context.Context) (types.Response, error) {
	id, _ := strconv.Atoi(data.Params("id"))

	resp := c.schoolService.FindSchoolById(id)

	return types.Response{
		Message: "Success get data",
		Data:    resp,
	}, nil
}

func (c *SchoolController) CreateSchool(data types.Request, ctx context.Context) (types.Response, error) {
	var body dto.School
	data.Body(&body)

	resp := c.schoolService.CreateSchool(body)

	return types.Response{
		Message: "Success create data",
		Data:    resp,
	}, nil
}

func (c *SchoolController) UpdateSchool(data types.Request, ctx context.Context) (types.Response, error) {
	id, _ := strconv.Atoi(data.Params("id"))
	var body dto.School
	data.Body(&body)

	resp := c.schoolService.UpdateSchool(body, id)

	return types.Response{
		Message: "Success update data",
		Data:    resp,
	}, nil
}

func (c *SchoolController) DeleteSchool(data types.Request, ctx context.Context) (types.Response, error) {
	id, _ := strconv.Atoi(data.Params("id"))

	resp := c.schoolService.DeleteSchool(id)

	return types.Response{
		Message: "Success delete data",
		Data:    resp,
	}, nil
}
