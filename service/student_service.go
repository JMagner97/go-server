package service

import (
	"context"
	"go-server/data/request"
	"go-server/data/response"
)

type StudentService interface {
	Create(ctx context.Context, request request.StudentCreateRequest) (bool, error)
	Update(ctx context.Context, request request.StudentUpdateRequest) (bool, error)
	Delete(ctx context.Context, studentid int) (bool, error)
	FindById(ctx context.Context, studentid int) (response.StudentResponse, error)
	FindAll(ctx context.Context) []response.StudentResponse
}
