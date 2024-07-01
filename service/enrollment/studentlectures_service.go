package enrollment

import (
	"context"
	"go-server/Model"
	"go-server/data/response"
)

type StudentLectureService interface {
	Create(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error)
	FindAll(ctx context.Context) ([]*Model.StudentLectures, error)
	Delete(ctx context.Context, email string, name string) (bool, error)
	FindById(ctx context.Context, studentlecture *Model.StudentLectures) (response.EnrollmentResponse, error)
}
