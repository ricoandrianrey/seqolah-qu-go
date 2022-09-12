package controllers

import (
	"context"
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
