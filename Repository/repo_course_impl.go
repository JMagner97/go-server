package repository

import (
	"context"
	"database/sql"
	"errors"
	model "go-server/Model"
	"go-server/helper"
)

type courseRepo struct{ db *sql.DB }

// Delete implements CourseRepo.
func (s *courseRepo) Delete(ctx context.Context, idcourse int) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Delete from courses where courseid = $1"
	_, err = tx.ExecContext(ctx, SQL, idcourse)
	helper.PanicIfError(err)
}

// FindAll implements CourseRepo.
func (s *courseRepo) FindAll(ctx context.Context) []model.Course {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Select * from courses"
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()
	var courses []model.Course
	for result.Next() {
		course := model.Course{}
		err = result.Scan(&course.CourseId, &course.CourseName, &course.Description, &course.StartYear, &course.EndYear, &course.DepartmentId, &course.ProfessorId)
		helper.PanicIfError(err)
		courses = append(courses, course)
	}
	return courses
}

// FindById implements CourseRepo.
func (s *courseRepo) FindById(ctx context.Context, idcourse int) (model.Course, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "Select * from courses where courseid = $1"
	result, err := tx.QueryContext(ctx, sql, idcourse)
	helper.PanicIfError(err)
	defer result.Close()
	course := model.Course{}
	if result.Next() {
		err = result.Scan(&course.CourseId, &course.CourseName, &course.Description, &course.StartYear, &course.EndYear, &course.DepartmentId, &course.ProfessorId)
		helper.PanicIfError(err)
		return course, nil
	} else {
		return course, errors.New("course not found")
	}
}

// Save implements CourseRepo.
func (s *courseRepo) Save(ctx context.Context, course model.Course) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	SQL := "insert into courses(courseid, coursename, startyear, endyear, description, departmentid, professorid) values ($1,$2,$3,$4,$5,$6,$7)"
	_, err = tx.ExecContext(ctx, SQL, course.CourseId, course.CourseName, course.StartYear, course.EndYear, course.Description, course.DepartmentId, course.ProfessorId)
	helper.PanicIfError(err)

}

// Update implements CourseRepo.
func (s *courseRepo) Update(ctx context.Context, course model.Course) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "update courses set coursename=$1, startyear=$2, endyear=$3, description=$4, departmentid=$5, professorid=$6 where courseid=$7"
	_, err = tx.ExecContext(ctx, sql, course.CourseName, course.StartYear, course.EndYear, course.Description, course.DepartmentId, course.ProfessorId, course.CourseId)
	helper.PanicIfError(err)
}

func NewCourseRepo(db *sql.DB) CourseRepo {
	return &courseRepo{db: db}
}
