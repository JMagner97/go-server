package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-server/Model"
	"go-server/helper"

	"github.com/lib/pq"
)

type studentlectureRepo struct{ db *sql.DB }

// Delete implements EnrollmentRepo.
func (s *studentlectureRepo) Delete(ctx context.Context, email string, name string) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	student := Model.Student{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM students WHERE email = $1", email).Scan(
		&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId,
	)
	if err == sql.ErrNoRows {
		return false, errors.New("student not found")
	} else if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	// Retrieve lecture by name
	lecture := Model.Lectures{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM lectures WHERE name = $1", name).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err == sql.ErrNoRows {
		return false, errors.New("lecture not found")
	} else if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	sql := `DELETE FROM studentlectures
			WHERE studentid = $1 and lectureid = $2`
	_, err = tx.ExecContext(ctx, sql, student.Id, lecture.LectureId)
	if err != nil {
		return false, handleStudentLectureSQLError(err)
	} else {
		return true, nil
	}
}

// FindAll implements EnrollmentRepo.
func (s *studentlectureRepo) FindAll(ctx context.Context) []Model.StudentLectures {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := `
	SELECT 
    students.name AS student_name,
    students.surname,
    students.email,
    lectures.name AS lecture_name,
    lectures.startyear,
    lectures.endyear,
    professors.surname AS professors_surname,
    departments.name AS department_name
	FROM 
   	 studentlectures
	JOIN 
    	students ON studentlectures.studentid = students.studentid
	JOIN 
    	lectures ON studentlectures.lectureid = lectures.lectureid
	JOIN 
    	professors ON lectures.professorid = professors.professorid
	JOIN 
    	departments ON lectures.departmentid = departments.departmentid;
`
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()
	var studentlectures []Model.StudentLectures
	for result.Next() {
		studentlecture := Model.StudentLectures{}
		err := result.Scan(&studentlecture.Student.Name, &studentlecture.Student.Surname, &studentlecture.Student.Email, &studentlecture.Lecture.LectureName, &studentlecture.Lecture.StartYear, &studentlecture.Lecture.EndYear, &studentlecture.Professor.Surname, &studentlecture.Department.Name)
		helper.PanicIfError(err)
		studentlectures = append(studentlectures, studentlecture)
	}
	return studentlectures
}

// FindById implements StudentLectureRepo.
func (s *studentlectureRepo) FindById(ctx context.Context, email string, name string) (Model.StudentLectures, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	student := Model.Student{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM students WHERE email = $1", email).Scan(
		&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId,
	)
	if err == sql.ErrNoRows {
		return Model.StudentLectures{}, errors.New("student not found")
	} else if err != nil {
		return Model.StudentLectures{}, handleStudentLectureSQLError(err)
	}
	lecture := Model.Lectures{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM lectures WHERE name = $1", name).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err == sql.ErrNoRows {
		return Model.StudentLectures{}, errors.New("lecture not found")
	} else if err != nil {
		return Model.StudentLectures{}, handleStudentLectureSQLError(err)
	}
	SQL := `
	SELECT 
    students.name AS student_name,
    students.surname,
    students.email,
    lectures.name AS lecture_name,
    lectures.startyear,
    lectures.endyear,
    professors.surname AS professors_surname,
    departments.name AS department_name
	FROM 
   	 studentlectures
	JOIN 
    	students ON studentlectures.studentid = students.studentid
	JOIN 
    	lectures ON studentlectures.lectureid = lectures.lectureid
	JOIN 
    	professors ON lectures.professorid = professors.professorid
	JOIN 
    	departments ON lectures.departmentid = departments.departmentid
	where students.studentid = $1 and lectures.lectureid = $2
`
	var studentlecture Model.StudentLectures
	result, err := tx.QueryContext(ctx, SQL, student.Id, lecture.LectureId)
	handleStudentLectureSQLError(err)
	if err != nil {
		return studentlecture, errors.New("enrollment not found")
	}
	defer result.Close()

	for result.Next() {
		studentlecture := Model.StudentLectures{}
		err := result.Scan(&studentlecture.Student.Name, &studentlecture.Student.Surname, &studentlecture.Student.Email, &studentlecture.Lecture.LectureName, &studentlecture.Lecture.StartYear, &studentlecture.Lecture.EndYear, &studentlecture.Professor.Surname, &studentlecture.Department.Name)
		//helper.PanicIfError(err)
		return studentlecture, handleStudentLectureSQLError(err)
	}
	return studentlecture, errors.New("enrollment not found")
}

// EnrollmentExist implements StudentLectureRepo.
func (s *studentlectureRepo) EnrollmentExist(ctx context.Context, email string, name string) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	student := Model.Student{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM students WHERE email = $1", email).Scan(
		&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId,
	)
	if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	// Retrieve lecture by name
	lecture := Model.Lectures{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM lectures WHERE name = $1", name).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	var existingIds int
	err = s.db.QueryRow("SELECT COUNT(*) FROM studentlectures where studentid=$1 and lectureid=$2", student.Id, lecture.LectureId).Scan(&existingIds)
	handleStudentLectureSQLError(err)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	return existingIds > 0, nil

}

// Save implements EnrollmentRepo.
func (s *studentlectureRepo) Save(ctx context.Context, email string, name string) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	student := Model.Student{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM students WHERE email = $1", email).Scan(
		&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId,
	)
	if err == sql.ErrNoRows {
		return false, errors.New("student not found")
	} else if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	// Retrieve lecture by name
	lecture := Model.Lectures{}
	err = tx.QueryRowContext(ctx, "SELECT * FROM lectures WHERE name = $1", name).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err == sql.ErrNoRows {
		return false, errors.New("lecture not found")
	} else if err != nil {
		return false, handleStudentLectureSQLError(err)
	}
	sql := `INSERT INTO studentlectures (studentid, lectureid) VALUES ($1,$2)`
	_, err = tx.ExecContext(ctx, sql, student.Id, lecture.LectureId)
	if err != nil {
		return false, handleStudentLectureSQLError(err)
	} else {
		return true, nil
	}
}

func NewStudentLectureRepo(db *sql.DB) StudentLectureRepo {
	return &studentlectureRepo{db: db}
}

func handleStudentLectureSQLError(err error) error {
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code {
		case "23505":
			return errors.New("this student-lecture association already exists")
		case "23503":
			return errors.New("foreign key violation")
		case "23502":
			return errors.New("not null violation")
		case "23514":
			return errors.New("check constraint violation")
		case "22001":
			return errors.New("string data right truncation")
		case "22003":
			return errors.New("numeric value out of range")
		case "22012":
			return errors.New("division by zero")
		default:
			return fmt.Errorf("database error: %w", pgErr)
		}
	}
	return err
}
