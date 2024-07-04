package controller

import (
	"go-server/data/request"
	"go-server/helper"
	"go-server/service"
	"net/http"

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
				Data:   studentRequest,
			}
			helper.WriteResponse(writer, WebResponse, http.StatusCreated)
		} else {
			webRepo := helper.WebResponse{
				Status: err.Error(),
			}
			helper.WriteResponse(writer, webRepo, http.StatusConflict)
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
		email := params.ByName("email")
		if email == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Email is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		//studentUpdate.Id = id
		result, errx := controller.StudentService.Update(requests.Context(), studentUpdate, email)
		if result {
			webRepo := helper.WebResponse{

				Status: "created",
				Data:   studentUpdate,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{

				//Status: "Error during update",
				Status: errx.Error(),
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

func (controller *StudentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	valid := service.CheckToken(requests)
	if valid {
		email := params.ByName("email")
		if email == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Email is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		result, errx := controller.StudentService.Delete(requests.Context(), email)
		if result {
			webRepo := helper.WebResponse{
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{
				//Status: "Error during delete",
				Status: errx.Error(),
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
		email := params.ByName("email")
		if email == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Email is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		result, errx := controller.StudentService.FindById(requests.Context(), email)
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
