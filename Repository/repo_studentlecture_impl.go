package repository

import (
	"context"
	"database/sql"
	"go-server/Model"
	"go-server/helper"
)

type studentlectureRepo struct{ db *sql.DB }

// Delete implements EnrollmentRepo.
func (s *studentlectureRepo) Delete(ctx context.Context, idstudent int, lectureid int) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := `DELETE FROM studentlectures
			WHERE studentid = $1 and lectureid = $2`
	_, err = tx.ExecContext(ctx, sql, idstudent, lectureid)
	if err != nil {
		return false, err
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
	Select students.name, surname, email,lectures.name,startyear, endyear 
				from studentlectures 
					 join students on studentlectures.studentid = students.studentid 
					 join lectures on studentlectures.lectureid = lectures.lectureid`
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()
	var studentlectures []Model.StudentLectures
	for result.Next() {
		studentlecture := Model.StudentLectures{}
		err := result.Scan(&studentlecture.Student.Name, &studentlecture.Student.Surname, &studentlecture.Lecture.LectureName, &studentlecture.Student.Email, &studentlecture.Lecture.StartYear, &studentlecture.Lecture.EndYear)
		helper.PanicIfError(err)
		studentlectures = append(studentlectures, studentlecture)
	}
	return studentlectures
}

// Save implements EnrollmentRepo.
func (s *studentlectureRepo) Save(ctx context.Context, studentlecture Model.StudentLectures) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	sql := `INSERT INTO studentlectures (studentid, lectureid) VALUES ($1,$2)`
	_, err = tx.ExecContext(ctx, sql, studentlecture.Student.Id, studentlecture.Lecture.LectureId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func NewStudentLectureRepo(db *sql.DB) StudentLectureRepo {
	return &studentlectureRepo{db: db}
}
