package lectures

import (
	"context"
	"go-server/data/request"
	"go-server/data/response"
)

type LectureService interface {
	Create(ctx context.Context, request request.LectureCreateRequest) (bool, error)
	Update(ctx context.Context, request request.LectureUpdateRequest) (bool, error)
	Delete(ctx context.Context, lectureid int) (bool, error)
	FindById(ctx context.Context, lectureid int) (response.LectureResponse, error)
	FindAll(ctx context.Context) []response.LectureResponse
}
