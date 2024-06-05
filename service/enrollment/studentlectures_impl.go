package enrollment

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
)

type StudentLectureServiceImpl struct {
	enrollmentRepo repository.StudentLectureRepo
}

// Create implements EnrollmentService.
func (s *StudentLectureServiceImpl) Create(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error) {
	result, errx := s.enrollmentRepo.Save(ctx, *studentlecture)
	return result, errx
}

// Delete implements EnrollmentService.
func (s *StudentLectureServiceImpl) Delete(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error) {
	result, errx := s.enrollmentRepo.Delete(ctx, studentlecture.Student.Id, studentlecture.Lecture.LectureId)
	return result, errx
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
		}
		response = append(response, enrollment)
	}
	return response, nil
}

func NewStudentLectureServiceImpl(studentLectureRepo repository.StudentLectureRepo) StudentLectureService {
	return &StudentLectureServiceImpl{enrollmentRepo: studentLectureRepo}
}
