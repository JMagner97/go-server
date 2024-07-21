package service

import (
	"context"
	"errors"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
	"go-server/data/response"
)

type StudentServiceImpl struct {
	StudentRepo repository.StudentRepo
}

func NewStudentServiceImpl(studentrepo repository.StudentRepo) StudentService {
	return &StudentServiceImpl{StudentRepo: studentrepo}
}

// Create implements StudentService.
func (b *StudentServiceImpl) Create(ctx context.Context, request request.StudentCreateRequest) (bool, error) {
	student := Model.Student{
		Name: request.Name,
		//Id:           request.Id,
		Username:     request.Username,
		Password:     request.Password,
		Role:         1,
		Surname:      request.Surname,
		Birthdate:    request.Data,
		Email:        request.Email,
		Address:      request.Address,
		DepartmentId: request.DepartmentId,
	}
	exists, err := b.StudentRepo.StudentExists(ctx, &student)
	if err != nil {
		return false, err
	}
	if !exists {
		innerResult, err := b.StudentRepo.Save(ctx, student)
		return innerResult, err
	} else {
		return false, errors.New("student already exists")
	}
}

// Delete implements StudentService.
func (b *StudentServiceImpl) Delete(ctx context.Context, email string) (bool, error) {
	student, err := b.StudentRepo.FindById(ctx, email)
	if err != nil {
		return false, err
	}
	result, errx := b.StudentRepo.Delete(ctx, student.Id)
	return result, errx
}

// FindAll implements StudentService.
func (b *StudentServiceImpl) FindAll(ctx context.Context) []response.StudentResponse {
	student := b.StudentRepo.FindAll(ctx)
	var studentRespo []response.StudentResponse
	for _, value := range student {
		student := response.StudentResponse{Name: value.Name, Surname: value.Surname, Data: value.Birthdate, Address: value.Address, Email: value.Email, DepartmentId: value.DepartmentId}
		studentRespo = append(studentRespo, student)
	}
	return studentRespo
}

// FindById implements StudentService.
func (b *StudentServiceImpl) FindById(ctx context.Context, email string) (response.StudentResponse, error) {
	student, err := b.StudentRepo.FindById(ctx, email)
	//helper.PanicIfError(err)
	studentResponse := response.StudentResponse{
		//Id:           student.Id,
		Username:     student.Username,
		Name:         student.Name,
		Surname:      student.Surname,
		Data:         student.Birthdate,
		Address:      student.Address,
		Email:        student.Email,
		DepartmentId: student.DepartmentId,
	}
	return studentResponse, err
}

// Update implements StudentService.
func (b *StudentServiceImpl) Update(ctx context.Context, request request.StudentUpdateRequest, email string) (bool, error) {
	student, err := b.StudentRepo.FindById(ctx, email)
	//helper.PanicIfError(err)
	if err != nil {
		return false, err
	}
	student.Name = request.Name
	student.Surname = request.Surname
	student.Birthdate = request.Data
	student.Address = request.Address
	student.Email = email
	student.DepartmentId = request.DepartmentId
	result, errx := b.StudentRepo.Update(ctx, student)
	return result, errx
}
