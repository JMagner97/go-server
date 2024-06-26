package controller

import (
	"go-server/data/request"
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
			WebResponse := helper.WebResponse{
				Status: "Created",
				Data:   nil,
			}
			helper.WriteResponse(writer, WebResponse, http.StatusCreated)
		} else {
			webRepo := helper.WebResponse{
				Status: "Error during insert",
				Data:   err,
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
		}
	} else {
		webRepo := helper.WebResponse{
			Status: "Error token not valid",
		}
		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)
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
			webRepo := helper.WebResponse{

				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{

				Status: "Error during update",
				Data:   errx,
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
		}
	} else {

		webRepo := helper.WebResponse{

			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)
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
			webRepo := helper.WebResponse{
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{
				Status: "Error during delete",
				Data:   errx,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNotFound)
		}
	} else {
		webRepo := helper.WebResponse{
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)
	}
}

func (controller *StudentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	valid := service.CheckToken(requests)
	if valid {
		result := controller.StudentService.FindAll(requests.Context())
		webRepo := helper.WebResponse{
			Status: "ok",
			Data:   result,
		}
		helper.WriteResponse(writer, webRepo, http.StatusOK)
	} else {
		webRepo := helper.WebResponse{

			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)
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
			webRepo := helper.WebResponse{
				Status: "Student not found",
			}
			helper.WriteResponse(writer, webRepo, http.StatusNotFound)
		} else {
			webRepo := helper.WebResponse{

				Status: "ok",
				Data:   result,
			}
			helper.WriteResponse(writer, webRepo, http.StatusOK)
		}
	} else {
		webRepo := helper.WebResponse{
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)
	}
}
