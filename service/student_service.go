package service

import (
	"context"
	"go-server/data/request"
	"go-server/data/response"
)

type StudentService interface {
	Create(ctx context.Context, request request.StudentCreateRequest)
	Update(ctx context.Context, request request.StudentUpdateRequest)
	Delete(ctx context.Context, studentid int)
	FindById(ctx context.Context, studentid int) response.StudentResponse
	FindAll(ctx context.Context) []response.StudentResponse
}
