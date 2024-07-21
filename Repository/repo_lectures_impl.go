package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "go-server/Model"
	"go-server/helper"

	"github.com/lib/pq"
)

type lectureRepo struct{ db *sql.DB }

// Delete implements CourseRepo.
func (s *lectureRepo) Delete(ctx context.Context, name string) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Delete from lectures where name = $1"
	_, err = tx.ExecContext(ctx, SQL, name)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// FindAll implements CourseRepo.
func (s *lectureRepo) FindAll(ctx context.Context) []model.Lectures {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Select * from lectures"
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()
	var lectures []model.Lectures
	for result.Next() {
		lecture := model.Lectures{}
		err = result.Scan(&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.StartYear, &lecture.EndYear, &lecture.ProfessorId, &lecture.DepartmentId)
		helper.PanicIfError(err)
		lectures = append(lectures, lecture)
	}
	return lectures
}

// FindById implements LectureRepo.
func (s *lectureRepo) FindById(ctx context.Context, name string) (model.Lectures, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "Select * from lectures where name = $1"
	result, err := tx.QueryContext(ctx, sql, name)
	helper.PanicIfError(err)
	defer result.Close()
	lectures := model.Lectures{}
	if result.Next() {
		err = result.Scan(&lectures.LectureId, &lectures.LectureName, &lectures.Description, &lectures.StartYear, &lectures.EndYear, &lectures.ProfessorId, &lectures.DepartmentId)
		helper.PanicIfError(err)
		return lectures, nil
	} else {
		return lectures, errors.New("course not found")
	}
}
func (s *lectureRepo) LectureExists(ctx context.Context, lecture *model.Lectures) (bool, error) {
	var existingId int
	err := s.db.QueryRow("SELECT lectureid FROM lectures WHERE name = $1", lecture.LectureName).Scan(&existingId)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, handleSQLError(err)
	}
	return existingId > 0, nil
}

// Save implements CourseRepo.
func (s *lectureRepo) Save(ctx context.Context, lectures model.Lectures) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	SQL := "insert into lectures(name, startyear, endyear, description, professorid, departmentid) values ($1,$2,$3,$4,$5,$6)"
	_, err = tx.ExecContext(ctx, SQL, lectures.LectureName, lectures.StartYear, lectures.EndYear, lectures.Description, lectures.ProfessorId, lectures.DepartmentId)
	helper.PanicIfError(err)

	if err != nil {
		return false, handleCourseSQLError(err)
	} else {
		return true, nil
	}
}

// Update implements CourseRepo.
func (s *lectureRepo) Update(ctx context.Context, lectures model.Lectures) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "update lectures set startyear=$1, endyear=$2, description=$3, professorid=$4, departmentid=$5 where name=$6"
	res, err := tx.ExecContext(ctx, sql, lectures.StartYear, lectures.EndYear, lectures.Description, lectures.ProfessorId, lectures.DepartmentId, lectures.LectureName)
	if err != nil {
		return false, errors.New("error during update")
	}
	count, errx := res.RowsAffected()
	handleCourseSQLError(errx)
	if count == 0 {
		return false, errors.New("course not found")
	}
	if err != nil {
		return false, handleCourseSQLError(err)
	} else {
		return true, nil
	}
}

func NewLectureRepo(db *sql.DB) *lectureRepo {
	return &lectureRepo{db: db}
}

func handleCourseSQLError(err error) error {
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code {
		case "23505":
			return errors.New("a course with this identifier already exists")
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

func (s *lectureRepo) FindByIds(ctx context.Context, lecturename string, departmentname string) ([]model.DepartmentLecture, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	department := model.Department{}
	sql1 := "select departmentid, name from departments where name = $1"
	err = tx.QueryRowContext(ctx, sql1, departmentname).Scan(
		&department.DepartmentId, &department.Name)
	if err == sql.ErrNoRows {
		return []model.DepartmentLecture{}, errors.New("department not found")
	} else if err != nil {
		return []model.DepartmentLecture{}, handleSQLError(err)
	}
	lecture := model.Lectures{}
	sql2 := "select * from lectures where name = $1"
	err = tx.QueryRowContext(ctx, sql2, lecturename).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err == sql.ErrNoRows {
		return []model.DepartmentLecture{}, errors.New("lecture not found")
	} else if err != nil {
		return []model.DepartmentLecture{}, handleSQLError(err)
	}
	SQL := `
	select departments.name, lectures.name, lectures.description, lectures.startyear, lectures.endyear 
	from departments 
	join lectures on departments.departmentid = lectures.departmentid 
	where departments.departmentid = $1 and lectures.lectureid = $2
	`
	var departmentLectures []model.DepartmentLecture
	result, err := tx.QueryContext(ctx, SQL, &department.DepartmentId, &lecture.LectureId)
	if err != nil {
		return departmentLectures, err
	}
	defer result.Close()

	for result.Next() {
		departmentLecture := model.DepartmentLecture{}
		err := result.Scan(&departmentLecture.Department.Name, &departmentLecture.Lectures.LectureName, &departmentLecture.Lectures.Description, &departmentLecture.Lectures.StartYear, &departmentLecture.Lectures.EndYear)
		helper.PanicIfError(err)
		departmentLectures = append(departmentLectures, departmentLecture)
	}
	if err != nil {
		return departmentLectures, errors.New("DepartmentLectures not found")
	}
	return departmentLectures, nil
}
