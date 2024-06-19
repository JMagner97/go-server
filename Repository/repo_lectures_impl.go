package repository

import (
	"context"
	"database/sql"
	"errors"
	model "go-server/Model"
	"go-server/helper"
)

type lectureRepo struct{ db *sql.DB }

// Delete implements CourseRepo.
func (s *lectureRepo) Delete(ctx context.Context, idcourse int) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Delete from lectures where lectureid = $1"
	_, err = tx.ExecContext(ctx, SQL, idcourse)
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
func (s *lectureRepo) FindById(ctx context.Context, idcourse int) (model.Lectures, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "Select * from lectures where lectureid = $1"
	result, err := tx.QueryContext(ctx, sql, idcourse)
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

// Save implements CourseRepo.
func (s *lectureRepo) Save(ctx context.Context, lectures model.Lectures) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	SQL := "insert into lectures(lectureid, name, startyear, endyear, description, professorid, departmentid) values ($1,$2,$3,$4,$5,$6,$7)"
	_, err = tx.ExecContext(ctx, SQL, lectures.LectureId, lectures.LectureName, lectures.StartYear, lectures.EndYear, lectures.Description, lectures.ProfessorId, lectures.DepartmentId)
	helper.PanicIfError(err)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Update implements CourseRepo.
func (s *lectureRepo) Update(ctx context.Context, lectures model.Lectures) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := "update lectures set name=$1, startyear=$2, endyear=$3, description=$4, professorid=$5, departmentid=$6 where lectureid=$7"
	_, err = tx.ExecContext(ctx, sql, lectures.LectureName, lectures.StartYear, lectures.EndYear, lectures.Description, lectures.ProfessorId, lectures.DepartmentId, lectures.LectureId)
	helper.PanicIfError(err)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func NewLectureRepo(db *sql.DB) *lectureRepo {
	return &lectureRepo{db: db}
}
