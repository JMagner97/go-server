package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-server/Model"
	"go-server/helper"
)

// struttura contenente un puntatore al db
type repo_stud_impl struct {
	Db *sql.DB
}

// accetta un argomento a puntatore e restituisce un'istanza del repository
func NewStudRepo(Db *sql.DB) StudentRepo {
	return &repo_stud_impl{Db: Db}
}

// Delete implements StudentRepo.
func (s *repo_stud_impl) Delete(ctx context.Context, idstudent int) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Delete from students where studentid = $1"
	_, errx := tx.ExecContext(ctx, SQL, idstudent)
	helper.PanicIfError(errx)
}

// FindAll implements StudentRepo.
func (s *repo_stud_impl) FindAll(ctx context.Context) []Model.Student {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	SQL := "Select * from students"

	result, errx := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errx)
	defer result.Close()
	var students []Model.Student
	for result.Next() {
		student := Model.Student{}
		err := result.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
}

// FindById implements StudentRepo.
func (s *repo_stud_impl) FindById(ctx context.Context, idstudent int) (Model.Student, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Select * from students where studentid = $1"
	result, errx := tx.QueryContext(ctx, SQL, idstudent)
	helper.PanicIfError(errx)
	defer result.Close()
	student := Model.Student{}

	if result.Next() {
		err := result.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student not found")
	}

}

// Save implements StudentRepo.
func (s *repo_stud_impl) Save(ctx context.Context, student Model.Student) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	//var SQL string
	//var args []interface{}
	//	if student.CourseId == 0 {
	SQL := "insert into students(studentid,name,surname,birthdate,address,email) values ($1,$2,$3,$4,$5,$6)"
	//args = []interface{}{student.Id, student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.CourseId}
	//	} else {
	//		SQL = "insert into studenti(idstudente,nome,cognome,datanascita,indirizzo,email) values ($1,$2,$3,$4,$5,$6)"
	//		args = []interface{}{student.Id, student.Name, student.Surname, student.Birthdate, student.Address, student.Email}
	//	}
	_, err = tx.ExecContext(ctx, SQL, student.Id, student.Name, student.Surname, student.Birthdate, student.Address, student.Email)
	helper.PanicIfError(err)
}

// Update implements StudentRepo.
func (s *repo_stud_impl) Update(ctx context.Context, student Model.Student) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	//var SQL string
	//var args []interface{}
	//if student.CourseId == 0 {
	SQL := "update students set name=$1, surname=$2, birthdate=$3, address=$4, email=$5 where studentid=$6"
	//args = []interface{}{student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.CourseId, student.Id}
	//	} else {
	//		SQL = "update studenti set nome=$1, cognome=$2, datanascita=$3, indirizzo=$4, email=$5 where idstudente=$6"
	//		args = []interface{}{student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.Id}
	//	}
	_, err = tx.ExecContext(ctx, SQL, student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.Id)
	helper.PanicIfError(err)
}
