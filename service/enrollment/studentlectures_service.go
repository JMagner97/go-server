package enrollment

import (
	"context"
	"go-server/Model"
)

type StudentLectureService interface {
	Create(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error)
	FindAll(ctx context.Context) ([]*Model.StudentLectures, error)
	Delete(ctx context.Context, studentlecture *Model.StudentLectures) (bool, error)
}
