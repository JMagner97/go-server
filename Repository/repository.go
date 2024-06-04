package repository

import (
	"context"
	model "go-server/Model"
)

type StudentRepo interface {
	Save(ctx context.Context, student model.Student)
	Update(ctx context.Context, student model.Student)
	Delete(ctx context.Context, idstudent int)
	FindById(ctx context.Context, idstudent int) (model.Student, error)
	FindAll(ctx context.Context) []model.Student
}

type CourseRepo interface {
	Save(ctx context.Context, course model.Course)
	Update(ctx context.Context, course model.Course)
	Delete(ctx context.Context, idcourse int)
	FindById(ctx context.Context, idcourse int) (model.Course, error)
	FindAll(ctx context.Context) []model.Course
}

type EnrollmentRepo interface {
	Save(ctx context.Context, enrollment model.Enrollment)
	Delete(ctx context.Context, idstudent int, idcourse int)
	FindAll(ctx context.Context) []model.Enrollment
}

type UserRepo interface {
	VerifyCredentials(username string, password string) bool
	VerifyIsAuthenticated(tokenString string) bool
	UpdateToken(username string, token string) bool
	Logout(tokenString string) bool
	Signup(username string, password string) (bool, error)
	VerifyUsername(username string) bool
}
