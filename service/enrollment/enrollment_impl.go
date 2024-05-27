package enrollment

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
)

type EnrollmentServiceImpl struct {
	enrollmentRepo repository.EnrollmentRepo
}

// Create implements EnrollmentService.
func (s *EnrollmentServiceImpl) Create(ctx context.Context, enrollment *Model.Enrollment) {
	s.enrollmentRepo.Save(ctx, *enrollment)
}

// Delete implements EnrollmentService.
func (s *EnrollmentServiceImpl) Delete(ctx context.Context, enrollment *Model.Enrollment) {
	s.enrollmentRepo.Delete(ctx, enrollment.Student.Id, enrollment.Corsue.CourseId)
}

// FindAll implements EnrollmentService.
func (s *EnrollmentServiceImpl) FindAll(ctx context.Context) ([]*Model.Enrollment, error) {
	list := s.enrollmentRepo.FindAll(ctx)
	var response []*Model.Enrollment
	for _, value := range list {
		enrollment := &Model.Enrollment{
			Student: Model.Student{
				Id:        value.Student.Id,
				Name:      value.Student.Name,
				Surname:   value.Student.Surname,
				Email:     value.Student.Email,
				Address:   value.Student.Address,
				Birthdate: value.Student.Birthdate,
			},
			Corsue: Model.Course{
				CourseId:     value.Corsue.CourseId,
				CourseName:   value.Corsue.CourseName,
				StartYear:    value.Corsue.StartYear,
				EndYear:      value.Corsue.EndYear,
				Description:  value.Corsue.Description,
				DepartmentId: value.Corsue.DepartmentId,
				ProfessorId:  value.Corsue.ProfessorId,
			},
		}
		response = append(response, enrollment)
	}
	return response, nil
}

func NewEnrollmentServiceImpl(enrollmentrepo repository.EnrollmentRepo) EnrollmentService {
	return &EnrollmentServiceImpl{enrollmentRepo: enrollmentrepo}
}
