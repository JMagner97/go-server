package lectures

import (
	"context"
	"go-server/data/request"
	"go-server/data/response"
)

type LectureService interface {
	Create(ctx context.Context, request request.LectureCreateRequest) (bool, error)
	Update(ctx context.Context, request request.LectureUpdateRequest, name string) (bool, error)
	Delete(ctx context.Context, name string) (bool, error)
	FindById(ctx context.Context, name string) (response.LectureResponse, error)
	FindAll(ctx context.Context) []response.LectureResponse
}
