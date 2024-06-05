package repository

import (
	"context"
	model "go-server/Model"
)

type StudentRepo interface {
	Save(ctx context.Context, student model.Student) (bool, error)
	Update(ctx context.Context, student model.Student) (bool, error)
	Delete(ctx context.Context, idstudent int) (bool, error)
	FindById(ctx context.Context, idstudent int) (model.Student, error)
	FindAll(ctx context.Context) []model.Student
}

type LectureRepo interface {
	Save(ctx context.Context, lecture model.Lectures) (bool, error)
	Update(ctx context.Context, lecture model.Lectures) (bool, error)
	Delete(ctx context.Context, lectureid int) (bool, error)
	FindById(ctx context.Context, lectureid int) (model.Lectures, error)
	FindAll(ctx context.Context) []model.Lectures
}

type StudentLectureRepo interface {
	Save(ctx context.Context, enrollment model.StudentLectures) (bool, error)
	Delete(ctx context.Context, idstudent int, idcourse int) (bool, error)
	FindAll(ctx context.Context) []model.StudentLectures
}

type UserRepo interface {
	VerifyCredentials(username string, password string) bool
	VerifyIsAuthenticated(tokenString string) bool
	UpdateToken(username string, token string) bool
	Logout(tokenString string) bool
	Signup(username string, password string) (bool, error)
	VerifyUsername(username string) bool
}
