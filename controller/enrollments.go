package controller

import (
	"go-server/Model"
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service/enrollment"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type EnrollmentController struct {
	EnrollmentService enrollment.EnrollmentService
}

func NewEnrollmentController(enrollmentService enrollment.EnrollmentService) *EnrollmentController {
	return &EnrollmentController{EnrollmentService: enrollmentService}
}
func (controller *EnrollmentController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	enrollmentRequest := request.EnrollmentRequest{}
	helper.ReadRequestBody(requests, &enrollmentRequest)
	enrollment := &Model.Enrollment{
		Student: Model.Student{Id: enrollmentRequest.StudentId},
		Corsue:  Model.Course{CourseId: enrollmentRequest.CourseId}, // Fixed field name
	}
	controller.EnrollmentService.Create(requests.Context(), enrollment)
	WebResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, WebResponse)
}
func (controller *EnrollmentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	enrollment, errx := controller.EnrollmentService.FindAll(requests.Context())
	if errx != nil {
		webRepo := response.WebResponse{
			Code:   500,
			Status: "error",
			Data:   errx.Error(),
		}
		helper.WriteResponse(writer, webRepo)
		return
	}

	var enrollmentResponses []response.EnrollmentResponse
	for i := 0; i < len(enrollment); i++ {
		enrollmentResponse := response.EnrollmentResponse{
			Name:       enrollment[i].Student.Name,
			Surname:    enrollment[i].Student.Surname,
			Email:      enrollment[i].Student.Email,
			CourseName: enrollment[i].Corsue.CourseName,
			StartYear:  enrollment[i].Corsue.StartYear,
			EndYear:    enrollment[i].Corsue.EndYear,
		}
		enrollmentResponses = append(enrollmentResponses, enrollmentResponse)
	}

	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   enrollmentResponses,
	}

	helper.WriteResponse(writer, webRepo)
}

func (controller *EnrollmentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentid := params.ByName("studentid")
	courseid := params.ByName("courseid")
	s_id, err := strconv.Atoi(studentid)
	helper.PanicIfError(err)
	c_id, err := strconv.Atoi(courseid)
	enrollment := &Model.Enrollment{
		Student: Model.Student{Id: s_id},
		Corsue:  Model.Course{CourseId: c_id},
	}
	helper.PanicIfError(err)
	controller.EnrollmentService.Delete(requests.Context(), enrollment)
	webRepo := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, webRepo)
}
