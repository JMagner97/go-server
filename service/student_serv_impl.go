package service

import (
	"context"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
	"go-server/data/response"
	"go-server/helper"
)

type StudentServiceImpl struct {
	StudentRepo repository.StudentRepo
}

func NewStudentServiceImpl(studentrepo repository.StudentRepo) StudentService {
	return &StudentServiceImpl{StudentRepo: studentrepo}
}

// Create implements StudentService.
func (b *StudentServiceImpl) Create(ctx context.Context, request request.StudentCreateRequest) {
	student := Model.Student{
		Name:      request.Name,
		Id:        request.Id,
		Surname:   request.Surname,
		Birthdate: request.Data,
		Email:     request.Email,
		Address:   request.Address,
	}
	b.StudentRepo.Save(ctx, student)
}

// Delete implements StudentService.
func (b *StudentServiceImpl) Delete(ctx context.Context, studentid int) {
	student, err := b.StudentRepo.FindById(ctx, studentid)
	helper.PanicIfError(err)
	b.StudentRepo.Delete(ctx, student.Id)
}

// FindAll implements StudentService.
func (b *StudentServiceImpl) FindAll(ctx context.Context) []response.StudentResponse {
	student := b.StudentRepo.FindAll(ctx)
	var studentRespo []response.StudentResponse
	for _, value := range student {
		student := response.StudentResponse{Id: value.Id, Name: value.Name, Surname: value.Surname, Data: value.Birthdate, Address: value.Address, Email: value.Email}
		studentRespo = append(studentRespo, student)
	}
	return studentRespo
}

// FindById implements StudentService.
func (b *StudentServiceImpl) FindById(ctx context.Context, studentid int) response.StudentResponse {
	student, err := b.StudentRepo.FindById(ctx, studentid)
	helper.PanicIfError(err)
	studentResponse := response.StudentResponse{
		Id:      student.Id,
		Name:    student.Name,
		Surname: student.Surname,
		Data:    student.Birthdate,
		Address: student.Address,
		Email:   student.Email,
	}
	return studentResponse
}

// Update implements StudentService.
func (b *StudentServiceImpl) Update(ctx context.Context, request request.StudentUpdateRequest) {
	student, err := b.StudentRepo.FindById(ctx, request.Id)
	helper.PanicIfError(err)
	student.Name = request.Name
	student.Surname = request.Surname
	student.Birthdate = request.Data
	student.Address = request.Address
	student.Email = request.Email
	b.StudentRepo.Update(ctx, student)
}
