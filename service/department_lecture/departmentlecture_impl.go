package departmentlecturego

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
)

type deparmentlectureimpl struct {
	departmentlecRepo *repository.DepartmentLectureRepo
}

// FindById implements DepartmentLectureService.
// FindById implements DepartmentLectureService.
func (s *deparmentlectureimpl) FindById(ctx context.Context, lecturename string, departmentname string) ([]*Model.DepartmentLecture, error) {
	result, err := (*s.departmentlecRepo).FindByIds(ctx, lecturename, departmentname)
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

func NewDepartmentLectureGo(departmentlecRepos *repository.DepartmentLectureRepo) DepartmentLectureService {
	return &deparmentlectureimpl{departmentlecRepo: departmentlecRepos}
}
