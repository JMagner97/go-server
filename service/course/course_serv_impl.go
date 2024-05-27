package course

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
)

type CourseServiceImpl struct {
	CourseRepo repository.CourseRepo
}

// Delete implements CourseService.
func (s *CourseServiceImpl) Delete(ctx context.Context, idcourse int) {
	course, err := s.CourseRepo.FindById(ctx, idcourse)
	helper.PanicIfError(err)
	s.CourseRepo.Delete(ctx, course.CourseId)
}

// FindAll implements CourseService.
func (s *CourseServiceImpl) FindAll(ctx context.Context) []response.CourseResponse {
	course := s.CourseRepo.FindAll(ctx)
	var courseRespo []response.CourseResponse
	for _, value := range course {
		course := response.CourseResponse{CourseId: value.CourseId, CourseName: value.CourseName, StartYear: value.StartYear, EndYear: value.EndYear, Description: value.Description, DepartmentId: value.DepartmentId, ProfessorId: value.ProfessorId}
		courseRespo = append(courseRespo, course)
	}
	return courseRespo
}

// FindById implements CourseService.
func (s *CourseServiceImpl) FindById(ctx context.Context, idcourse int) response.CourseResponse {
	course, err := s.CourseRepo.FindById(ctx, idcourse)
	helper.PanicIfError(err)
	courseResponse := response.CourseResponse{
		CourseId:     course.CourseId,
		CourseName:   course.CourseName,
		StartYear:    course.StartYear,
		EndYear:      course.EndYear,
		Description:  course.Description,
		DepartmentId: course.DepartmentId,
		ProfessorId:  course.ProfessorId,
	}
	return courseResponse
}

// Save implements CourseService.
func (s *CourseServiceImpl) Create(ctx context.Context, request request.CourseCreateRequest) {
	course := Model.Course{
		CourseId:     request.CourseId,
		CourseName:   request.CourseName,
		StartYear:    request.StartYear,
		EndYear:      request.EndYear,
		Description:  request.Description,
		DepartmentId: request.DepartmentId,
		ProfessorId:  request.ProfessorId,
	}
	s.CourseRepo.Save(ctx, course)
}

// Update implements CourseService.
func (s *CourseServiceImpl) Update(ctx context.Context, request request.CourseUpdateRequest) {
	course, err := s.CourseRepo.FindById(ctx, request.CourseId)
	helper.PanicIfError(err)
	course.CourseName = request.CourseName
	course.StartYear = request.StartYear
	course.EndYear = request.EndYear
	course.Description = request.Description
	course.DepartmentId = request.DepartmentId
	course.ProfessorId = request.ProfessorId
	s.CourseRepo.Update(ctx, course)
}

func NewCourseServiceImpl(courserepo repository.CourseRepo) CourseService {
	return &CourseServiceImpl{CourseRepo: courserepo}
}
