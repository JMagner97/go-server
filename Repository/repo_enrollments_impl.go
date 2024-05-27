package repository

import (
	"context"
	"database/sql"
	"go-server/Model"
	"go-server/helper"
)

type enrollmentRepo struct{ db *sql.DB }

// Delete implements EnrollmentRepo.
func (s *enrollmentRepo) Delete(ctx context.Context, idstudent int, courseid int) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := `DELETE FROM enrollments
			WHERE studentid = $1 and courseid = $2`
	_, err = tx.ExecContext(ctx, sql, idstudent, courseid)
	helper.PanicIfError(err)
}

// FindAll implements EnrollmentRepo.
func (s *enrollmentRepo) FindAll(ctx context.Context) []Model.Enrollment {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := `Select name, surname, email,coursename,startyear, endyear 
			from enrollments 
				 join students on enrollments.studentid = students.studentid 
				 join courses on enrollments.courseid = courses.courseid`
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()
	var enrollments []Model.Enrollment
	for result.Next() {
		enrollment := Model.Enrollment{}
		err := result.Scan(&enrollment.Student.Name, &enrollment.Student.Surname, &enrollment.Corsue.CourseName, &enrollment.Student.Email, &enrollment.Corsue.StartYear, &enrollment.Corsue.EndYear)
		helper.PanicIfError(err)
		enrollments = append(enrollments, enrollment)
	}
	return enrollments
}

// Save implements EnrollmentRepo.
func (s *enrollmentRepo) Save(ctx context.Context, enrollment Model.Enrollment) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := `INSERT INTO enrollments (studentid, courseid) VALUES ($1,$2)`
	_, err = tx.ExecContext(ctx, sql, enrollment.Student.Id, enrollment.Corsue.CourseId)
	helper.PanicIfError(err)
}

func NewEnrollmentRepo(db *sql.DB) EnrollmentRepo {
	return &enrollmentRepo{db: db}
}
