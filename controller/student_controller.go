package controller

import (
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StudentController struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) *StudentController {
	return &StudentController{StudentService: studentService}
}

func (controller *StudentController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		studentRequest := request.StudentCreateRequest{}
		helper.ReadRequestBody(requests, &studentRequest)
		controller.StudentService.Create(requests.Context(), studentRequest)
		WebResponse := response.WebResponse{
			Code:   200,
			Status: "ok",
			Data:   nil,
		}
		helper.WriteResponse(writer, WebResponse)
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}
		helper.WriteResponse(writer, webRepo)
	}
}

func (controller *StudentController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentUpdate := request.StudentUpdateRequest{}
	helper.ReadRequestBody(requests, &studentUpdate)
	studentid := params.ByName("idstudente")
	id, err := strconv.Atoi(studentid)
	helper.PanicIfError(err)
	studentUpdate.Id = id
	controller.StudentService.Update(requests.Context(), studentUpdate)
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, webRepo)
}

func (controller *StudentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	valid := service.CheckToken(requests)
	if valid {
		studentid := params.ByName("idstudente")
		id, err := strconv.Atoi(studentid)
		helper.PanicIfError(err)
		controller.StudentService.Delete(requests.Context(), id)
		webRepo := response.WebResponse{
			Code:   200,
			Status: "ok",
			Data:   nil,
		}
		helper.WriteResponse(writer, webRepo)
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo)
	}
}

func (controller *StudentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	valid := service.CheckToken(requests)
	if valid {
		result := controller.StudentService.FindAll(requests.Context())
		webRepo := response.WebResponse{
			Code:   200,
			Status: "ok",
			Data:   result,
		}
		helper.WriteResponse(writer, webRepo)
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo)
	}
}

func (controller *StudentController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		studentid := params.ByName("idstudente")
		id, err := strconv.Atoi(studentid)
		helper.PanicIfError(err)
		result := controller.StudentService.FindById(requests.Context(), id)
		webRepo := response.WebResponse{
			Code:   200,
			Status: "ok",
			Data:   result,
		}
		helper.WriteResponse(writer, webRepo)
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo)
	}
}
