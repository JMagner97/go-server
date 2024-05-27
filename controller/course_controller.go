package controller

import (
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service/course"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CourseController struct {
	courseService course.CourseService
}

func NewCourseController(courseService course.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

func (controller *CourseController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	courseRequest := request.CourseCreateRequest{}
	helper.ReadRequestBody(requests, &courseRequest)
	controller.courseService.Create(requests.Context(), courseRequest)
	WebResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, WebResponse)
}

func (controller *CourseController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	courseUpdate := request.CourseUpdateRequest{}
	helper.ReadRequestBody(requests, &courseUpdate)
	courseid := params.ByName("courseid")
	id, err := strconv.Atoi(courseid)
	helper.PanicIfError(err)
	courseUpdate.CourseId = id
	controller.courseService.Update(requests.Context(), courseUpdate)
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, webRepo)
}

func (controller *CourseController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	courseid := params.ByName("courseid")
	id, err := strconv.Atoi(courseid)
	helper.PanicIfError(err)
	controller.courseService.Delete(requests.Context(), id)
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, webRepo)
}

func (controller *CourseController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.courseService.FindAll(requests.Context())
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helper.WriteResponse(writer, webRepo)
}

func (controller *CourseController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	courseid := params.ByName("courseid")
	id, err := strconv.Atoi(courseid)
	helper.PanicIfError(err)
	result := controller.courseService.FindById(requests.Context(), id)
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helper.WriteResponse(writer, webRepo)
}
