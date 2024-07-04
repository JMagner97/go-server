package enrollment

import (
	"context"
	"errors"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/response"
)

type StudentLectureServiceImpl struct {
	enrollmentRepo repository.StudentLectureRepo
}

// FindById implements StudentLectureService.
func (s *StudentLectureServiceImpl) FindById(ctx context.Context, studentlecture *Model.StudentLectures) (response.EnrollmentResponse, error) {
	result, err := s.enrollmentRepo.FindById(ctx, studentlecture.Student.Email, studentlecture.Lecture.LectureName)
	enrollment := response.EnrollmentResponse{
		Name:             result.Student.Name,
		Surname:          result.Student.Surname,
		Email:            result.Student.Email,
		LectureName:      result.Lecture.LectureName,
		StartYear:        result.Lecture.StartYear,
		EndYear:          result.Lecture.EndYear,
		ProfessorSurname: result.Professor.Surname,
		DepartmentName:   result.Department.Name,
	}
	return enrollment, err
}

// Create implements EnrollmentService.
func (s *StudentLectureServiceImpl) Create(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error) {
	exists, err := s.enrollmentRepo.EnrollmentExist(ctx, studentlecture.Student.Email, studentlecture.Lecture.LectureName)
	if err != nil {
		return false, err
	}
	if !exists {
		result, errx := s.enrollmentRepo.Save(ctx, studentlecture.Student.Email, studentlecture.Lecture.LectureName)
		return result, errx
	} else {
		return false, errors.New("enrollment already exists")
	}
}

// Delete implements EnrollmentService.
func (s *StudentLectureServiceImpl) Delete(ctx context.Context, email string, name string) (bool, error) {
	exist, _ := s.enrollmentRepo.EnrollmentExist(ctx, email, name)
	if exist {
		result, errx := s.enrollmentRepo.Delete(ctx, email, name)
		return result, errx
	} else {
		return false, errors.New("enrollment not found")
	}
}

// FindAll implements EnrollmentService.
func (s *StudentLectureServiceImpl) FindAll(ctx context.Context) ([]*Model.StudentLectures, error) {
	list := s.enrollmentRepo.FindAll(ctx)
	var response []*Model.StudentLectures
	for _, value := range list {
		enrollment := &Model.StudentLectures{
			Student: Model.Student{
				Id:        value.Student.Id,
				Name:      value.Student.Name,
				Surname:   value.Student.Surname,
				Email:     value.Student.Email,
				Address:   value.Student.Address,
				Birthdate: value.Student.Birthdate,
			},
			Lecture: Model.Lectures{
				LectureId:   value.Lecture.LectureId,
				LectureName: value.Lecture.LectureName,
				StartYear:   value.Lecture.StartYear,
				EndYear:     value.Lecture.EndYear,
				Description: value.Lecture.Description,
				ProfessorId: value.Lecture.ProfessorId,
			},
			Professor: Model.Professor{
				ProfessorId: value.Professor.ProfessorId,
				Surname:     value.Professor.Surname,
			},
			Department: Model.Department{
				//DepartmentId: value.Department.DepartmentId,
				Name: value.Department.Name,
			},
		}
		response = append(response, enrollment)
	}
	return response, nil
}

func NewStudentLectureServiceImpl(studentLectureRepo repository.StudentLectureRepo) StudentLectureService {
	return &StudentLectureServiceImpl{enrollmentRepo: studentLectureRepo}
}
