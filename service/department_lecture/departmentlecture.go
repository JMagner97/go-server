package departmentlecturego

import (
	"context"
	"go-server/Model"
)

type DepartmentLectureService interface {
	FindById(sctx context.Context, lname string, dname string) ([]*Model.DepartmentLecture, error)
}
