package repository

import (
	"database/sql"
	en "go-server/EnumRole"
	"go-server/helper"
	utility "go-server/utility"
)

// struttura contenente un puntatore al db
type repo_user struct {
	Db *sql.DB
}

func NewUserRepo(Db *sql.DB) UserRepo {
	return &repo_user{Db: Db}
}

// VerifyUsername implements UserRepo.
func (s *repo_user) VerifyUsername(username string) bool {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Select COUNT(*) as us from users where username =$1"
	result, err := tx.Query(SQL, username)
	helper.PanicIfError(err)
	defer result.Close()
	var us int
	for result.Next() {

		result.Scan(&us)
	}
	return us > 0
}

// SignUp implements UserRepo.
func (s *repo_user) Signup(username string, password string) (bool, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Insert into users (username, password) values ($1, $2)"
	_, err = tx.Exec(SQL, username, password)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateToken implements UserRepo.
func (s *repo_user) UpdateToken(username string, token string) bool {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Update users set token = $1 where username = $2"
	_, err = tx.Exec(SQL, token, username)
	helper.PanicIfError(err)
	return err != nil
}

func (s *repo_user) VerifyCredentials(username string, password string) int {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx) // Corrected function name
	SQL := "Select roleid as us from users where username =$1 and password =$2"
	result, err := tx.Query(SQL, username, password) // Corrected receiver
	helper.PanicIfError(err)
	defer result.Close()
	for result.Next() {
		var us int
		err := result.Scan(&us)
		helper.PanicIfError(err)
		switch us {
		case 1:
			return en.Admin
		case 2:
			return en.Professor
		case 3:
			return en.Student
		default:
			return en.Unknown
		}

	}
	return en.Unknown
}

func (s *repo_user) VerifyIsAuthenticated(tokenString string) bool {
	_, _, valid := utility.IsValidToken(tokenString)
	if valid {
		tx, err := s.Db.Begin()
		helper.PanicIfError(err)
		defer helper.CommirOrRollback(tx)
		SQL := "Select COUNT(*) as us from users where token =$1 "
		result, err := tx.Query(SQL, tokenString)
		helper.PanicIfError(err)
		defer result.Close()
		for result.Next() {
			var us int
			err := result.Scan(&us)
			helper.PanicIfError(err)
			if us == 1 {
				return true
			}
		}
	}
	return false
}

func (s *repo_user) Logout(tokenString string) bool {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)
	SQL := "Update users set token = null where token = $1"
	_, errx := tx.Exec(SQL, tokenString)

	if errx != nil {
		helper.PanicIfError(errx)
		return false
	}
	return true
}
