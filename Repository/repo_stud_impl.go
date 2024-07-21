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

// struttura contenente un puntatore al db
type repo_stud_impl struct {
	Db *sql.DB
}

// accetta un argomento a puntatore e restituisce un'istanza del repository
func NewStudRepo(Db *sql.DB) StudentRepo {
	return &repo_stud_impl{Db: Db}
}

// Delete implements StudentRepo.
func (s *repo_stud_impl) Delete(ctx context.Context, idstudent int) (bool, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Delete from students where studentid = $1"
	_, errx := tx.ExecContext(ctx, SQL, idstudent)
	if errx != nil {
		return false, handleSQLError(errx)
	} else {
		return true, nil
	}
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
		err := result.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
} //pagination

// FindById implements StudentRepo.
func (s *repo_stud_impl) FindById(ctx context.Context, email string) (Model.Student, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Select * from students where email = $1"
	result, errx := tx.QueryContext(ctx, SQL, email)
	//helper.PanicIfError(errx)

	student := Model.Student{}
	if errx != nil {
		return student, errors.New("student not found")
	}
	defer result.Close()
	if result.Next() {
		err := result.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.Address, &student.Email, &student.DepartmentId)
		//helper.PanicIfError(err)
		return student, handleSQLError(err)
	} else {
		return student, errors.New("student not found")
	}

}

func (s *repo_stud_impl) StudentExists(ctx context.Context, student *Model.Student) (bool, error) {
	var existingId int
	err := s.Db.QueryRow("SELECT studentid FROM students WHERE email = $1", student.Email).Scan(&existingId)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, handleSQLError(err)
	}
	return existingId > 0, nil
}

// Save implements StudentRepo.
func (s *repo_stud_impl) Save(ctx context.Context, student Model.Student) (bool, error) {
	// Inizia una transazione
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return false, handleSQLError(err)
	}
	defer tx.Rollback() // Assicurati che la transazione sia annullata in caso di errore

	// Prima, inserisci l'utente nella tabella Users e ottieni l'userID generato
	userSQL := "INSERT INTO users(username, password, roleID) VALUES ($1, $2, $3) RETURNING userID"
	var userID int
	err = tx.QueryRowContext(ctx, userSQL, student.Username, student.Password, student.Role).Scan(&userID)
	if err != nil {
		return false, handleSQLError(err)
	}

	// Poi, inserisci lo studente nella tabella Students con userID ottenuto
	studentSQL := "INSERT INTO students(studentID, name, surname, birthDate, address, email, departmentID) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = tx.ExecContext(ctx, studentSQL, userID, student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.DepartmentId)
	if err != nil {
		return false, handleSQLError(err)
	}

	// Se tutto va bene, conferma la transazione
	err = tx.Commit()
	if err != nil {
		return false, handleSQLError(err)
	}

	return true, nil
}

// Update implements StudentRepo.
func (s *repo_stud_impl) Update(ctx context.Context, student Model.Student) (bool, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	//var SQL string
	//var args []interface{}
	//if student.CourseId == 0 {
	SQL := "update students set name=$1, surname=$2, birthdate=$3, address=$4, departmentid=$5 where email=$6"
	//args = []interface{}{student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.CourseId, student.Id}
	//	} else {
	//		SQL = "update studenti set nome=$1, cognome=$2, datanascita=$3, indirizzo=$4, email=$5 where idstudente=$6"
	//		args = []interface{}{student.Name, student.Surname, student.Birthdate, student.Address, student.Email, student.Id}
	//	}
	res, err := tx.ExecContext(ctx, SQL, student.Name, student.Surname, student.Birthdate, student.Address, student.DepartmentId, student.Email)
	count, errx := res.RowsAffected()
	handleSQLError(errx)
	if count == 0 {
		return false, errors.New("student not found")
	}
	if err != nil {
		return false, handleSQLError(err)
	} else {
		return true, nil
	}
}
func handleSQLError(err error) error {
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code {
		case "23505":
			return errors.New("a student with this email already exists")
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
