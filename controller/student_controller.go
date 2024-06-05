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
		res, err := controller.StudentService.Create(requests.Context(), studentRequest)
		if res {
			WebResponse := response.WebResponse{
				Code:   200,
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, WebResponse)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Error during insert",
				Data:   err,
			}
			helper.WriteResponse(writer, webRepo)
		}
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}
		helper.WriteResponse(writer, webRepo)
	}
}

func (controller *StudentController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		studentUpdate := request.StudentUpdateRequest{}
		helper.ReadRequestBody(requests, &studentUpdate)
		studentid := params.ByName("idstudente")
		id, err := strconv.Atoi(studentid)
		helper.PanicIfError(err)
		studentUpdate.Id = id
		result, errx := controller.StudentService.Update(requests.Context(), studentUpdate)
		if result {
			webRepo := response.WebResponse{
				Code:   200,
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Error during update",
				Data:   errx,
			}
			helper.WriteResponse(writer, webRepo)
		}
	} else {

		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo)
	}
}

func (controller *StudentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	valid := service.CheckToken(requests)
	if valid {
		studentid := params.ByName("idstudente")
		id, err := strconv.Atoi(studentid)
		helper.PanicIfError(err)
		result, errx := controller.StudentService.Delete(requests.Context(), id)
		if result {
			webRepo := response.WebResponse{
				Code:   200,
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Error during delete",
				Data:   errx,
			}
			helper.WriteResponse(writer, webRepo)
		}
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
		result, errx := controller.StudentService.FindById(requests.Context(), id)
		if errx != nil {
			webRepo := response.WebResponse{
				Code:   404,
				Status: "Student not found",
			}
			helper.WriteResponse(writer, webRepo)
		} else {
			webRepo := response.WebResponse{
				Code:   200,
				Status: "ok",
				Data:   result,
			}
			helper.WriteResponse(writer, webRepo)
		}
	} else {
		webRepo := response.WebResponse{
			Code:   403,
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo)
	}
}
