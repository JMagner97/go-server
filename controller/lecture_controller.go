package controller

import (
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service"
	lecture "go-server/service/lectures"
	"net/http"
	"strconv"

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
			WebResponse := response.WebResponse{
				Code:   200,
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, WebResponse)
		} else {
			webRepo := response.WebResponse{
				Code:   403,
				Status: "Error during create",
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

func (controller *LectureController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		lectureUpdate := request.LectureUpdateRequest{}
		helper.ReadRequestBody(requests, &lectureUpdate)
		lectureid := params.ByName("lectureid")
		id, err := strconv.Atoi(lectureid)
		helper.PanicIfError(err)
		lectureUpdate.LectureId = id
		result, errx := controller.lectureService.Update(requests.Context(), lectureUpdate)
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

func (controller *LectureController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		lectureid := params.ByName("lectureid")
		id, err := strconv.Atoi(lectureid)
		helper.PanicIfError(err)
		result, errx := controller.lectureService.Delete(requests.Context(), id)
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

func (controller *LectureController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		result := controller.lectureService.FindAll(requests.Context())
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

func (controller *LectureController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		lectureid := params.ByName("lectureid")
		id, err := strconv.Atoi(lectureid)
		helper.PanicIfError(err)
		result, errx := controller.lectureService.FindById(requests.Context(), id)
		if errx != nil {
			webRepo := response.WebResponse{
				Code:   404,
				Status: "Lectures not found",
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
