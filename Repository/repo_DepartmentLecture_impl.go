package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-server/Model"
	"go-server/helper"
)

type DepartmentLecture struct {
	db *sql.DB
}

// FindByIds implements DepartmentLectureRepo.
func (s *DepartmentLecture) FindByIds(ctx context.Context, lecturename string, departmentname string) ([]Model.DepartmentLecture, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	department := Model.Department{}
	sql1 := "select * from departments where departmentid = $1"
	err = tx.QueryRowContext(ctx, sql1, departmentname).Scan(
		&department.DepartmentId, &department.Name)
	if err == sql.ErrNoRows {
		return []Model.DepartmentLecture{}, errors.New("department not found")
	} else if err != nil {
		return []Model.DepartmentLecture{}, handleSQLError(err)
	}
	lecture := Model.Lectures{}
	sql2 := "select * from lectures where name = $1"
	err = tx.QueryRowContext(ctx, sql2, lecturename).Scan(
		&lecture.LectureId, &lecture.LectureName, &lecture.Description, &lecture.ProfessorId, &lecture.DepartmentId, &lecture.StartYear, &lecture.EndYear,
	)
	if err == sql.ErrNoRows {
		return []Model.DepartmentLecture{}, errors.New("lecture not found")
	} else if err != nil {
		return []Model.DepartmentLecture{}, handleSQLError(err)
	}
	SQL := `
	select departments.name, lectures.name, lectures.description, lectures.startyear, lectures.endyear 
	from departments 
	join lectures on departments.departmentid = lectures.departmentid 
	where departments.name = $1 and lectures.name = $2
	`
	var departmentLectures []Model.DepartmentLecture
	result, err := tx.QueryContext(ctx, SQL, &department.DepartmentId, &lecture.LectureId)
	if err != nil {
		return departmentLectures, err
	}
	defer result.Close()

	for result.Next() {
		departmentLecture := Model.DepartmentLecture{}
		err := result.Scan(&departmentLecture.Department.Name, &departmentLecture.Lectures.LectureName, &departmentLecture.Lectures.Description, &departmentLecture.Lectures.StartYear, &departmentLecture.Lectures.EndYear)
		helper.PanicIfError(err)
		departmentLectures = append(departmentLectures, departmentLecture)
	}
	if err != nil {
		return departmentLectures, errors.New("DepartmentLectures not found")
	}
	return departmentLectures, nil
}

func NewDepartmentLectureRepository(db *sql.DB) DepartmentLectureRepo {
	return &DepartmentLecture{db: db}
}
