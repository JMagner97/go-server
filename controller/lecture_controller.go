package controller

import (
	"go-server/data/request"
	"go-server/helper"
	"go-server/service"
	lecture "go-server/service/lectures"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LectureController struct {
	lectureService lecture.LectureService
}

func NewLecturesController(lectureService lecture.LectureService) *LectureController {
	return &LectureController{
		lectureService: lectureService,
	}
}

func (controller *LectureController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		courseRequest := request.LectureCreateRequest{}
		helper.ReadRequestBody(requests, &courseRequest)
		result, errx := controller.lectureService.Create(requests.Context(), courseRequest)
		if result {
			WebResponse := helper.WebResponse{
				Status: "ok",
				Data:   courseRequest,
			}
			helper.WriteResponse(writer, WebResponse, http.StatusCreated)
		} else {
			webRepo := helper.WebResponse{
				//Status: "Error during create",
				Status: errx.Error(),
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

func (controller *LectureController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		lectureUpdate := request.LectureUpdateRequest{}
		helper.ReadRequestBody(requests, &lectureUpdate)
		name := params.ByName("name")
		if name == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Name is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		result, errx := controller.lectureService.Update(requests.Context(), lectureUpdate, name)
		if result {
			webRepo := helper.WebResponse{

				Status: "ok",
				Data:   nil,
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

func (controller *LectureController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		name := params.ByName("name")
		if name == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Name is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		result, errx := controller.lectureService.Delete(requests.Context(), name)
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

func (controller *LectureController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		result := controller.lectureService.FindAll(requests.Context())
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

func (controller *LectureController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		name := params.ByName("name")
		if name == "" {
			webRepo := helper.WebResponse{
				Status: "Error",
				Data:   "Name is required",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}
		result, errx := controller.lectureService.FindById(requests.Context(), name)
		if errx != nil {
			webRepo := helper.WebResponse{
				Status: "Lectures not found",
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
