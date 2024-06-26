package controller

import (
	"go-server/Model"
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service"
	"go-server/service/enrollment"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LectureStudentController struct {
	StudentLectureService enrollment.StudentLectureService
}

func NewStudentLectureController(enrollmentService enrollment.StudentLectureService) *LectureStudentController {
	return &LectureStudentController{StudentLectureService: enrollmentService}
}
func (controller *LectureStudentController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		enrollmentRequest := request.EnrollmentRequest{}
		helper.ReadRequestBody(requests, &enrollmentRequest)
		enrollment := &Model.StudentLectures{
			Student: Model.Student{Id: enrollmentRequest.StudentId},
			Lecture: Model.Lectures{LectureId: enrollmentRequest.LectureId},
		}
		result, errx := controller.StudentLectureService.Create(requests.Context(), enrollment)
		if result {
			WebResponse := helper.WebResponse{
				Status: "success",
				//Data:   result,
			}
			helper.WriteResponse(writer, WebResponse, http.StatusCreated)
		} else {
			webRepo := helper.WebResponse{

				Status: "Error during insert",
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
func (controller *LectureStudentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		enrollment, errx := controller.StudentLectureService.FindAll(requests.Context())
		if errx != nil {
			webRepo := helper.WebResponse{
				Status: "error",
				Data:   errx.Error(),
			}
			helper.WriteResponse(writer, webRepo, http.StatusNotFound)
			return
		}

		var enrollmentResponses []response.EnrollmentResponse
		for i := 0; i < len(enrollment); i++ {
			enrollmentResponse := response.EnrollmentResponse{
				Name:             enrollment[i].Student.Name,
				Surname:          enrollment[i].Student.Surname,
				Email:            enrollment[i].Student.Email,
				LectureName:      enrollment[i].Lecture.LectureName,
				StartYear:        enrollment[i].Lecture.StartYear,
				EndYear:          enrollment[i].Lecture.EndYear,
				ProfessorSurname: enrollment[i].Professor.Surname,
				DepartmentName:   enrollment[i].Department.Name,
			}
			enrollmentResponses = append(enrollmentResponses, enrollmentResponse)
		}

		webRepo := helper.WebResponse{
			Status: "ok",
			Data:   enrollmentResponses,
		}

		helper.WriteResponse(writer, webRepo, http.StatusOK)
	} else {
		webRepo := helper.WebResponse{
			Status: "Error token not valid",
		}

		helper.WriteResponse(writer, webRepo, http.StatusUnauthorized)

	}
}

func (controller *LectureStudentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	valid := service.CheckToken(requests)
	if valid {
		studentid := params.ByName("studentid")
		lectureid := params.ByName("lectureid")
		s_id, err := strconv.Atoi(studentid)
		if err != nil {
			webRepo := helper.WebResponse{
				Status: "Error during parsing studentid",
				Data:   err,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
			return
		}
		c_id, err := strconv.Atoi(lectureid)
		if err != nil {
			webRepo := helper.WebResponse{
				Status: "Error during parsing lectureid",
				Data:   err,
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}

		enrollment := &Model.StudentLectures{
			Student: Model.Student{Id: s_id},
			Lecture: Model.Lectures{LectureId: c_id},
		}
		helper.PanicIfError(err)
		result, err := controller.StudentLectureService.Delete(requests.Context(), enrollment)
		if result {
			webRepo := helper.WebResponse{
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{
				Status: "Error during delete",
				Data:   err,
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
