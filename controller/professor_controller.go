package controller

import (
	re "go-server/data/request"
	"go-server/helper"
	professorservice "go-server/service/professor_service"
	"net/http"
)

type ProfessorController struct {
	ProfessorService professorservice.ProfessorService
}

func NewProfessorsController(ProfessorService professorservice.ProfessorService) *ProfessorController {
	return &ProfessorController{
		ProfessorService: ProfessorService,
	}
}

func (controller *ProfessorController) Create(writer http.ResponseWriter, request *http.Request) {
	professorRequest := re.ProfessorRequest{}
	helper.ReadRequestBody(request, &professorRequest)
	res, err := controller.ProfessorService.Create(request.Context(), professorRequest)
	if res {
		WebResponse := helper.WebResponse{
			Status: "Created",
			Data:   professorRequest,
		}
		helper.WriteResponse(writer, WebResponse, http.StatusCreated)
	} else {
		webRepo := helper.WebResponse{
			Status: err.Error(),
		}
		helper.WriteResponse(writer, webRepo, http.StatusConflict)
	}
}
