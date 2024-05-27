package course

import (
	"context"
	"go-server/data/request"
	"go-server/data/response"
)

type CourseService interface {
	Create(ctx context.Context, request request.CourseCreateRequest)
	Update(ctx context.Context, request request.CourseUpdateRequest)
	Delete(ctx context.Context, idcourse int)
	FindById(ctx context.Context, idcourse int) response.CourseResponse
	FindAll(ctx context.Context) []response.CourseResponse
}
