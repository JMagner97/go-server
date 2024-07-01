package repository

import (
	"context"
	model "go-server/Model"
)

type StudentRepo interface {
	Save(ctx context.Context, student model.Student) (bool, error)
	Update(ctx context.Context, student model.Student) (bool, error)
	Delete(ctx context.Context, idstudent int) (bool, error)
	FindById(ctx context.Context, email string) (model.Student, error)
	FindAll(ctx context.Context) []model.Student
	StudentExists(ctx context.Context, student *model.Student) (bool, error)
}

type LectureRepo interface {
	Save(ctx context.Context, lecture model.Lectures) (bool, error)
	Update(ctx context.Context, lecture model.Lectures) (bool, error)
	Delete(ctx context.Context, name string) (bool, error)
	FindById(ctx context.Context, name string) (model.Lectures, error)
	FindAll(ctx context.Context) []model.Lectures
	LectureExists(ctx context.Context, lecture *model.Lectures) (bool, error)
}

type StudentLectureRepo interface {
	Save(ctx context.Context, email string, name string) (bool, error)
	Delete(ctx context.Context, email string, name string) (bool, error)
	FindAll(ctx context.Context) []model.StudentLectures
	EnrollmentExist(ctx context.Context, email string, name string) (bool, error)
	FindById(ctx context.Context, email string, name string) (model.StudentLectures, error)
}

type UserRepo interface {
	VerifyCredentials(username string, password string) bool
	VerifyIsAuthenticated(tokenString string) bool
	UpdateToken(username string, token string) bool
	Logout(tokenString string) bool
	Signup(username string, password string) (bool, error)
	VerifyUsername(username string) bool
}
