package lectures

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
	"go-server/data/response"
)

type LectureServiceImpl struct {
	lectureRepo repository.LectureRepo
}

// Delete implements CourseService.
func (s *LectureServiceImpl) Delete(ctx context.Context, idcourse int) (bool, error) {
	lecture, err := s.lectureRepo.FindById(ctx, idcourse)
	//helper.PanicIfError(err)
	if err != nil {
		return false, err
	}
	result, errx := s.lectureRepo.Delete(ctx, lecture.LectureId)
	return result, errx
}

// FindAll implements CourseService.
func (s *LectureServiceImpl) FindAll(ctx context.Context) []response.LectureResponse {
	lecture := s.lectureRepo.FindAll(ctx)
	var lectureRespo []response.LectureResponse
	for _, value := range lecture {
		lecture := response.LectureResponse{
			LectureId:   value.LectureId,
			LectureName: value.LectureName,
			StartYear:   value.StartYear,
			EndYear:     value.EndYear,
			Description: value.Description,
			ProfessorId: value.ProfessorId,
		}
		lectureRespo = append(lectureRespo, lecture)
	}
	return lectureRespo
}

// FindById implements CourseService.
func (s *LectureServiceImpl) FindById(ctx context.Context, idcourse int) (response.LectureResponse, error) {
	lecture, err := s.lectureRepo.FindById(ctx, idcourse)
	//helper.PanicIfError(err)
	lectureResponse := response.LectureResponse{
		LectureId:   lecture.LectureId,
		LectureName: lecture.LectureName,
		StartYear:   lecture.StartYear,
		EndYear:     lecture.EndYear,
		Description: lecture.Description,
		ProfessorId: lecture.ProfessorId,
	}
	return lectureResponse, err
}

// Save implements CourseService.
func (s *LectureServiceImpl) Create(ctx context.Context, request request.LectureCreateRequest) (bool, error) {
	lecture := Model.Lectures{
		LectureId:   request.LectureId,
		LectureName: request.LectureName,
		StartYear:   request.StartYear,
		EndYear:     request.EndYear,
		Description: request.Description,
		ProfessorId: request.ProfessorId,
	}
	result, err := s.lectureRepo.Save(ctx, lecture)
	return result, err
}

// Update implements CourseService.
func (s *LectureServiceImpl) Update(ctx context.Context, request request.LectureUpdateRequest) (bool, error) {
	lecture, err := s.lectureRepo.FindById(ctx, request.LectureId)
	//helper.PanicIfError(err)
	if err != nil {
		return false, err
	}
	lecture.LectureName = request.LectureName
	lecture.StartYear = request.StartYear
	lecture.EndYear = request.EndYear
	lecture.Description = request.Description
	lecture.ProfessorId = request.ProfessorId
	result, errx := s.lectureRepo.Update(ctx, lecture)
	return result, errx
}

func NewLectureServiceImpl(lecturerepo repository.LectureRepo) LectureService {
	return &LectureServiceImpl{lectureRepo: lecturerepo}
}
