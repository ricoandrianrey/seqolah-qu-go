package controllers

import (
	"encoding/json"
	"net/http"
	"seqolah-qu/services"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type SchoolController struct {
	r             *chi.Mux
	schoolService *services.SchoolServiceImpl
}

func NewSchoolController(r *chi.Mux, schoolService *services.SchoolServiceImpl) *SchoolController {
	return &SchoolController{r, schoolService}
}

func (c *SchoolController) FindSchoolById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(paramId)

		school := c.schoolService.FindSchoolById(id)
		b, _ := json.Marshal(school)
		w.Write(b)
	}
}
