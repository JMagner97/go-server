package lectures

import (
	"context"
	"errors"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
	"go-server/data/response"
)

type LectureServiceImpl struct {
	lectureRepo repository.LectureRepo
}

// FindByIds implements LectureService.
func (s *LectureServiceImpl) FindByIds(ctx context.Context, lecname string, depname string) ([]*Model.DepartmentLecture, error) {
	result, err := s.lectureRepo.FindByIds(ctx, lecname, depname)
	var responses []*Model.DepartmentLecture
	for _, value := range result {
		departmentlec := &Model.DepartmentLecture{
			Department: Model.Department{
				Name: value.Department.Name,
			},
			Lectures: Model.Lectures{
				LectureName: value.Lectures.LectureName,
				Description: value.Lectures.Description,
				StartYear:   value.Lectures.StartYear,
				EndYear:     value.Lectures.EndYear,
			},
		}
		responses = append(responses, departmentlec)
	}

	// Return the response
	return responses, err
}

// Delete implements CourseService.
func (s *LectureServiceImpl) Delete(ctx context.Context, name string) (bool, error) {
	lecture, err := s.lectureRepo.FindById(ctx, name)
	//helper.PanicIfError(err)
	if err != nil {
		return false, err
	}
	result, errx := s.lectureRepo.Delete(ctx, lecture.LectureName)
	return result, errx
}

// FindAll implements CourseService.
func (s *LectureServiceImpl) FindAll(ctx context.Context) []response.LectureResponse {
	lecture := s.lectureRepo.FindAll(ctx)
	var lectureRespo []response.LectureResponse
	for _, value := range lecture {
		lecture := response.LectureResponse{
			//LectureId:    value.LectureId,
			LectureName:  value.LectureName,
			StartYear:    value.StartYear,
			EndYear:      value.EndYear,
			Description:  value.Description,
			ProfessorId:  value.ProfessorId,
			DepartmentId: value.DepartmentId,
		}
		lectureRespo = append(lectureRespo, lecture)
	}
	return lectureRespo
}

// FindById implements CourseService.
func (s *LectureServiceImpl) FindById(ctx context.Context, name string) (response.LectureResponse, error) {
	lecture, err := s.lectureRepo.FindById(ctx, name)
	//helper.PanicIfError(err)
	lectureResponse := response.LectureResponse{
		//LectureId:    lecture.LectureId,
		LectureName:  lecture.LectureName,
		StartYear:    lecture.StartYear,
		EndYear:      lecture.EndYear,
		Description:  lecture.Description,
		ProfessorId:  lecture.ProfessorId,
		DepartmentId: lecture.DepartmentId,
	}
	return lectureResponse, err
}

// Save implements CourseService.
func (s *LectureServiceImpl) Create(ctx context.Context, request request.LectureCreateRequest) (bool, error) {
	lecture := Model.Lectures{
		//LectureId:    request.LectureId,
		LectureName:  request.LectureName,
		StartYear:    request.StartYear,
		EndYear:      request.EndYear,
		Description:  request.Description,
		ProfessorId:  request.ProfessorId,
		DepartmentId: request.DepartmentId,
	}
	exists, err := s.lectureRepo.LectureExists(ctx, &lecture)
	if err != nil {
		return false, err
	}
	if !exists {
		result, err := s.lectureRepo.Save(ctx, lecture)
		return result, err
	} else {
		return false, errors.New("lecture already exists")
	}
}

// Update implements CourseService.
func (s *LectureServiceImpl) Update(ctx context.Context, request request.LectureUpdateRequest, name string) (bool, error) {
	lecture, err := s.lectureRepo.FindById(ctx, name)
	//helper.PanicIfError(err)
	if err != nil {
		return false, err
	}
	lecture.LectureName = name
	lecture.StartYear = request.StartYear
	lecture.EndYear = request.EndYear
	lecture.Description = request.Description
	lecture.ProfessorId = request.ProfessorId
	lecture.DepartmentId = request.DepartmentId
	result, errx := s.lectureRepo.Update(ctx, lecture)
	return result, errx
}

func NewLectureServiceImpl(lecturerepo repository.LectureRepo) LectureService {
	return &LectureServiceImpl{lectureRepo: lecturerepo}
}
