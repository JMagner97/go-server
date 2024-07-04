package controller

import (
	"go-server/Model"
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
	"go-server/service"
	"go-server/service/enrollment"
	"net/http"

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
			Student: Model.Student{Email: enrollmentRequest.StudentEmail},
			Lecture: Model.Lectures{LectureName: enrollmentRequest.LectureName},
		}
		result, errx := controller.StudentLectureService.Create(requests.Context(), enrollment)
		if result {
			result1, err := controller.StudentLectureService.FindById(requests.Context(), enrollment)
			if err != nil {
				WebResponse := helper.WebResponse{
					Status: err.Error(),
					//Data:   enrollmentResponse,
				}
				helper.WriteResponse(writer, WebResponse, http.StatusNotFound)
			} else {
				enrollmentResponse := response.EnrollmentResponse{
					Name:             result1.Name,
					Surname:          result1.Surname,
					Email:            result1.Email,
					LectureName:      result1.LectureName,
					StartYear:        result1.StartYear,
					EndYear:          result1.EndYear,
					ProfessorSurname: result1.ProfessorSurname,
					DepartmentName:   result1.DepartmentName,
				}
				WebResponse := helper.WebResponse{
					Status: "created",
					Data:   enrollmentResponse,
				}
				helper.WriteResponse(writer, WebResponse, http.StatusCreated)
			}
		} else {
			webRepo := helper.WebResponse{

				//Status: "Error during insert",
				Status: errx.Error(),
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
		email := params.ByName("email")
		name := params.ByName("name")
		if email == "" || name == "" {
			webRepo := helper.WebResponse{
				//Status: "Error",
				Status: "Error in parameters",
			}
			helper.WriteResponse(writer, webRepo, http.StatusBadRequest)
			return
		}

		//	enrollment := &Model.StudentLectures{
		//		Student: Model.Student{Email: email},
		//		Lecture: Model.Lectures{LectureName: name},
		//	}

		result, err := controller.StudentLectureService.Delete(requests.Context(), email, name)
		if result {
			webRepo := helper.WebResponse{
				Status: "ok",
				Data:   nil,
			}
			helper.WriteResponse(writer, webRepo, http.StatusNoContent)
		} else {
			webRepo := helper.WebResponse{
				//Status: "Error during delete",
				Data: err.Error(),
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
